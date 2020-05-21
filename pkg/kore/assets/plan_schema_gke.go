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

package assets

// GKEPlanSchema is the JSON schema used to describe and validate GKE Plans
const GKEPlanSchema = `
{
	"$id": "https://appvia.io/schemas/gke/plan.json",
	"$schema": "http://json-schema.org/draft-07/schema#",
	"description": "GKE Cluster Plan Schema",
	"type": "object",
	"additionalProperties": false,
	"required": [
		"authorizedMasterNetworks",
		"authProxyAllowedIPs",
		"description",
		"domain",
		"enableDefaultTrafficBlock",
		"enableHTTPLoadBalancer",
		"enableHorizontalPodAutoscaler",
		"enableIstio",
		"enablePrivateEndpoint",
		"enablePrivateNetwork",
		"enableShieldedNodes",
		"enableStackDriverLogging",
		"enableStackDriverMetrics",
		"inheritTeamMembers",
		"maintenanceWindow",
		"network",
		"nodePools",
		"region",
		"version",
		"releaseChannel"
	],
	"properties": {
		"authorizedMasterNetworks": {
			"type": "array",
			"description": "The networks which are allowed to access the master control plane.",
			"items": {
				"type": "object",
				"additionalProperties": false,
				"required": [
					"name",
					"cidr"
				],
				"properties": {
					"name": {
						"type": "string",
						"minLength": 1
					},
					"cidr": {
						"type": "string",
						"format": "1.2.3.4/16"
					}
				}
			},
			"minItems": 1
		},
		"authProxyAllowedIPs": {
			"title": "Auth Proxy Allowed IP Ranges",
			"type": "array",
			"description": "The networks which are allowed to connect to this cluster (e.g. via kubectl).",
			"items": {
				"type": "string",
				"format": "1.2.3.4/16"
			},
			"minItems": 1
		},
		"clusterUsers": {
			"type": "array",
			"description": "Users who should be allowed to access this cluster.",
			"items": {
				"type": "object",
				"additionalProperties": false,
				"required": [
					"username",
					"roles"
				],
				"properties": {
					"username": {
						"type": "string",
						"minLength": 1
					},
					"roles": {
						"type": "array",
						"items": {
							"type": "string",
							"minLength": 1,
							"enum": [ "view", "edit", "admin", "cluster-admin" ]
						},
						"minItems": 1
					}
				}
			}
		},
		"defaultTeamRole": {
			"type": "string",
			"description": "The default role that team members have on this cluster.",
			"enum": [ "view", "edit", "admin", "cluster-admin" ]
		},
		"description": {
			"type": "string",
			"description": "Meaningful description of this cluster.",
			"minLength": 1
		},
		"domain": {
			"type": "string",
			"description": "The domain for this cluster.",
			"minLength": 1,
			"immutable": true
		},
		"enableDefaultTrafficBlock": {
			"type": "boolean"
		},
		"enableHTTPLoadBalancer": {
			"type": "boolean"
		},
		"enableHorizontalPodAutoscaler": {
			"type": "boolean",
			"immutable": true
		},
		"enableIstio": {
			"type": "boolean",
			"immutable": true
		},
		"enablePrivateEndpoint": {
			"type": "boolean",
			"immutable": true
		},
		"enablePrivateNetwork": {
			"type": "boolean",
			"immutable": true
		},
		"enableShieldedNodes": {
			"type": "boolean",
			"description": "Shielded nodes provide additional verifications of the node OS and VM, with enhanced rootkit and bootkit protection applied",
			"immutable": true
		},
		"enableStackDriverLogging": {
			"type": "boolean",
			"immutable": true
		},
		"enableStackDriverMetrics": {
			"type": "boolean",
			"immutable": true
		},
		"inheritTeamMembers": {
			"type": "boolean"
		},
		"maintenanceWindow": {
			"type": "string",
			"description": "Time of day to allow maintenance operations to be performed by the cloud provider on this cluster.",
			"format": "hh:mm",
			"immutable": true
		},
		"network": {
			"type": "string",
			"minLength": 1,
			"description": "The GCP network that this cluster should reside on. If specified, this network must exist.",
			"immutable": true
		},
		"nodePools": {
			"type": "array",
			"items": {
				"type": "object",
				"additionalProperties": false,
				"required": [
					"name",
					"enableAutoscaler",
					"enableAutoupgrade",
					"enableAutorepair",
					"version",
					"minSize",
					"maxSize",
					"size",
					"machineType",
					"imageType",
					"diskSize"
				],
				"properties": {
					"name": {
						"type": "string",
						"minLength": 1,
						"description": "Name of this node pool. Must be unique within the cluster.",
						"immutable": true
					},
					"enableAutoupgrade": {
						"type": "boolean",
						"description": "Enable to update this node pool updated when new GKE versions are made available by GCP - must be enabled if a release channel is selected",
						"default": true
					},
					"version": {
						"type": "string",
						"description": "Node pool version, '-' to use same version as cluster (recommended). Must be blank if cluster follows a release channel. Must be within 2 minor versions of the master version (e.g. for master version 1.16, this must be 1.14, 1.15 or 1.16)",
						"pattern": "^($|-|latest|[0-9]+\\.[0-9]+($|\\.[0-9]+($|\\-gke\\.[0-9]+)))$",
						"default": "",
						"examples": [
							"-", "latest", "1.15 (latest 1.15.x-gke.y)", "1.15.1 (latest 1.15.1-gke.x)", "1.15.1-gke.6 (exact GKE version)"
						]
					},
					"enableAutoscaler": {
						"type": "boolean",
						"default": true,
						"description": "Add and remove nodes automatically based on load"
					},
					"enableAutorepair": {
						"type": "boolean",
						"default": true,
						"description": "Automatically repair any failed nodes within this node pool."
					},
					"minSize": {
						"type": "number",
						"multipleOf": 1,
						"minimum": 1,
						"default": 1,
						"description": "The minimum nodes this pool should contain (if auto-scale enabled)"
					},
					"maxSize": {
						"type": "number",
						"multipleOf": 1,
						"minimum": 1,
						"default": 10,
						"description": "The maximum nodes this pool should contain (if auto-scale enabled)"
					},
					"size": {
						"type": "number",
						"multipleOf": 1,
						"minimum": 1,
						"default": 1,
						"description": "How many nodes to build when provisioning this pool - if autoscaling enabled, this will be the initial size",
						"immutable": true
					},
					"maxPodsPerNode": {
						"type": "number",
						"multipleOf": 1,
						"description": "The maximum number of pods that can be scheduled onto each node of this pool",
						"default": 110,
						"maximum": 110,
						"minimum": 8,
						"immutable": true
					},
					"machineType": {
						"type": "string",
						"description": "The type of nodes used for this node pool",
						"pattern": "^[a-z][0-9]\\-(micro|small|medium|standard\\-[0-9]+|highmem\\-[0-9]+|highcpu\\-[0-9]+|ultramem\\-[0-9]+|megamem\\-[0-9]+)$",
						"default": "n2-standard-2",
						"minLength": 1,
						"immutable": true
					},
					"imageType": {
						"type": "string",
						"enum": [ "COS", "COS_CONTAINERD", "UBUNTU", "UBUNTU_CONTAINERD", "WINDOWS_LTSC", "WINDOWS_SAC" ],
						"description": "The image type used by the nodes",
						"default": "COS"
					},
					"diskSize": {
						"type": "number",
						"description": "The amount of storage in GiB provisioned on the nodes in this group",
						"multipleOf": 1,
						"default": 100,
						"minimum": 10,
						"maximum": 65536,
						"immutable": true
					},
					"preemptible": {
						"type": "boolean",
						"description": "Whether to use pre-emptible nodes (cheaper, but can and will be terminated at any time, use with care).",
						"default": false,
						"immutable": true
					},
					"labels": {
						"type": "object",
						"propertyNames": {
						  "minLength": 1,
						  "pattern": "^[a-zA-Z0-9\\-\\.\\_]+"
					    },
						"additionalProperties": { "type": "string" },
						"description": "A set of labels to help Kubernetes workloads find this group",
						"default": {},
						"immutable": true
					}
				}
			},
			"minItems": 1
		},
		"region": {
			"type": "string",
			"minLength": 1,
			"examples": ["europe-west2", "us-east1"],
			"immutable": true
		},
		"releaseChannel": {
			"type": "string",
			"description": "Follow a GKE release channel to control the auto-upgrade of your cluster - if set, auto-upgrade will be true on all node groups",
			"enum": ["REGULAR", "STABLE", "RAPID", "UNSPECIFIED"]
		},
		"version": {
			"type": "string",
			"description": "Kubernetes version - must be blank if release channel specified.",
			"pattern": "^($|-|latest|[0-9]+\\.[0-9]+($|\\.[0-9]+($|\\-gke\\.[0-9]+)))$",
			"examples": [
				"- (GKE default)", "1.15 (latest 1.15.x)", "1.15.1", "1.15.1-gke.6 (exact GKE patch version, not recommended)", "latest"
			]
		}
	},
	"allOf": [
		{
			"if": {
				"properties": {
					"inheritTeamMembers": {
						"const": true
					}
				},
				"required": ["inheritTeamMembers"]
			},
			"then": {
				"properties": {
					"defaultTeamRole": {
						"minLength": 1
					}
				},
				"required": ["defaultTeamRole"]
			},
			"else": {
			}
		},
		{
			"$comment": "Require auto-upgrade if releaseChannel not unspecified",
			"if": {
				"properties": {
					"releaseChannel": {
						"const": "UNSPECIFIED"
					}
				}
			},
			"then": {
				"properties": {
					"version": {
						"pattern": "^(-|latest|[0-9]+\\.[0-9]+($|\\.[0-9]+($|\\-gke\\.[0-9]+)))$"
					},
					"nodeGroups": {
						"properties": {
							"version": {
								"pattern": "^(-|latest|[0-9]+\\.[0-9]+($|\\.[0-9]+($|\\-gke\\.[0-9]+)))$"
							}
						}
					}
				}
			},
			"else": {
				"properties": {
					"version": {
						"const": ""
					},
					"nodeGroups": {
						"properties": {
							"enableAutoupgrade": {
								"const": true
							},
							"version": {
								"const": ""
							}
						}
					}
				}
			}
		}
	]
}
`
