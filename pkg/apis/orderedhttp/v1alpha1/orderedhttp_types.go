package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// OrderedHttpSpec defines the desired state of OrderedHttp
// +k8s:openapi-gen=true
type OrderedHttpSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Replicas int32 `json:"replicas"`
}

// OrderedHttpStatus defines the observed state of OrderedHttp
// +k8s:openapi-gen=true
type OrderedHttpStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	PodNames []string `json:"podnames"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OrderedHttp is the Schema for the orderedhttps API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type OrderedHttp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OrderedHttpSpec   `json:"spec,omitempty"`
	Status OrderedHttpStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OrderedHttpList contains a list of OrderedHttp
type OrderedHttpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrderedHttp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OrderedHttp{}, &OrderedHttpList{})
}
