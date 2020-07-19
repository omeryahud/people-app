/*


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
	v1 "github.com/openshift/custom-resource-status/conditions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PeopleAppSpec defines the desired state of PeopleApp
type PeopleAppSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Defines desired state for frontend pods
	FrontendSpec FrontendSpec `json:"frontendSpec"`

	// Defines desired state for backend pods
	BackendSpec BackendSpec `json:"backendSpec"`

	// Defines desired state for database pods
	DatabaseSpec DatabaseSpec `json:"databaseSpec"`
}

// FrontendSpec defines the desired state for frontend deployment
type FrontendSpec struct {
	// Replicas is the replica count for pods
	Replicas *int32 `json:"replicas,omitempty"`

	// HttpPort is the port pods should listen on
	HttpPort int32 `json:"httpPort"`
}

// BackendSpec defines the desired state for backend deployment
type BackendSpec struct {
	// Replicas is the replica count for pods
	Replicas *int32 `json:"replicas,omitempty"`

	// HttpPort is the port pods should listen on
	HttpPort int32 `json:"httpPort"`
}

// DatabaseSpec defines the desired state for database deployment
type DatabaseSpec struct {
	// Replicas is the replica count for  pods
	Replicas *int32 `json:"replicas,omitempty"`

	// HttpPort is the port pods should listen on
	HttpPort int32 `json:"httpPort"`
}

// PeopleAppStatus defines the observed state of PeopleApp
type PeopleAppStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Conditions hold the condition of the deployment
	Conditions []v1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// PeopleApp is the Schema for the peopleapps API
type PeopleApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PeopleAppSpec   `json:"spec,omitempty"`
	Status PeopleAppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PeopleAppList contains a list of PeopleApp
type PeopleAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PeopleApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PeopleApp{}, &PeopleAppList{})
}
