// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha1.S3Bucket":     schema_pkg_apis_managed_v1alpha1_S3Bucket(ref),
		"github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha1.Velero":       schema_pkg_apis_managed_v1alpha1_Velero(ref),
		"github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha1.VeleroSpec":   schema_pkg_apis_managed_v1alpha1_VeleroSpec(ref),
		"github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha1.VeleroStatus": schema_pkg_apis_managed_v1alpha1_VeleroStatus(ref),
	}
}

func schema_pkg_apis_managed_v1alpha1_S3Bucket(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "S3Bucket defines the observed state of Velero",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is the name of the S3 bucket created to store Velero backup details",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"provisioned": {
						SchemaProps: spec.SchemaProps{
							Description: "Provisioned is true once the bucket has been initially provisioned.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"lastSyncTimestamp": {
						SchemaProps: spec.SchemaProps{
							Description: "LastSyncTimestamp is the time that the bucket policy was last synced.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
				},
				Required: []string{"provisioned"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_managed_v1alpha1_Velero(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Velero is the Schema for the veleros API",
				Type:        []string{"object"},
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
							Ref: ref("github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha1.VeleroSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha1.VeleroStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha1.VeleroSpec", "github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha1.VeleroStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_managed_v1alpha1_VeleroSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VeleroSpec defines the desired state of Velero",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_managed_v1alpha1_VeleroStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VeleroStatus defines the observed state of Velero",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"s3Bucket": {
						SchemaProps: spec.SchemaProps{
							Description: "S3Bucket contains details of the S3 storage bucket for backups",
							Ref:         ref("github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha1.S3Bucket"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/managed-velero-operator/pkg/apis/managed/v1alpha1.S3Bucket"},
	}
}
