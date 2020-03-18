/**
 * Copyright 2020 Appvia Ltd <info@appvia.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package eks

import (
	"context"
	"errors"
	"fmt"

	core "github.com/appvia/kore/pkg/apis/core/v1"
	eksv1alpha1 "github.com/appvia/kore/pkg/apis/eks/v1alpha1"
	"github.com/appvia/kore/pkg/utils/kubernetes"
	log "github.com/sirupsen/logrus"
	kerrors "k8s.io/apimachinery/pkg/api/errors"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	finalizerName = "eks.compute.kore.appvia.io"
	// ComponentClusterCreator is the name of the component for the UI
	ComponentClusterCreator = "Cluster Creator"
	// ComponentClusterBootstrap is the component name for seting up cloud credentials
	ComponentClusterBootstrap = "Cluster Initialize Access"
)

func (t *eksCtrl) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	ctx := context.Background()
	logger := log.WithFields(log.Fields{
		"name":      request.NamespacedName.Name,
		"namespace": request.NamespacedName.Namespace,
	})
	logger.Debug("attempting to reconcile aws eks cluster")

	resource := &eksv1alpha1.EKS{}
	if err := t.mgr.GetClient().Get(ctx, request.NamespacedName, resource); err != nil {
		if kerrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}

		return reconcile.Result{}, err
	}

	finalizer := kubernetes.NewFinalizer(t.mgr.GetClient(), finalizerName)

	// @step: are we deleting the resource
	if finalizer.IsDeletionCandidate(resource) {
		return t.Delete(request)
	}

	requeue, err := func() (bool, error) {

		logger.Debug("retrieving the gke cluster credential")
		// @step: first we need to check if we have access to the credentials
		credentials, err := t.GetCredentials(ctx, resource, resource.Namespace)
		if err != nil {
			logger.WithError(err).Error("trying to retrieve cloud credentials")

			resource.Status.Conditions.SetCondition(core.Component{
				Name:    ComponentClusterCreator,
				Message: "You do not have permission to the credentials",
				Status:  core.SuccessStatus,
			})

			return false, err
		}
		logger.Info("Found EKSCredential")

		client, err := NewClient(credentials, resource)
		if err != nil {
			logger.WithError(err).Error("attempting to create the cluster client")

			resource.Status.Conditions.SetCondition(core.Component{
				Detail:  err.Error(),
				Name:    ComponentClusterCreator,
				Message: "Failed to create EKS client, please check credentials",
				Status:  core.FailureStatus,
			})

			return false, err
		}
		logger.Info("Checking cluster existence")

		found, err := client.Exists()
		if err != nil {
			resource.Status.Conditions.SetCondition(core.Component{
				Detail:  err.Error(),
				Name:    ComponentClusterCreator,
				Message: "Failed to check for cluster existence",
				Status:  core.FailureStatus,
			})

			return false, err
		}

		if !found {
			status, found := resource.Status.Conditions.HasComponent(ComponentClusterCreator)
			if !found || status != core.PendingStatus {
				resource.Status.Conditions.SetCondition(core.Component{
					Name:    ComponentClusterCreator,
					Message: "Provisioning the EKS cluster in AWS",
					Status:  core.PendingStatus,
				})
				resource.Status.Status = core.PendingStatus

				return true, nil
			}

			logger.Debug("creating a new gke cluster in gcp")
			if _, err = client.Create(); err != nil {
				logger.WithError(err).Error("attempting to create cluster")

				resource.Status.Conditions.SetCondition(core.Component{
					Name:    ComponentClusterCreator,
					Message: "Failed trying to provision the cluster",
					Detail:  err.Error(),
				})
				resource.Status.Status = core.FailureStatus

				return false, err
			}
		} else {
			// TODO - client needs to manage migrations
		}

		// Get cluster status
		cluster, err := client.DescribeEKS()
		if err != nil {
			return false, err
		}
		// TODO: use strong type status check
		if *cluster.Status == "ERROR" {
			return false, fmt.Errorf("Cluster has ERROR status:%s", resource.Spec.Name)
		}
		if *cluster.Status != "ACTIVE" {
			// not ready, reque no errors
			return true, nil
		}
		// Active cluster
		resource.Status.CACertificate = cluster.CertificateAuthority.String()
		resource.Status.Endpoint = fmt.Sprintf("https://%s", cluster.Endpoint)
		resource.Status.Status = core.SuccessStatus

		// @step: update the state as provisioned
		resource.Status.Conditions.SetCondition(core.Component{
			Name:    ComponentClusterCreator,
			Message: "Cluster has been provisioned",
			Status:  core.SuccessStatus,
		})

		// @step: set the bootstrap as pending if required
		resource.Status.Conditions.SetCondition(core.Component{
			Name:    ComponentClusterBootstrap,
			Message: "Accessing the eks cluster",
			Status:  core.PendingStatus,
		})

		logger.Info("attempting to bootstrap the gke cluster")

		/*

			boot, err := NewBootstrapClient(resource, client.)
			if err != nil {
				logger.WithError(err).Error("trying to create bootstrap client")

				return false, err
			}
			if err := boot.Bootstrap(ctx, t.mgr.GetClient()); err != nil {
				logger.WithError(err).Error("trying to bootstrap gke cluster")

				return false, err
			}
		*/
		resource.Status.Conditions.SetCondition(core.Component{
			Name:    ComponentClusterBootstrap,
			Message: "Successfully initialised the cluster",
			Status:  core.SuccessStatus,
		})

		return false, nil
	}()
	if err != nil {
		resource.Status.Status = core.FailureStatus
	}

	if err := t.mgr.GetClient().Status().Update(ctx, resource); err != nil {
		logger.WithError(err).Error("updating the status of eks cluster")

		return reconcile.Result{}, err
	}

	if err == nil {
		if finalizer.NeedToAdd(resource) {
			logger.Info("adding our finalizer to the team resource")

			if err := finalizer.Add(resource); err != nil {
				return reconcile.Result{}, err
			}

			return reconcile.Result{Requeue: true}, nil
		}
	}

	if requeue {
		return reconcile.Result{Requeue: true}, nil
	}

	return reconcile.Result{}, nil
}

// GetCredentials returns the cloud credential
func (t *eksCtrl) GetCredentials(ctx context.Context, cluster *eksv1alpha1.EKS, team string) (*eksv1alpha1.EKSCredentials, error) {
	// @step: is the team permitted access to this credentials
	permitted, err := t.Teams().Team(team).Allocations().IsPermitted(ctx, cluster.Spec.Credentials)
	if err != nil {
		log.WithError(err).Error("attempting to check for permission on credentials")

		return nil, fmt.Errorf("attempting to check for permission on credentials")
	}

	if !permitted {
		log.Warn("trying to build gke cluster unallocated permissions")

		return nil, errors.New("you do not have permissions to the gke credentials")
	}

	// @step: retrieve the credentials
	creds := &eksv1alpha1.EKSCredentials{}

	return creds, t.mgr.GetClient().Get(ctx,
		types.NamespacedName{
			Namespace: cluster.Spec.Credentials.Namespace,
			Name:      cluster.Spec.Credentials.Name,
		}, creds,
	)
}
