package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type OpenDjList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []OpenDj `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type OpenDj struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              OpenDjSpec   `json:"spec"`
	Status            OpenDjStatus `json:"status,omitempty"`
}

type OpenDjSpec struct {
	RootPasswordSecret *corev1.LocalObjectReference `json:"rootPasswordSecret,omitempty"`
	LdifFileConfig *corev1.LocalObjectReference `json:"ldifFileConfig,omitempty"`
	Members int32 `json:"members,omitempty"`
	Version string `json:"version,omitempty"`
	Config LdapConfig `json:"config,omitempty"`
}

type LdapConfig struct {
	Port int32 `json:"port,omitempty"`
	LdapPort int32 `json:"ldapPort,omitempty"`
	BaseDn string `json:"baseDn,omitempty"`
	RootUserDn string `json:"rootUserDn,omitempty"`
	HostName string `json:"hostName,omitempty"`
}

type OpenDjStatus struct {
	Nodes []string `json:"nodes"`
	Active string `json:"active"`
}
