// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/splicemaahs/orderedhttp-operator/pkg/apis/orderedhttp/v1alpha1.OrderedHttp":       schema_pkg_apis_orderedhttp_v1alpha1_OrderedHttp(ref),
		"github.com/splicemaahs/orderedhttp-operator/pkg/apis/orderedhttp/v1alpha1.OrderedHttpSpec":   schema_pkg_apis_orderedhttp_v1alpha1_OrderedHttpSpec(ref),
		"github.com/splicemaahs/orderedhttp-operator/pkg/apis/orderedhttp/v1alpha1.OrderedHttpStatus": schema_pkg_apis_orderedhttp_v1alpha1_OrderedHttpStatus(ref),
	}
}

func schema_pkg_apis_orderedhttp_v1alpha1_OrderedHttp(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OrderedHttp is the Schema for the orderedhttps API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
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
							Ref: ref("github.com/splicemaahs/orderedhttp-operator/pkg/apis/orderedhttp/v1alpha1.OrderedHttpSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/splicemaahs/orderedhttp-operator/pkg/apis/orderedhttp/v1alpha1.OrderedHttpStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/splicemaahs/orderedhttp-operator/pkg/apis/orderedhttp/v1alpha1.OrderedHttpSpec", "github.com/splicemaahs/orderedhttp-operator/pkg/apis/orderedhttp/v1alpha1.OrderedHttpStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_orderedhttp_v1alpha1_OrderedHttpSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OrderedHttpSpec defines the desired state of OrderedHttp",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_orderedhttp_v1alpha1_OrderedHttpStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OrderedHttpStatus defines the observed state of OrderedHttp",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
