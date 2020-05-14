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

package clusterappman

import (
	"testing"

	"github.com/appvia/kore/pkg/clusterapp"
	"github.com/appvia/kore/pkg/utils/kubernetes"
	"github.com/stretchr/testify/assert"
	cc "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestLoadAllManifests(t *testing.T) {
	options, err := clusterapp.GetClientOptions()
	assert.NoError(t, err)
	client := cc.NewFakeClientWithScheme(options.Scheme)

	// Load all manifests - we won't need them to re-create any API access
	err = LoadAllManifests(client, kubernetes.KubernetesAPI{})
	assert.NoError(t, err)
}
