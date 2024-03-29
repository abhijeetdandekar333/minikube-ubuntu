//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureVMSchedulerStart) DeepCopyInto(out *AzureVMSchedulerStart) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureVMSchedulerStart.
func (in *AzureVMSchedulerStart) DeepCopy() *AzureVMSchedulerStart {
	if in == nil {
		return nil
	}
	out := new(AzureVMSchedulerStart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureVMSchedulerStart) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureVMSchedulerStartList) DeepCopyInto(out *AzureVMSchedulerStartList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureVMSchedulerStart, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureVMSchedulerStartList.
func (in *AzureVMSchedulerStartList) DeepCopy() *AzureVMSchedulerStartList {
	if in == nil {
		return nil
	}
	out := new(AzureVMSchedulerStartList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureVMSchedulerStartList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureVMSchedulerStartSpec) DeepCopyInto(out *AzureVMSchedulerStartSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureVMSchedulerStartSpec.
func (in *AzureVMSchedulerStartSpec) DeepCopy() *AzureVMSchedulerStartSpec {
	if in == nil {
		return nil
	}
	out := new(AzureVMSchedulerStartSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureVMSchedulerStartStatus) DeepCopyInto(out *AzureVMSchedulerStartStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureVMSchedulerStartStatus.
func (in *AzureVMSchedulerStartStatus) DeepCopy() *AzureVMSchedulerStartStatus {
	if in == nil {
		return nil
	}
	out := new(AzureVMSchedulerStartStatus)
	in.DeepCopyInto(out)
	return out
}
