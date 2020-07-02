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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ECSInstanceSpec defines the desired state of ECSInstance
type ECSInstanceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Region ID to create an ECS Instance
	RegionID string `json:"regionID"`

	// Instance type for an ECS instance
	InstanceType string `json:"instanceType"`

	// Image ID for an ECS instance
	ImageID string `json:"imageID"`

	// Charge type for an ECS isntance
	InstanceChargeType string `json:"instanceChargeType,omitempty"`

	// Instance name
	InstanceName string `json:"instanceName,omitempty"`
}

// ECSInstanceStatus defines the observed state of ECSInstance
type ECSInstanceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Region ID to create an ECS Instance
	InstanceID string `json:"instanceID"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=ecs

// ECSInstance is the Schema for the ecsinstances API
type ECSInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ECSInstanceSpec   `json:"spec,omitempty"`
	Status ECSInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ECSInstanceList contains a list of ECSInstance
type ECSInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ECSInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ECSInstance{}, &ECSInstanceList{})
}
