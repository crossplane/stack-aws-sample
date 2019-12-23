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
	"github.com/crossplaneio/crossplane-runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MinimalAWSSpec defines the desired state of MinimalAWS
type MinimalAWSSpec struct {
	// CredentialsSecretRef refers to the secret and its key that contains
	// the required credentials to connect to AWS.
	CredentialsSecretRef v1alpha1.SecretKeySelector `json:"credentialsSecretRef"`

	// Region of the resources that will be deployed.
	Region string `json:"region"`
}

// MinimalAWSStatus defines the observed state of MinimalAWS
type MinimalAWSStatus struct {
	v1alpha1.ConditionedStatus `json:",inline"`
}

func (mg *MinimalAWS) GetCondition(ct v1alpha1.ConditionType) v1alpha1.Condition {
	return mg.Status.GetCondition(ct)
}

func (mg *MinimalAWS) SetConditions(c ...v1alpha1.Condition) {
	mg.Status.SetConditions(c...)
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:subresource:status

// MinimalAWS is the Schema for the minimalaws API
type MinimalAWS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MinimalAWSSpec   `json:"spec,omitempty"`
	Status MinimalAWSStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MinimalAWSList contains a list of MinimalAWS
type MinimalAWSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MinimalAWS `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MinimalAWS{}, &MinimalAWSList{})
}
