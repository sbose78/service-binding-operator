// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/apps/v1alpha1.ServiceBinding":       schema_pkg_apis_apps_v1alpha1_ServiceBinding(ref),
		"./pkg/apis/apps/v1alpha1.ServiceBindingStatus": schema_pkg_apis_apps_v1alpha1_ServiceBindingStatus(ref),
	}
}

func schema_pkg_apis_apps_v1alpha1_ServiceBinding(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceBinding expresses intent to bind an operator-backed service with an application workload.",
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
							Ref: ref("./pkg/apis/apps/v1alpha1.ServiceBindingSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/apps/v1alpha1.ServiceBindingStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/apps/v1alpha1.ServiceBindingSpec", "./pkg/apis/apps/v1alpha1.ServiceBindingStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_apps_v1alpha1_ServiceBindingStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceBindingStatus defines the observed state of ServiceBinding",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Description: "Conditions describes the state of the operator's reconciliation functionality.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/custom-resource-status/conditions/v1.Condition"),
									},
								},
							},
						},
					},
					"secretRef": {
						SchemaProps: spec.SchemaProps{
							Description: "Secret is the name of the intermediate secret",
							Ref:         ref("k8s.io/api/core/v1.LocalObjectReference"),
						},
					},
					"applications": {
						SchemaProps: spec.SchemaProps{
							Description: "ApplicationObjects contains all the application objects filtered by label",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/apps/v1alpha1.BoundApplication"),
									},
								},
							},
						},
					},
				},
				Required: []string{"conditions", "secretRef"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/apps/v1alpha1.BoundApplication", "github.com/openshift/custom-resource-status/conditions/v1.Condition", "k8s.io/api/core/v1.LocalObjectReference"},
	}
}
