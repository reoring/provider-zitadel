//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdpGithub) DeepCopyInto(out *IdpGithub) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdpGithub.
func (in *IdpGithub) DeepCopy() *IdpGithub {
	if in == nil {
		return nil
	}
	out := new(IdpGithub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdpGithub) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdpGithubInitParameters) DeepCopyInto(out *IdpGithubInitParameters) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.IsAutoCreation != nil {
		in, out := &in.IsAutoCreation, &out.IsAutoCreation
		*out = new(bool)
		**out = **in
	}
	if in.IsAutoUpdate != nil {
		in, out := &in.IsAutoUpdate, &out.IsAutoUpdate
		*out = new(bool)
		**out = **in
	}
	if in.IsCreationAllowed != nil {
		in, out := &in.IsCreationAllowed, &out.IsCreationAllowed
		*out = new(bool)
		**out = **in
	}
	if in.IsLinkingAllowed != nil {
		in, out := &in.IsLinkingAllowed, &out.IsLinkingAllowed
		*out = new(bool)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdpGithubInitParameters.
func (in *IdpGithubInitParameters) DeepCopy() *IdpGithubInitParameters {
	if in == nil {
		return nil
	}
	out := new(IdpGithubInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdpGithubList) DeepCopyInto(out *IdpGithubList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IdpGithub, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdpGithubList.
func (in *IdpGithubList) DeepCopy() *IdpGithubList {
	if in == nil {
		return nil
	}
	out := new(IdpGithubList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdpGithubList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdpGithubObservation) DeepCopyInto(out *IdpGithubObservation) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IsAutoCreation != nil {
		in, out := &in.IsAutoCreation, &out.IsAutoCreation
		*out = new(bool)
		**out = **in
	}
	if in.IsAutoUpdate != nil {
		in, out := &in.IsAutoUpdate, &out.IsAutoUpdate
		*out = new(bool)
		**out = **in
	}
	if in.IsCreationAllowed != nil {
		in, out := &in.IsCreationAllowed, &out.IsCreationAllowed
		*out = new(bool)
		**out = **in
	}
	if in.IsLinkingAllowed != nil {
		in, out := &in.IsLinkingAllowed, &out.IsLinkingAllowed
		*out = new(bool)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdpGithubObservation.
func (in *IdpGithubObservation) DeepCopy() *IdpGithubObservation {
	if in == nil {
		return nil
	}
	out := new(IdpGithubObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdpGithubParameters) DeepCopyInto(out *IdpGithubParameters) {
	*out = *in
	if in.ClientID != nil {
		in, out := &in.ClientID, &out.ClientID
		*out = new(string)
		**out = **in
	}
	out.ClientSecretSecretRef = in.ClientSecretSecretRef
	if in.IsAutoCreation != nil {
		in, out := &in.IsAutoCreation, &out.IsAutoCreation
		*out = new(bool)
		**out = **in
	}
	if in.IsAutoUpdate != nil {
		in, out := &in.IsAutoUpdate, &out.IsAutoUpdate
		*out = new(bool)
		**out = **in
	}
	if in.IsCreationAllowed != nil {
		in, out := &in.IsCreationAllowed, &out.IsCreationAllowed
		*out = new(bool)
		**out = **in
	}
	if in.IsLinkingAllowed != nil {
		in, out := &in.IsLinkingAllowed, &out.IsLinkingAllowed
		*out = new(bool)
		**out = **in
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdpGithubParameters.
func (in *IdpGithubParameters) DeepCopy() *IdpGithubParameters {
	if in == nil {
		return nil
	}
	out := new(IdpGithubParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdpGithubSpec) DeepCopyInto(out *IdpGithubSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdpGithubSpec.
func (in *IdpGithubSpec) DeepCopy() *IdpGithubSpec {
	if in == nil {
		return nil
	}
	out := new(IdpGithubSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdpGithubStatus) DeepCopyInto(out *IdpGithubStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdpGithubStatus.
func (in *IdpGithubStatus) DeepCopy() *IdpGithubStatus {
	if in == nil {
		return nil
	}
	out := new(IdpGithubStatus)
	in.DeepCopyInto(out)
	return out
}
