/*
Copyright 2021.

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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TemplateSecretSpec defines the desired state of TemplateSecret
type TemplateSecretSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Template     TemplateSource `json:"template"`
	Replacements []ReplacementSpec
}

// TemplateSource defines a template from something else
type TemplateSource struct {
	// Use an inline map as a source for the template
	// +optional
	Inline map[string]string `json:"inline,omitempty"`
	// Use a secret as a source for the template
	// +optional
	Secret corev1.SecretProjection `json:"secret,omitempty"`
	// Use a configmap as a source for the tamplate
	// +optional
	ConfigMap corev1.ConfigMapProjection `json:"configMap,omitempty"`
}

// ReplacementSpec defines a match string and corresponding replacement to make in the template
type ReplacementSpec struct {
	Match       string            `json:"match"`
	Replacement ReplacementSource `json:"replacement"`
}

// ReplacementSource points to a value to use as a replacement in a ReplacementSpec
type ReplacementSource struct {
	// Selects a key of a ConfigMap.
	// +optional
	ConfigMapKeyRef *corev1.ConfigMapKeySelector `json:"configMapKeyRef,omitempty"`
	// Selects a key of a secret in the pod's namespace
	// +optional
	SecretKeyRef *corev1.SecretKeySelector `json:"secretKeyRef,omitempty"`
}

// TemplateSecretStatus defines the observed state of TemplateSecret
type TemplateSecretStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TemplateSecret is the Schema for the templatesecrets API
type TemplateSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TemplateSecretSpec   `json:"spec,omitempty"`
	Status TemplateSecretStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TemplateSecretList contains a list of TemplateSecret
type TemplateSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TemplateSecret `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TemplateSecret{}, &TemplateSecretList{})
}
