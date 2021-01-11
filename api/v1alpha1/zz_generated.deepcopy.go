// +build !ignore_autogenerated

/*
Copyright 2021.

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
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchSpec) DeepCopyInto(out *MatchSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchSpec.
func (in *MatchSpec) DeepCopy() *MatchSpec {
	if in == nil {
		return nil
	}
	out := new(MatchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplacementSource) DeepCopyInto(out *ReplacementSource) {
	*out = *in
	if in.ConfigMapKeyRef != nil {
		in, out := &in.ConfigMapKeyRef, &out.ConfigMapKeyRef
		*out = new(v1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(v1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplacementSource.
func (in *ReplacementSource) DeepCopy() *ReplacementSource {
	if in == nil {
		return nil
	}
	out := new(ReplacementSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplacementSpec) DeepCopyInto(out *ReplacementSpec) {
	*out = *in
	out.Match = in.Match
	in.Replacement.DeepCopyInto(&out.Replacement)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplacementSpec.
func (in *ReplacementSpec) DeepCopy() *ReplacementSpec {
	if in == nil {
		return nil
	}
	out := new(ReplacementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSecret) DeepCopyInto(out *TemplateSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSecret.
func (in *TemplateSecret) DeepCopy() *TemplateSecret {
	if in == nil {
		return nil
	}
	out := new(TemplateSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemplateSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSecretList) DeepCopyInto(out *TemplateSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TemplateSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSecretList.
func (in *TemplateSecretList) DeepCopy() *TemplateSecretList {
	if in == nil {
		return nil
	}
	out := new(TemplateSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemplateSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSecretSpec) DeepCopyInto(out *TemplateSecretSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	if in.Replacements != nil {
		in, out := &in.Replacements, &out.Replacements
		*out = make([]ReplacementSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSecretSpec.
func (in *TemplateSecretSpec) DeepCopy() *TemplateSecretSpec {
	if in == nil {
		return nil
	}
	out := new(TemplateSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSecretStatus) DeepCopyInto(out *TemplateSecretStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSecretStatus.
func (in *TemplateSecretStatus) DeepCopy() *TemplateSecretStatus {
	if in == nil {
		return nil
	}
	out := new(TemplateSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSource) DeepCopyInto(out *TemplateSource) {
	*out = *in
	if in.Inline != nil {
		in, out := &in.Inline, &out.Inline
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Secret.DeepCopyInto(&out.Secret)
	in.ConfigMap.DeepCopyInto(&out.ConfigMap)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSource.
func (in *TemplateSource) DeepCopy() *TemplateSource {
	if in == nil {
		return nil
	}
	out := new(TemplateSource)
	in.DeepCopyInto(out)
	return out
}
