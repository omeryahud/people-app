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
	conditionsv1 "github.com/openshift/custom-resource-status/conditions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PeopleAppSpec defines the desired state of PeopleApp
type PeopleAppSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// FrontendConfig holds configuration for the frontend deployment
	FrontendConfig FrontendConfig `json:"frontendConfig"`

	// BackendConfig holds configuration for the backend deployment
	BackendConfig BackendConfig `json:"backendConfig"`

	// DatabseConfig holds configuration for the database config
	DatabaseConfig DatabaseConfig `json:"databaseConfig"`
}

// FrontendConfig defines the configurations of the frontend deployment
type FrontendConfig struct {
	// Image is the image used for frontend pods
	Image string `json:"image"`
	// Replicas is the desired amount of frontend pods
	Replicas int32 `json:"replicas"`
}

// BackendConfig defines the configurations of the backend deployment
type BackendConfig struct {
	// Image is the image used for backend pods
	Image string `json:"image"`
	// Replicas is the desired amount of backend pods
	Replicas int32 `json:"replicas"`
}

// DatabaseConfig defines the configurations of the database deployment
type DatabaseConfig struct {
	// Image is the image used for database pods
	Image string `json:"image"`
	// Replicas is the desired amount of database pods
	Replicas int32 `json:"replicas"`
}

// PeopleAppStatus defines the observed state of PeopleApp
type PeopleAppStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []conditionsv1.Condition `json:"conditions,omitempty"`
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
