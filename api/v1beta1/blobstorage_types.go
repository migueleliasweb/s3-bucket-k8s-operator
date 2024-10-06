package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BlobstorageSpec defines the desired state of Blobstorage
type BlobstorageSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Blobstorage. Edit blobstorage_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// BlobstorageStatus defines the observed state of Blobstorage
type BlobstorageStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// Blobstorage is the Schema for the blobstorages API
type Blobstorage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BlobstorageSpec   `json:"spec,omitempty"`
	Status BlobstorageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BlobstorageList contains a list of Blobstorage
type BlobstorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Blobstorage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Blobstorage{}, &BlobstorageList{})
}
