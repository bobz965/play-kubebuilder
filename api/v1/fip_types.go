/*
Copyright 2022.

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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FipConditionType is a valid value for FipCondition.Type
// +enum
type FipConditionType string

// These are valid conditions of Fip.
const (
	// L2 FIP 只有两种状态，Exists 表示L1 存在fip, Released 表示L1 不存在fip
	// FipReleased 已被释放
	FipReleased FipConditionType = "Released"
	// FipExists 在下层存在对应的资源
	FipExists FipConditionType = "Exists"
)

type FipCondition struct {
	// Type is the type of the condition.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	Type FipConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=FipConditionType"`
	// Status is the status of the condition.
	// Can be True, False, Unknown.
	Status corev1.ConditionStatus `json:"status" protobuf:"bytes,2,opt,name=status,casttype=ConditionStatus"`
	// Last time we probed the condition.
	// +optional
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty" protobuf:"bytes,3,opt,name=lastProbeTime"`
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty" protobuf:"bytes,4,opt,name=lastTransitionTime"`
	// Unique, one-word, CamelCase reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,5,opt,name=reason"`
	// Human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty" protobuf:"bytes,6,opt,name=message"`
}

// FipSpec defines the desired state of Fip
type FipSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	EIP        string `json:"eip"`
	InternalIp string `json:"internalIp"`
}

// FipStatus defines the observed state of Fip
type FipStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Conditions represents the latest state of the object
	// +optional
	// +listType=map
	// +listMapKey=type
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []FipCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:shortName=fip
//+kubebuilder:printcolumn:name="Released",type=string,JSONPath=`.status.conditions[?(@.type=='Released')].status`
//+kubebuilder:printcolumn:name="Exists",type=string,JSONPath=`.status.conditions[?(@.type=='Exists')].status`

// Fip is the Schema for the fips API
type Fip struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FipSpec   `json:"spec,omitempty"`
	Status FipStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// FipList contains a list of Fip
type FipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Fip `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Fip{}, &FipList{})
}
