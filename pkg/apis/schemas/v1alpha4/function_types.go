/*
Copyright 2019 The SchemaHero Authors

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

package v1alpha4

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NotImplementedFunctionSchema struct{}

type FunctionSchema struct {
	Postgres    *PostgresqlFunctionSchema     `json:"postgres,omitempty" yaml:"postgres,omitempty"`
	Mysql       *NotImplementedFunctionSchema `json:"mysql,omitempty" yaml:"mysql,omitempty"`
	CockroachDB *NotImplementedFunctionSchema `json:"cockroachdb,omitempty" yaml:"cockroachdb,omitempty"`
	RQLite      *NotImplementedFunctionSchema `json:"rqlite,omitempty" yaml:"rqlite,omitempty"`
	SQLite      *NotImplementedFunctionSchema `json:"sqlite,omitempty" yaml:"sqlite,omitempty"`
	TimescaleDB *NotImplementedFunctionSchema `json:"timescaledb,omitempty" yaml:"timescaledb,omitempty"`
	Cassandra   *NotImplementedFunctionSchema `json:"cassandra,omitempty" yaml:"cassandra,omitempty"`
}

// FunctionSpec defines the desired state of Function
type FunctionSpec struct {
	Database         string `json:"database" yaml:"database"`
	Name             string `json:"name" yaml:"name"`
	RemoveOnDeletion bool   `json:"removeOnDeletion,omitempty" yaml:"removeOnDeletion,omitempty"`

	Schema *FunctionSchema `json:"schema,omitempty" yaml:"schema,omitempty"`
}

// FunctionStatus defines the observed state of Function
type FunctionStatus struct {
	AppliedAt int64 `json:"appliedAt,omitempty" yaml:"appliedAt,omitempty"`

	Phase string `json:"phase,omitempty" yaml:"phase,omitempty"`

	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Function is the Schema for the function API
// +kubebuilder:printcolumn:name="Namespace",type=string,JSONPath=`.metadata.namespace`,priority=1
// +kubebuilder:printcolumn:name="Function",type=string,JSONPath=`.spec.name`
// +kubebuilder:printcolumn:name="Database",type=string,JSONPath=`.spec.database`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +k8s:openapi-gen=true
type Function struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FunctionSpec   `json:"spec,omitempty"`
	Status FunctionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FunctionList contains a list of Function
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Function `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Function{}, &FunctionList{})
}
