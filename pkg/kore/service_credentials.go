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

package kore

import (
	"context"

	"github.com/appvia/kore/pkg/utils"

	clustersv1 "github.com/appvia/kore/pkg/apis/clusters/v1"

	"github.com/appvia/kore/pkg/utils/jsonschema"

	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/store"
	"github.com/appvia/kore/pkg/utils/validation"

	log "github.com/sirupsen/logrus"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
)

// ServiceCredentials returns the an interface for handling service credentials
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . ServiceCredentials
type ServiceCredentials interface {
	// Delete is used to delete service credentials
	Delete(context.Context, string) (*servicesv1.ServiceCredentials, error)
	// Get returns a specific service credentials
	Get(context.Context, string) (*servicesv1.ServiceCredentials, error)
	// GetSchema returns the service credential schema
	GetSchema(context.Context, string) (string, error)
	// List returns a list of service credentials
	List(context.Context) (*servicesv1.ServiceCredentialsList, error)
	// Update is used to update service credentials
	Update(context.Context, *servicesv1.ServiceCredentials) error
}

type serviceCredentialsImpl struct {
	*hubImpl
	// team is the name
	team string
}

// Delete is used to delete service credentials
func (s *serviceCredentialsImpl) Delete(ctx context.Context, name string) (*servicesv1.ServiceCredentials, error) {
	logger := log.WithFields(log.Fields{
		"serviceCredentials": name,
		"team":               s.team,
	})
	logger.Info("attempting to delete the service credentials")

	original, err := s.Get(ctx, name)
	if err != nil {
		if err == ErrNotFound {
			return nil, err
		}

		logger.WithError(err).Error("failed to retrieve the service credentials")

		return nil, err
	}

	return original, s.Store().Client().Delete(ctx, store.DeleteOptions.From(original))
}

// List returns a list of service credentials we have access to
func (s *serviceCredentialsImpl) List(ctx context.Context) (*servicesv1.ServiceCredentialsList, error) {
	list := &servicesv1.ServiceCredentialsList{}

	return list, s.Store().Client().List(ctx,
		store.ListOptions.InNamespace(s.team),
		store.ListOptions.InTo(list),
	)
}

// Get returns specific service credentials
func (s *serviceCredentialsImpl) Get(ctx context.Context, name string) (*servicesv1.ServiceCredentials, error) {
	serviceCredentials := &servicesv1.ServiceCredentials{}

	if err := s.Store().Client().Get(ctx,
		store.GetOptions.InNamespace(s.team),
		store.GetOptions.InTo(serviceCredentials),
		store.GetOptions.WithName(name),
	); err != nil {
		if kerrors.IsNotFound(err) {
			return nil, ErrNotFound
		}

		log.WithError(err).Error("failed to retrieve the service credentials")
		return nil, err
	}

	return serviceCredentials, nil
}

// GetSchema returns the service plan schema
func (s *serviceCredentialsImpl) GetSchema(ctx context.Context, name string) (string, error) {
	serviceCredentials, err := s.Get(ctx, name)
	if err != nil {
		return "", err
	}

	service, err := s.Teams().Team(serviceCredentials.Spec.Service.Namespace).Services().Get(ctx, serviceCredentials.Spec.Service.Name)
	if err != nil {
		return "", err
	}

	plan, err := s.ServicePlans().Get(ctx, service.Spec.Plan)
	if err != nil {
		return "", err
	}
	if plan.Spec.Schema != "" {
		return plan.Spec.Schema, nil
	}

	kind, err := s.ServiceKinds().Get(ctx, plan.Spec.Kind)
	if err != nil {
		return "", err
	}

	return kind.Spec.Schema, nil
}

