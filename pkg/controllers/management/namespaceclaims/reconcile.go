/*
 * Copyright (C) 2019  Rohith Jayawardene <gambol99@gmail.com>
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package namespaceclaims

import (
	"context"
	"time"

	clustersv1 "github.com/appvia/kore/pkg/apis/clusters/v1"
	core "github.com/appvia/kore/pkg/apis/core/v1"
	"github.com/appvia/kore/pkg/controllers"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/utils/kubernetes"
	log "github.com/sirupsen/logrus"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	// finalizerName is our finalizer name
	finalizerName = "namespaceclaims.kore.appvia.io"
)

// Reconcile is resposible for reconciling the resource
func (a *nsCtrl) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	ctx := context.Background()

	// --- Logic ---
	// we have a client to the remote kubernetes cluster
	// we check if the team has a team namespace policy
	// we need to check the namespace is there and if not create it
	// we need to check the rolebinding exists and if not create it
	// we need to check that all the members of the team are in the binding
	// we set ourselves as the finalizer on the resource if not there already
	// we set the status of the resource to Success and the Phase is Installed
	// we sit back, relax and contain our smug smile

	logger := log.WithFields(log.Fields{
		"name":      request.Name,
		"namespace": request.Namespace,
	})
	logger.Debug("attempting to reconcile the nameresource claim")

	// @step: retrieve the resource from the api
	resource := &clustersv1.NamespaceClaim{}
	if err := a.mgr.GetClient().Get(ctx, request.NamespacedName, resource); err != nil {
		if !kerrors.IsNotFound(err) {
			return reconcile.Result{}, err
		}

		return reconcile.Result{}, nil
	}
	original := resource.DeepCopy()

	// @step: create a finalizer for the resource
	finalizer := kubernetes.NewFinalizer(a.mgr.GetClient(), finalizerName)

	if resource.GetDeletionTimestamp() != nil {
		if finalizer.IsDeletionCandidate(resource) {
			return a.Delete(request)
		}

		return reconcile.Result{}, nil
	}

	result, err := func() (reconcile.Result, error) {
		// @step: ensure the namespace is for a cluster you own
		if resource.Spec.Cluster.Namespace != resource.Namespace {
			resource.Status.Status = core.FailureStatus
			resource.Status.Conditions = []core.Condition{{
				Detail:  "access denied",
				Message: "cannot create namespace on cluster not owned by you",
			}}

			return reconcile.Result{}, nil
		}

		// @step: check the status of the cluster
		cluster := &clustersv1.Kubernetes{}
		if err := a.mgr.GetClient().Get(context.Background(), types.NamespacedName{
			Name:      resource.Spec.Cluster.Name,
			Namespace: resource.Spec.Cluster.Namespace,
		}, cluster); err != nil {
			if !kerrors.IsNotFound(err) {
				logger.WithError(err).Error("trying to retrieve the cluster")

				return reconcile.Result{}, err
			}

			// @checkpoint the cluster is not available yet
			resource.Status.Status = core.PendingStatus
			resource.Status.Conditions = []core.Condition{{
				Detail:  "cluster does not exist",
				Message: "no cluster: " + resource.Spec.Cluster.Name + " exist for this namespace",
			}}

			// @TODO we probably need a way of escaping this loop?
			return reconcile.Result{RequeueAfter: 3 * time.Minute}, nil
		}

		// @step: ignore the resource if already bootstrapped
		status, found := cluster.Status.Components.GetStatus("provision")
		if !found {
			logger.Warn("cluster does not have a status on the provisioning yet")

			resource.Status.Status = core.PendingStatus
			resource.Status.Conditions = []core.Condition{{
				Detail:  "cluster is pending, retrying later",
				Message: "cluster: " + resource.Spec.Cluster.Name + " is still pending",
			}}

			return reconcile.Result{RequeueAfter: 3 * time.Minute}, nil
		}
		switch status.Status {
		case core.PendingStatus:
			logger.Warn("cluster provision is not successfully yet, waiting")

			resource.Status.Status = core.PendingStatus
			resource.Status.Conditions = []core.Condition{{
				Detail:  "cluster has failed to provision, will retry",
				Message: "cluster " + resource.Spec.Cluster.Name + " is still pending",
			}}

			return reconcile.Result{RequeueAfter: 3 * time.Minute}, nil
		case core.SuccessStatus:
		default:
			resource.Status.Status = core.PendingStatus
			resource.Status.Conditions = []core.Condition{{
				Detail:  "cluster has failed to provision, will retry",
				Message: "cluster " + resource.Spec.Cluster.Name + " is in a failed state",
			}}

			return reconcile.Result{RequeueAfter: 3 * time.Minute}, nil
		}

		// @step: create credentials for the cluster
		client, err := controllers.CreateClientFromSecret(context.Background(), a.mgr.GetClient(),
			resource.Spec.Cluster.Namespace, resource.Spec.Cluster.Name)
		if err != nil {
			logger.WithError(err).Error("trying to create client from cluster secret")

			return reconcile.Result{}, err
		}

		// @step: ensure the namespace claim exists
		if err := kubernetes.EnsureNamespace(ctx, client, &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name:        resource.Spec.Name,
				Labels:      resource.Spec.Labels,
				Annotations: resource.Spec.Annotations,
			},
		}); err != nil {
			logger.WithError(err).Error("trying to provision the namespace in remote cluster")

			return reconcile.Result{}, err
		}

		// @step we need to check the rolebinding exists and if not create it
		logger.Debug("ensuring the binding to the namespace admin exists")

		binding := &rbacv1.RoleBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name:      RoleBindingName,
				Namespace: resource.Spec.Name,
				Labels: map[string]string{
					kore.Label("owned"): "true",
				},
			},
			RoleRef: rbacv1.RoleRef{
				APIGroup: rbacv1.GroupName,
				Kind:     "ClusterRole",
				Name:     ClusterRoleName,
			},
		}

		// @step: retrieve all the users in the team
		users, err := a.Teams().Team(request.Namespace).Members().List(ctx)
		if err != nil {
			logger.WithError(err).Error("trying to retrieve a list of members in the team")

			return reconcile.Result{}, err
		}

		logger.WithField(
			"users", len(users.Items),
		).Debug("found the x members in the team")

		for _, x := range users.Items {
			binding.Subjects = append(binding.Subjects, rbacv1.Subject{
				APIGroup: rbacv1.GroupName,
				Kind:     rbacv1.UserKind,
				Name:     x.Spec.Username,
			})
		}

		// @step: ensuring the binding exists
		if _, err := kubernetes.CreateOrUpdate(ctx, client, binding); err != nil {
			logger.WithError(err).Error("trying to ensure the namespace team binding")

			return reconcile.Result{}, err
		}

		resource.Status.Status = core.SuccessStatus
		resource.Status.Conditions = []core.Condition{}

		return reconcile.Result{}, nil
	}()
	if err != nil {
		logger.WithError(err).Error("trying to reconcile the nameresource claim")

		resource.Status.Status = core.FailureStatus
		resource.Status.Conditions = []core.Condition{{
			Message: "failed trying to reconcile the nameresource claim",
			Detail:  err.Error(),
		}}
	} else {
		if finalizer.NeedToAdd(resource) {
			if err := finalizer.Add(resource); err != nil {
				logger.WithError(err).Error("trying to add the finalizer")

				return reconcile.Result{}, err
			}

			return reconcile.Result{Requeue: true}, nil
		}
	}

	if err := a.mgr.GetClient().Status().Patch(ctx, resource, client.MergeFrom(original)); err != nil {
		logger.WithError(err).Error("trying to update the status of the resource")

		return reconcile.Result{}, err
	}

	return result, err
}
