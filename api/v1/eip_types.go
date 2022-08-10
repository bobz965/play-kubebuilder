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

// EipConditionType is a valid value for EipCondition.Type
// +enum
type EipConditionType string

// These are valid conditions of Eip.
const (
	// EipAllocated 已成功分配 IP
	EipAllocated EipConditionType = "Allocated"
	// EipBound 已被绑定给容器使用
	EipBound EipConditionType = "Bound"
	// EipReleased 已被释放
	EipReleased EipConditionType = "Released"
	// EipExists 在下层存在对应的资源
	EipExists EipConditionType = "Exists"
)

type EipCondition struct {
	// Type is the type of the condition.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
	Type EipConditionType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=EipConditionType"`
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

// EipSpec defines the desired state of Eip
type EipSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// address 手动指定期望的 IP 地址, IP 格式
	// +optional
	Address IPSet `json:"address,omitempty"`
}

// EipStatus defines the observed state of Eip
type EipStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +optional
	Address IPSet `json:"address,omitempty"`

	// Conditions represents the latest state of the object
	// +optional
	// +listType=map
	// +listMapKey=type
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []EipCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:shortName=eip
//+kubebuilder:printcolumn:name="IPv4",type=string,JSONPath=`.status.address.ipv4`
//+kubebuilder:printcolumn:name="IPv6",type=string,JSONPath=`.status.address.ipv6`
//+kubebuilder:printcolumn:name="Allocated",type=string,JSONPath=`.status.conditions[?(@.type=='Allocated')].status`
//+kubebuilder:printcolumn:name="Bound",type=string,JSONPath=`.status.conditions[?(@.type=='Bound')].status`
//+kubebuilder:printcolumn:name="Released",type=string,JSONPath=`.status.conditions[?(@.type=='Released')].status`
//+kubebuilder:printcolumn:name="Exists",type=string,JSONPath=`.status.conditions[?(@.type=='Exists')].status`

// Eip is the Schema for the eips API
type Eip struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EipSpec   `json:"spec,omitempty"`
	Status EipStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// EipList contains a list of Eip
type EipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Eip `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Eip{}, &EipList{})
}
