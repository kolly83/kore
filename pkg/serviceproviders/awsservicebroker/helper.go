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

package awsservicebroker

import (
	"context"
	"errors"
	"fmt"
	"strings"

	eksv1alpha1 "github.com/appvia/kore/pkg/apis/eks/v1alpha1"
	servicesv1 "github.com/appvia/kore/pkg/apis/services/v1"
	"github.com/appvia/kore/pkg/controllers"
	"github.com/appvia/kore/pkg/kore"

	"github.com/aws/aws-sdk-go/aws/awserr"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"
)

func isAWSErr(err error, code string, message string) bool {
	var awsErr awserr.Error
	if errors.As(err, &awsErr) {
		return awsErr.Code() == code && strings.Contains(awsErr.Message(), message)
	}
	return false
}

func isAWSErrRequestFailureStatusCode(err error, statusCode int) bool {
	var awsErr awserr.RequestFailure
	if errors.As(err, &awsErr) {
		return awsErr.StatusCode() == statusCode
	}
	return false
}

func getServiceAccountToken(ctx context.Context, client client.Client, namespace, name string) (_ *corev1.Secret, _ error) {
	sa := &corev1.ServiceAccount{}
	err := client.Get(ctx, types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}, sa)
	if err != nil {
		return nil, fmt.Errorf("failed to get serviceaccount %q: %w", name, err)
	}
	if len(sa.Secrets) <= 0 {
		return nil, fmt.Errorf("no secrets found in serviceaccount %q", name)
	}

	return getSecret(ctx, client, namespace, sa.Secrets[0].Name)
}

func getSecret(ctx context.Context, client client.Client, namespace, name string) (*corev1.Secret, error) {
	secret := &corev1.Secret{}
	err := client.Get(ctx, types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}, secret)
	if err != nil {
		return nil, fmt.Errorf("failed to get secret %q: %w", name, err)
	}

	return secret, nil
}

func getCredentials(ctx kore.Context, serviceProvider *servicesv1.ServiceProvider) (awsAccessKeyID string, awsSecretAccessKey string, _ error) {
	eksCredentials := &eksv1alpha1.EKSCredentials{}
	err := ctx.Client().Get(ctx, serviceProvider.Spec.Credentials.NamespacedName(), eksCredentials)
	if err != nil {
		return "", "", fmt.Errorf("failed to load the service provider credentials: %w", err)
	}

	if eksCredentials.Spec.CredentialsRef != nil {
		secret, err := controllers.GetDecodedSecret(ctx, ctx.Client(), eksCredentials.Spec.CredentialsRef)
		if err != nil {
			return "", "", fmt.Errorf("failed to load the credentials secret: %w", err)
		}
		awsAccessKeyID = secret.Spec.Data["access_key_id"]
		awsSecretAccessKey = secret.Spec.Data["access_secret_key"]
	} else {
		awsAccessKeyID = eksCredentials.Spec.AccessKeyID
		awsSecretAccessKey = eksCredentials.Spec.SecretAccessKey
	}

	return
}
