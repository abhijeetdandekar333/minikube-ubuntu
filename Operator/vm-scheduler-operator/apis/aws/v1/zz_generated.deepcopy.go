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
func (in *AWSVMSchedulerStart) DeepCopyInto(out *AWSVMSchedulerStart) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVMSchedulerStart.
func (in *AWSVMSchedulerStart) DeepCopy() *AWSVMSchedulerStart {
	if in == nil {
		return nil
	}
	out := new(AWSVMSchedulerStart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSVMSchedulerStart) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVMSchedulerStartList) DeepCopyInto(out *AWSVMSchedulerStartList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSVMSchedulerStart, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVMSchedulerStartList.
func (in *AWSVMSchedulerStartList) DeepCopy() *AWSVMSchedulerStartList {
	if in == nil {
		return nil
	}
	out := new(AWSVMSchedulerStartList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSVMSchedulerStartList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVMSchedulerStartSpec) DeepCopyInto(out *AWSVMSchedulerStartSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVMSchedulerStartSpec.
func (in *AWSVMSchedulerStartSpec) DeepCopy() *AWSVMSchedulerStartSpec {
	if in == nil {
		return nil
	}
	out := new(AWSVMSchedulerStartSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSVMSchedulerStartStatus) DeepCopyInto(out *AWSVMSchedulerStartStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSVMSchedulerStartStatus.
func (in *AWSVMSchedulerStartStatus) DeepCopy() *AWSVMSchedulerStartStatus {
	if in == nil {
		return nil
	}
	out := new(AWSVMSchedulerStartStatus)
	in.DeepCopyInto(out)
	return out
}
