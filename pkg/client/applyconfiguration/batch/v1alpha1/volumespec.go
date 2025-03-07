/*
Copyright The Volcano Authors.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// VolumeSpecApplyConfiguration represents a declarative configuration of the VolumeSpec type for use
// with apply.
type VolumeSpecApplyConfiguration struct {
	MountPath       *string                       `json:"mountPath,omitempty"`
	VolumeClaimName *string                       `json:"volumeClaimName,omitempty"`
	VolumeClaim     *v1.PersistentVolumeClaimSpec `json:"volumeClaim,omitempty"`
}

// VolumeSpecApplyConfiguration constructs a declarative configuration of the VolumeSpec type for use with
// apply.
func VolumeSpec() *VolumeSpecApplyConfiguration {
	return &VolumeSpecApplyConfiguration{}
}

// WithMountPath sets the MountPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MountPath field is set to the value of the last call.
func (b *VolumeSpecApplyConfiguration) WithMountPath(value string) *VolumeSpecApplyConfiguration {
	b.MountPath = &value
	return b
}

// WithVolumeClaimName sets the VolumeClaimName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeClaimName field is set to the value of the last call.
func (b *VolumeSpecApplyConfiguration) WithVolumeClaimName(value string) *VolumeSpecApplyConfiguration {
	b.VolumeClaimName = &value
	return b
}

// WithVolumeClaim sets the VolumeClaim field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeClaim field is set to the value of the last call.
func (b *VolumeSpecApplyConfiguration) WithVolumeClaim(value v1.PersistentVolumeClaimSpec) *VolumeSpecApplyConfiguration {
	b.VolumeClaim = &value
	return b
}
