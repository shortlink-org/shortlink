/*
Copyright 2022 Viktor Login.

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

type ShortDBHeartBeat struct {
	// Heartbeat interval in milliseconds
	Interval int `json:"interval,omitempty"`
}

type ShortDBKeepAlive struct {
	// Keep alive is enabled
	Enabled bool `json:"enabled,omitempty"`
	// Keep alive messages interval in seconds
	Interval int `json:"interval,omitempty"`
}

// ShortDBSpec defines the desired state of ShortDB
type ShortDBSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Deployments count
	Deployments int `json:"deployments,omitempty"`
	// Maximum memory limit in GB (for example, 4)
	MaxMemory *int `json:"maxMemory,omitempty"`
	// Maximum cpu limit (100 = full power)
	MaxCPU *int `json:"maxCPU,omitempty"`
	// Heartbeat configuration
	HeartBeat *ShortDBHeartBeat `json:"heartbeat,omitempty"`
	// Keepalive configuration
	Keepalive *ShortDBKeepAlive `json:"keepalive,omitempty"`
}

// ShortDBStatus defines the observed state of ShortDB
type ShortDBStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// All nodes are prepared and ready
	Deployed bool `json:"deployed"`
	// How many nodes isn't available
	BrokenNodes int `json:"brokenNodes"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ShortDB is the Schema for the shortdbs API
type ShortDB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ShortDBSpec   `json:"spec,omitempty"`
	Status ShortDBStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ShortDBList contains a list of ShortDB
type ShortDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ShortDB `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ShortDB{}, &ShortDBList{})
}
