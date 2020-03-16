/**
 * Copyright (C) 2020 Appvia Ltd <info@appvia.io>
 *
 * This file is part of kore.
 *
 * kore is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 2 of the License, or
 * (at your option) any later version.
 *
 * kore is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with kore.  If not, see <http://www.gnu.org/licenses/>.
 */

package projectclaim

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	configv1 "github.com/appvia/kore/pkg/apis/config/v1"
	gcp "github.com/appvia/kore/pkg/apis/gcp/v1alpha1"

	cloudresourcemanager "google.golang.org/api/cloudresourcemanager/v1"
	iam "google.golang.org/api/iam/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetServiceAccountKeyID returns the service account id portion from the name
// e.g kore-demo-1@kore-demo-1.iam.gserviceaccount.com/keys/35ac8390ee7d579af69108b8e32d1d05163ac73d
func GetServiceAccountKeyID(name string) string {
	items := strings.Split(name, "/")

	return items[len(items)-1]
}

// findOldestServiceAccountKey is used to filter out the oldest service accout
// @note: i can't find anyway around this at the moment; the service acount key include
// a default service account key which is managed by google and cannot be delete. There
//
func findOldestServiceAccountKey(keys []*iam.ServiceAccountKey) (*iam.ServiceAccountKey, error) {
	// @step: set the oldest to the first for now
	current := keys[0]
	oldest, err := time.Parse(time.RFC3339, current.ValidBeforeTime)
	if err != nil {
		return nil, err
	}

	// @step: find the oldest key
	for i := 1; i < len(keys); i++ {
		before, err := time.Parse(time.RFC3339, keys[i].ValidBeforeTime)
		if err != nil {
			return nil, err
		}
		if oldest.After(before) {
			current = keys[i]
			oldest = before
		}
	}

	return current, nil
}

// IsCredentialsValid checks the secert is cool
func IsCredentialsValid(secret *configv1.Secret) error {
	for _, name := range []string{ServiceAccountKey, ExpiryKey, ProjectIDKey, ProjectNameKey} {
		if _, found := secret.Spec.Data[name]; !found {
			return fmt.Errorf("secret does not have the %s field, we need to regenerate one", name)
		}
	}

	expires := secret.Spec.Data[ExpiryKey]
	if _, err := strconv.ParseInt(expires, 10, 64); err != nil {
		return fmt.Errorf("invalid expires field")
	}

	return nil
}

// CreateCredentialsSecret returns a project credentials secret
func CreateCredentialsSecret(project *gcp.ProjectClaim, name string, values map[string]string) *configv1.Secret {
	secret := &configv1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: configv1.GroupVersion.String(),
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: project.Namespace,
			OwnerReferences: []metav1.OwnerReference{{
				APIVersion: gcp.GroupVersion.String(),
				Controller: &isTrue,
				Kind:       "ProjectClaim",
				Name:       project.GetName(),
				UID:        project.GetUID(),
			}},
		},
		Spec: configv1.SecretSpec{
			Data:        values,
			Description: fmt.Sprintf("GCP Project credentials for team project: %s", project.Name),
			Type:        configv1.GenericSecret,
		},
	}

	return secret
}

// IsProject checks if the project exists
func IsProject(ctx context.Context, client *cloudresourcemanager.Service, name string) (*cloudresourcemanager.Project, bool, error) {
	list, err := client.Projects.List().Context(ctx).Do()
	if err != nil {
		return nil, false, err
	}

	for _, x := range list.Projects {
		if x.Name == name {
			return x, true, nil
		}
	}

	return nil, false, nil
}
