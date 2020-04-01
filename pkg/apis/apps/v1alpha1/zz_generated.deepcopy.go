// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	conditionsv1 "github.com/openshift/custom-resource-status/conditions/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSelector) DeepCopyInto(out *ApplicationSelector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	out.GroupVersionResource = in.GroupVersionResource
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSelector.
func (in *ApplicationSelector) DeepCopy() *ApplicationSelector {
	if in == nil {
		return nil
	}
	out := new(ApplicationSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackingServiceSelector) DeepCopyInto(out *BackingServiceSelector) {
	*out = *in
	out.GroupVersionKind = in.GroupVersionKind
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackingServiceSelector.
func (in *BackingServiceSelector) DeepCopy() *BackingServiceSelector {
	if in == nil {
		return nil
	}
	out := new(BackingServiceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindingData) DeepCopyInto(out *BindingData) {
	*out = *in
	in.TypedLocalObjectReference.DeepCopyInto(&out.TypedLocalObjectReference)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindingData.
func (in *BindingData) DeepCopy() *BindingData {
	if in == nil {
		return nil
	}
	out := new(BindingData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BoundApplication) DeepCopyInto(out *BoundApplication) {
	*out = *in
	out.GroupVersionKind = in.GroupVersionKind
	out.LocalObjectReference = in.LocalObjectReference
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BoundApplication.
func (in *BoundApplication) DeepCopy() *BoundApplication {
	if in == nil {
		return nil
	}
	out := new(BoundApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingRequest) DeepCopyInto(out *ServiceBindingRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingRequest.
func (in *ServiceBindingRequest) DeepCopy() *ServiceBindingRequest {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceBindingRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingRequestList) DeepCopyInto(out *ServiceBindingRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceBindingRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingRequestList.
func (in *ServiceBindingRequestList) DeepCopy() *ServiceBindingRequestList {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceBindingRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingRequestSpec) DeepCopyInto(out *ServiceBindingRequestSpec) {
	*out = *in
	if in.Binding != nil {
		in, out := &in.Binding, &out.Binding
		*out = new(BindingData)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomEnvVar != nil {
		in, out := &in.CustomEnvVar, &out.CustomEnvVar
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BackingServiceSelector != nil {
		in, out := &in.BackingServiceSelector, &out.BackingServiceSelector
		*out = new(BackingServiceSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.BackingServiceSelectors != nil {
		in, out := &in.BackingServiceSelectors, &out.BackingServiceSelectors
		*out = new([]BackingServiceSelector)
		if **in != nil {
			in, out := *in, *out
			*out = make([]BackingServiceSelector, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	in.ApplicationSelector.DeepCopyInto(&out.ApplicationSelector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingRequestSpec.
func (in *ServiceBindingRequestSpec) DeepCopy() *ServiceBindingRequestSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceBindingRequestStatus) DeepCopyInto(out *ServiceBindingRequestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditionsv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.BindingData.DeepCopyInto(&out.BindingData)
	if in.Applications != nil {
		in, out := &in.Applications, &out.Applications
		*out = make([]BoundApplication, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceBindingRequestStatus.
func (in *ServiceBindingRequestStatus) DeepCopy() *ServiceBindingRequestStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceBindingRequestStatus)
	in.DeepCopyInto(out)
	return out
}
