/*
Copyright 2022 hef.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ServerSpec defines the desired state of Server
type ServerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Number of CPUs [1 - 10]
	Cpu int `json:"cpu,omitempty"`
	// Enable Encryption [true|false]
	Encryption bool `json:"encryption,omitempty"`
	// Enable High Availability [true|false]
	Ha bool `json:"ha,omitempty"`
	// Operating System ["CentOS 7.9 64bit", "CentOS 8.3 64bit", "Debian 9.13 64Bit", "FreeBSD 12.2 64bit", "Ubuntu 18.04 LTS 64bit"]
	Os string `json:"os,omitempty"`
	//Ram in MB, valid values are [512, 1024, 1536, 2048, 2560, 3072, 4096, 5120, 6144, 7168, 8192]
	Ram int `json:"ram,omitempty"`
	// Storage to allocate, in GB, valid values are multiples of 10
	Storage int `json:"storage,omitempty"`
}

// ServerStatus defines the observed state of Server
type ServerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Ip   string `json:"ip,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Server is the Schema for the servers API
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServerSpec   `json:"spec,omitempty"`
	Status ServerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ServerList contains a list of Server
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}
