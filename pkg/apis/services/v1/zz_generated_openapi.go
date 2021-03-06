// +build !ignore_autogenerated

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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/appvia/kore/pkg/apis/services/v1.Service":                  schema_pkg_apis_services_v1_Service(ref),
		"github.com/appvia/kore/pkg/apis/services/v1.ServiceCredentials":       schema_pkg_apis_services_v1_ServiceCredentials(ref),
		"github.com/appvia/kore/pkg/apis/services/v1.ServiceCredentialsSpec":   schema_pkg_apis_services_v1_ServiceCredentialsSpec(ref),
		"github.com/appvia/kore/pkg/apis/services/v1.ServiceCredentialsStatus": schema_pkg_apis_services_v1_ServiceCredentialsStatus(ref),
		"github.com/appvia/kore/pkg/apis/services/v1.ServiceKind":              schema_pkg_apis_services_v1_ServiceKind(ref),
		"github.com/appvia/kore/pkg/apis/services/v1.ServiceKindSpec":          schema_pkg_apis_services_v1_ServiceKindSpec(ref),
		"github.com/appvia/kore/pkg/apis/services/v1.ServicePlan":              schema_pkg_apis_services_v1_ServicePlan(ref),
		"github.com/appvia/kore/pkg/apis/services/v1.ServicePlanSpec":          schema_pkg_apis_services_v1_ServicePlanSpec(ref),
		"github.com/appvia/kore/pkg/apis/services/v1.ServiceProvider":          schema_pkg_apis_services_v1_ServiceProvider(ref),
		"github.com/appvia/kore/pkg/apis/services/v1.ServiceProviderSpec":      schema_pkg_apis_services_v1_ServiceProviderSpec(ref),
		"github.com/appvia/kore/pkg/apis/services/v1.ServiceProviderStatus":    schema_pkg_apis_services_v1_ServiceProviderStatus(ref),
		"github.com/appvia/kore/pkg/apis/services/v1.ServiceSpec":              schema_pkg_apis_services_v1_ServiceSpec(ref),
		"github.com/appvia/kore/pkg/apis/services/v1.ServiceStatus":            schema_pkg_apis_services_v1_ServiceStatus(ref),
	}
}