// Update is used to update service credentials
func (s *serviceCredentialsImpl) Update(ctx context.Context, serviceCreds *servicesv1.ServiceCredentials) error {
	existing, err := s.Get(ctx, serviceCreds.Name)
	if err != nil && err != ErrNotFound {
		return err
	}

	if existing != nil {
		verr := validation.NewError("%q failed validation", serviceCreds.Name)
		if existing.Spec.Kind != serviceCreds.Spec.Kind {
			verr.AddFieldErrorf("kind", validation.ReadOnly, "can not be changed after the service credentials was created")
		}
		if verr.HasErrors() {
			return verr
		}
	}

	if serviceCreds.Namespace == "" {
		serviceCreds.Namespace = s.team
	}

	if serviceCreds.Namespace != s.team {
		return validation.NewError("%q failed validation", serviceCreds.Name).WithFieldErrorf(
			"namespace",
			validation.InvalidValue,
			"must be the same as the team name: %q",
			s.team,
		)
	}

	service, err := s.validateService(ctx, serviceCreds)
	if err != nil {
		return err
	}

	if err := s.validateCluster(ctx, serviceCreds); err != nil {
		return err
	}

	if err := s.validateConfiguration(ctx, service, serviceCreds, existing); err != nil {
		return err
	}

	return s.Store().Client().Update(ctx,
		store.UpdateOptions.To(serviceCreds),
		store.UpdateOptions.WithCreate(true),
	)
}

func (s *serviceCredentialsImpl) validateConfiguration(
	ctx context.Context,
	service *servicesv1.Service,
	serviceCreds, existing *servicesv1.ServiceCredentials,
) error {
	schema, err := s.ServicePlans().GetCredentialSchema(ctx, service.Spec.Plan)
	if err != nil {
		return err
	}

	if schema == "" && !utils.ApiExtJSONEmpty(serviceCreds.Spec.Configuration) {
		if existing == nil || !utils.ApiExtJSONEquals(serviceCreds.Spec.Configuration, existing.Spec.Configuration) {
			return validation.NewError("%q failed validation", serviceCreds.Name).
				WithFieldErrorf(
					"configuration",
					validation.ReadOnly,
					"the service provider doesn't have a JSON schema to validate the configuration",
				)
		}
	}

	if schema != "" {
		if err := jsonschema.Validate(
			schema,
			"configuration",
			serviceCreds.Spec.Configuration,
		); err != nil {
			return err
		}
	}

	return nil
}

func (s *serviceCredentialsImpl) validateService(ctx context.Context, serviceCreds *servicesv1.ServiceCredentials) (*servicesv1.Service, error) {
	if serviceCreds.Spec.Service.Namespace != serviceCreds.Namespace {
		return nil, validation.NewError("%q failed validation", serviceCreds.Name).WithFieldErrorf(
			"service.namespace",
			validation.InvalidValue,
			"must be the same as the team name: %q",
			s.team,
		)
	}

	if !serviceCreds.Spec.Service.HasGroupVersionKind(servicesv1.ServiceGVK) {
		return nil, validation.NewError("%q failed validation", serviceCreds.Name).WithFieldErrorf(
			"service",
			validation.InvalidValue,
			"must have type of %s",
			servicesv1.ServiceGVK,
		)
	}

	service, err := s.Teams().Team(s.team).Services().Get(ctx, serviceCreds.Spec.Service.Name)
	if err != nil {
		if kerrors.IsNotFound(err) {
			return nil, validation.NewError("%q failed validation", serviceCreds.Name).WithFieldError(
				"service.name",
				validation.MustExist,
				"does not exist",
			)
		}
	}
	return service, nil
}

func (s *serviceCredentialsImpl) validateCluster(ctx context.Context, serviceCreds *servicesv1.ServiceCredentials) error {
	if serviceCreds.Spec.Cluster.Namespace != serviceCreds.Namespace {
		return validation.NewError("%q failed validation", serviceCreds.Name).WithFieldErrorf(
			"cluster.namespace",
			validation.InvalidValue,
			"must be the same as the team name: %q",
			s.team,
		)
	}

	if !serviceCreds.Spec.Cluster.HasGroupVersionKind(clustersv1.ClusterGroupVersionKind) {
		return validation.NewError("%q failed validation", serviceCreds.Name).WithFieldErrorf(
			"cluster",
			validation.InvalidValue,
			"must have type of %s",
			clustersv1.ClusterGroupVersionKind,
		)
	}

	_, err := s.Teams().Team(s.team).Clusters().Get(ctx, serviceCreds.Spec.Cluster.Name)
	if err != nil {
		if kerrors.IsNotFound(err) {
			return validation.NewError("%q failed validation", serviceCreds.Name).WithFieldError(
				"cluster.name",
				validation.MustExist,
				"does not exist",
			)
		}
	}
	return nil
}
