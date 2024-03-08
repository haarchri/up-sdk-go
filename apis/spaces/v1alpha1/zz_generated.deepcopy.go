//go:build !ignore_autogenerated

/*
Copyright 2020 The Upbound Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMetadata) DeepCopyInto(out *ResourceMetadata) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMetadata.
func (in *ResourceMetadata) DeepCopy() *ResourceMetadata {
	if in == nil {
		return nil
	}
	out := new(ResourceMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSelector) DeepCopyInto(out *ResourceSelector) {
	*out = *in
	if in.LabelSelectors != nil {
		in, out := &in.LabelSelectors, &out.LabelSelectors
		*out = make([]v1.LabelSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSelector.
func (in *ResourceSelector) DeepCopy() *ResourceSelector {
	if in == nil {
		return nil
	}
	out := new(ResourceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStoreProvisioningFailure) DeepCopyInto(out *SecretStoreProvisioningFailure) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1beta1.SecretStoreStatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStoreProvisioningFailure.
func (in *SecretStoreProvisioningFailure) DeepCopy() *SecretStoreProvisioningFailure {
	if in == nil {
		return nil
	}
	out := new(SecretStoreProvisioningFailure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStoreProvisioningSuccess) DeepCopyInto(out *SecretStoreProvisioningSuccess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStoreProvisioningSuccess.
func (in *SecretStoreProvisioningSuccess) DeepCopy() *SecretStoreProvisioningSuccess {
	if in == nil {
		return nil
	}
	out := new(SecretStoreProvisioningSuccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedExternalSecret) DeepCopyInto(out *SharedExternalSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedExternalSecret.
func (in *SharedExternalSecret) DeepCopy() *SharedExternalSecret {
	if in == nil {
		return nil
	}
	out := new(SharedExternalSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedExternalSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedExternalSecretList) DeepCopyInto(out *SharedExternalSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SharedExternalSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedExternalSecretList.
func (in *SharedExternalSecretList) DeepCopy() *SharedExternalSecretList {
	if in == nil {
		return nil
	}
	out := new(SharedExternalSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedExternalSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedExternalSecretProvisioningFailure) DeepCopyInto(out *SharedExternalSecretProvisioningFailure) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1beta1.ClusterExternalSecretStatusCondition, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedExternalSecretProvisioningFailure.
func (in *SharedExternalSecretProvisioningFailure) DeepCopy() *SharedExternalSecretProvisioningFailure {
	if in == nil {
		return nil
	}
	out := new(SharedExternalSecretProvisioningFailure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedExternalSecretProvisioningSuccess) DeepCopyInto(out *SharedExternalSecretProvisioningSuccess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedExternalSecretProvisioningSuccess.
func (in *SharedExternalSecretProvisioningSuccess) DeepCopy() *SharedExternalSecretProvisioningSuccess {
	if in == nil {
		return nil
	}
	out := new(SharedExternalSecretProvisioningSuccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedExternalSecretSpec) DeepCopyInto(out *SharedExternalSecretSpec) {
	*out = *in
	if in.ExternalSecretMetadata != nil {
		in, out := &in.ExternalSecretMetadata, &out.ExternalSecretMetadata
		*out = new(ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	in.ControlPlaneSelector.DeepCopyInto(&out.ControlPlaneSelector)
	in.NamespaceSelector.DeepCopyInto(&out.NamespaceSelector)
	in.ExternalSecretSpec.DeepCopyInto(&out.ExternalSecretSpec)
	if in.RefreshInterval != nil {
		in, out := &in.RefreshInterval, &out.RefreshInterval
		*out = new(v1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedExternalSecretSpec.
func (in *SharedExternalSecretSpec) DeepCopy() *SharedExternalSecretSpec {
	if in == nil {
		return nil
	}
	out := new(SharedExternalSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedExternalSecretStatus) DeepCopyInto(out *SharedExternalSecretStatus) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.Failed != nil {
		in, out := &in.Failed, &out.Failed
		*out = make([]SharedExternalSecretProvisioningFailure, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Provisioned != nil {
		in, out := &in.Provisioned, &out.Provisioned
		*out = make([]SharedExternalSecretProvisioningSuccess, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedExternalSecretStatus.
func (in *SharedExternalSecretStatus) DeepCopy() *SharedExternalSecretStatus {
	if in == nil {
		return nil
	}
	out := new(SharedExternalSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedSecretStore) DeepCopyInto(out *SharedSecretStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedSecretStore.
func (in *SharedSecretStore) DeepCopy() *SharedSecretStore {
	if in == nil {
		return nil
	}
	out := new(SharedSecretStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedSecretStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedSecretStoreList) DeepCopyInto(out *SharedSecretStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SharedSecretStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedSecretStoreList.
func (in *SharedSecretStoreList) DeepCopy() *SharedSecretStoreList {
	if in == nil {
		return nil
	}
	out := new(SharedSecretStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SharedSecretStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedSecretStoreSpec) DeepCopyInto(out *SharedSecretStoreSpec) {
	*out = *in
	if in.SecretStoreMetadata != nil {
		in, out := &in.SecretStoreMetadata, &out.SecretStoreMetadata
		*out = new(ResourceMetadata)
		(*in).DeepCopyInto(*out)
	}
	in.ControlPlaneSelector.DeepCopyInto(&out.ControlPlaneSelector)
	in.NamespaceSelector.DeepCopyInto(&out.NamespaceSelector)
	in.Provider.DeepCopyInto(&out.Provider)
	if in.RetrySettings != nil {
		in, out := &in.RetrySettings, &out.RetrySettings
		*out = new(v1beta1.SecretStoreRetrySettings)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedSecretStoreSpec.
func (in *SharedSecretStoreSpec) DeepCopy() *SharedSecretStoreSpec {
	if in == nil {
		return nil
	}
	out := new(SharedSecretStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedSecretStoreStatus) DeepCopyInto(out *SharedSecretStoreStatus) {
	*out = *in
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.Failed != nil {
		in, out := &in.Failed, &out.Failed
		*out = make([]SecretStoreProvisioningFailure, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Provisioned != nil {
		in, out := &in.Provisioned, &out.Provisioned
		*out = make([]SecretStoreProvisioningSuccess, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedSecretStoreStatus.
func (in *SharedSecretStoreStatus) DeepCopy() *SharedSecretStoreStatus {
	if in == nil {
		return nil
	}
	out := new(SharedSecretStoreStatus)
	in.DeepCopyInto(out)
	return out
}