func schema_pkg_apis_services_v1_Service(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Service is a managed service instance",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/services/v1.ServiceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/services/v1.ServiceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/services/v1.ServiceSpec", "github.com/appvia/kore/pkg/apis/services/v1.ServiceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_services_v1_ServiceCredentials(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceCredentials is credentials provisioned by a service into the target namespace",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/services/v1.ServiceCredentialsSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/services/v1.ServiceCredentialsStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/services/v1.ServiceCredentialsSpec", "github.com/appvia/kore/pkg/apis/services/v1.ServiceCredentialsStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_services_v1_ServiceCredentialsSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceCredentialsSpec defines the the desired status for service credentials",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind refers to the service type",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"service": {
						SchemaProps: spec.SchemaProps{
							Description: "Service contains the reference to the service object",
							Ref:         ref("github.com/appvia/kore/pkg/apis/core/v1.Ownership"),
						},
					},
					"cluster": {
						SchemaProps: spec.SchemaProps{
							Description: "Cluster contains the reference to the cluster where the credentials will be saved as a secret",
							Ref:         ref("github.com/appvia/kore/pkg/apis/core/v1.Ownership"),
						},
					},
					"clusterNamespace": {
						SchemaProps: spec.SchemaProps{
							Description: "ClusterNamespace is the target namespace in the cluster where the secret will be created",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"secretName": {
						SchemaProps: spec.SchemaProps{
							Description: "SecretName is the Kubernetes Secret's name that will contain the service access information If not set the secret's name will default to `Name`",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configuration": {
						SchemaProps: spec.SchemaProps{
							Description: "Configuration are the configuration values for this service credentials It will be used by the service provider to provision the credentials",
							Ref:         ref("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.JSON"),
						},
					},
				},
				Required: []string{"kind"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/core/v1.Ownership", "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.JSON"},
	}
}

func schema_pkg_apis_services_v1_ServiceCredentialsStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceCredentialsStatus defines the observed state of a service",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"components": {
						SchemaProps: spec.SchemaProps{
							Description: "Components is a collection of component statuses",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/appvia/kore/pkg/apis/core/v1.Component"),
									},
								},
							},
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status is the overall status of the service",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Message is the description of the current status",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"providerID": {
						SchemaProps: spec.SchemaProps{
							Description: "ProviderID is the service credentials identifier in the service provider",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"providerData": {
						SchemaProps: spec.SchemaProps{
							Description: "ProviderData is provider specific data",
							Ref:         ref("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.JSON"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/core/v1.Component", "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.JSON"},
	}
}

func schema_pkg_apis_services_v1_ServiceKind(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceKind is a service type",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/services/v1.ServiceKindSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/services/v1.ServiceKindSpec", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_services_v1_ServiceKindSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceKindSpec defines the state of a service kind",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"enabled": {
						SchemaProps: spec.SchemaProps{
							Description: "Enabled is true if the service kind can be used",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"displayName": {
						SchemaProps: spec.SchemaProps{
							Description: "DisplayName refers to the display name of the service type",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"summary": {
						SchemaProps: spec.SchemaProps{
							Description: "Summary provides a short title summary for the service kind",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Description is a detailed description of the service kind",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"imageURL": {
						SchemaProps: spec.SchemaProps{
							Description: "ImageURL is a thumbnail for the service kind",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"documentationURL": {
						SchemaProps: spec.SchemaProps{
							Description: "DocumentationURL refers to the documentation page for this service",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"schema": {
						SchemaProps: spec.SchemaProps{
							Description: "Schema is the JSON schema for the plan",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"credentialSchema": {
						SchemaProps: spec.SchemaProps{
							Description: "CredentialSchema is the JSON schema for credentials created for service using this plan",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"providerData": {
						SchemaProps: spec.SchemaProps{
							Description: "ProviderData is provider specific data",
							Ref:         ref("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.JSON"),
						},
					},
				},
				Required: []string{"enabled", "summary"},
			},
		},
		Dependencies: []string{
			"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.JSON"},
	}
}

func schema_pkg_apis_services_v1_ServicePlan(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServicePlan is a template for a service",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/services/v1.ServicePlanSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/services/v1.ServicePlanSpec", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_services_v1_ServicePlanSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServicePlanSpec defines the desired state of Service plan",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind refers to the service type this is a plan for",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"displayName": {
						SchemaProps: spec.SchemaProps{
							Description: "DisplayName refers to the display name of the service type",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"labels": {
						SchemaProps: spec.SchemaProps{
							Description: "Labels is a collection of labels for this plan",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"summary": {
						SchemaProps: spec.SchemaProps{
							Description: "Summary provides a short title summary for the plan",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Description is a detailed description of the service plan",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configuration": {
						SchemaProps: spec.SchemaProps{
							Description: "Configuration are the key+value pairs describing a service configuration",
							Ref:         ref("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.JSON"),
						},
					},
					"schema": {
						SchemaProps: spec.SchemaProps{
							Description: "Schema is the JSON schema for the plan",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"credentialSchema": {
						SchemaProps: spec.SchemaProps{
							Description: "CredentialSchema is the JSON schema for credentials created for service using this plan",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"providerData": {
						SchemaProps: spec.SchemaProps{
							Description: "ProviderData is provider specific data",
							Ref:         ref("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.JSON"),
						},
					},
				},
				Required: []string{"kind", "summary"},
			},
		},
		Dependencies: []string{
			"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.JSON"},
	}
}

func schema_pkg_apis_services_v1_ServiceProvider(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceProvider is a template for a service provider",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/services/v1.ServiceProviderSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/appvia/kore/pkg/apis/services/v1.ServiceProviderStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/services/v1.ServiceProviderSpec", "github.com/appvia/kore/pkg/apis/services/v1.ServiceProviderStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_services_v1_ServiceProviderSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceProviderSpec defines the desired state of a Service provider",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type refers to the service provider type",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"summary": {
						SchemaProps: spec.SchemaProps{
							Description: "Summary provides a short title summary for the provider",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Description is a detailed description of the service provider",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configuration": {
						SchemaProps: spec.SchemaProps{
							Description: "Configuration are the key+value pairs describing a service provider",
							Ref:         ref("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.JSON"),
						},
					},
					"credentials": {
						SchemaProps: spec.SchemaProps{
							Description: "Credentials is a reference to the credentials object to use",
							Ref:         ref("github.com/appvia/kore/pkg/apis/core/v1.Ownership"),
						},
					},
				},
				Required: []string{"type", "summary"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/core/v1.Ownership", "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.JSON"},
	}
}

func schema_pkg_apis_services_v1_ServiceProviderStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceProviderStatus defines the observed state of a service provider",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status is the overall status of the service",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Message is the description of the current status",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"components": {
						SchemaProps: spec.SchemaProps{
							Description: "Components is a collection of component statuses",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/appvia/kore/pkg/apis/core/v1.Component"),
									},
								},
							},
						},
					},
					"supportedKinds": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "SupportedKinds contains all the supported service kinds",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/core/v1.Component"},
	}
}

func schema_pkg_apis_services_v1_ServiceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceSpec defines the desired state of a service",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind refers to the service type",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"plan": {
						SchemaProps: spec.SchemaProps{
							Description: "Plan is the name of the service plan which was used to create this service",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"cluster": {
						SchemaProps: spec.SchemaProps{
							Description: "Cluster contains the reference to the cluster where the service will be created",
							Ref:         ref("github.com/appvia/kore/pkg/apis/core/v1.Ownership"),
						},
					},
					"clusterNamespace": {
						SchemaProps: spec.SchemaProps{
							Description: "ClusterNamespace is the target namespace in the cluster where there the service will be created",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configuration": {
						SchemaProps: spec.SchemaProps{
							Description: "Configuration are the configuration values for this service It will contain values from the plan + overrides by the user This will provide a simple interface to calculate diffs between plan and service configuration",
							Ref:         ref("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.JSON"),
						},
					},
					"credentials": {
						SchemaProps: spec.SchemaProps{
							Description: "Credentials is a reference to the credentials object to use",
							Ref:         ref("github.com/appvia/kore/pkg/apis/core/v1.Ownership"),
						},
					},
				},
				Required: []string{"kind", "plan"},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/core/v1.Ownership", "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.JSON"},
	}
}

func schema_pkg_apis_services_v1_ServiceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceStatus defines the observed state of a service",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"components": {
						SchemaProps: spec.SchemaProps{
							Description: "Components is a collection of component statuses",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/appvia/kore/pkg/apis/core/v1.Component"),
									},
								},
							},
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status is the overall status of the service",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "Message is the description of the current status",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"providerID": {
						SchemaProps: spec.SchemaProps{
							Description: "ProviderID is the service identifier in the service provider",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"providerData": {
						SchemaProps: spec.SchemaProps{
							Description: "ProviderData is provider specific data",
							Ref:         ref("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.JSON"),
						},
					},
					"plan": {
						SchemaProps: spec.SchemaProps{
							Description: "Plan is the name of the service plan which was used to create this service",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configuration": {
						SchemaProps: spec.SchemaProps{
							Description: "Configuration are the applied configuration values for this service",
							Ref:         ref("k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.JSON"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/appvia/kore/pkg/apis/core/v1.Component", "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1.JSON"},
	}
}
