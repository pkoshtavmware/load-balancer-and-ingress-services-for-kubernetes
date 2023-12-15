/*
* Copyright 2022-2023 VMware, Inc.
* All Rights Reserved.
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*   http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// Code generated by a tool; DO NOT EDIT.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type L4RuleSpec struct {
	AnalyticsPolicy          *AnalyticsPolicy     `json:"analyticsPolicy,omitempty"`
	AnalyticsProfileRef      *string              `json:"analyticsProfileRef,omitempty"`
	ApplicationProfileRef    *string              `json:"applicationProfileRef,omitempty"`
	BackendProperties        []*BackendProperties `json:"backendProperties,omitempty"`
	Services                 []*Service           `json:"listenerProperties,omitempty"`
	LoadBalancerIP           *string              `json:"loadBalancerIP,omitempty"`
	NetworkProfileRef        *string              `json:"networkProfileRef,omitempty"`
	NetworkSecurityPolicyRef *string              `json:"networkSecurityPolicyRef,omitempty"`
	PerformanceLimits        *PerformanceLimits   `json:"performanceLimits,omitempty"`
	SecurityPolicyRef        *string              `json:"securityPolicyRef,omitempty"`
	SslKeyAndCertificateRefs []string             `json:"sslKeyAndCertificateRefs,omitempty"`
	SslProfileRef            *string              `json:"sslProfileRef,omitempty"`
	VsDatascriptRefs         []string             `json:"vsDatascriptRefs,omitempty"`
}

type L4RuleStatus struct {
	Status string `json:"status,omitempty"`
	Error  string `json:"error"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type L4Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              L4RuleSpec   `json:"spec,omitempty"`
	Status            L4RuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type L4RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []L4Rule `json:"items"`
}