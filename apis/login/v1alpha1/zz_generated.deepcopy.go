//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoginPolicy) DeepCopyInto(out *LoginPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoginPolicy.
func (in *LoginPolicy) DeepCopy() *LoginPolicy {
	if in == nil {
		return nil
	}
	out := new(LoginPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoginPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoginPolicyInitParameters) DeepCopyInto(out *LoginPolicyInitParameters) {
	*out = *in
	if in.AllowDomainDiscovery != nil {
		in, out := &in.AllowDomainDiscovery, &out.AllowDomainDiscovery
		*out = new(bool)
		**out = **in
	}
	if in.AllowExternalIdp != nil {
		in, out := &in.AllowExternalIdp, &out.AllowExternalIdp
		*out = new(bool)
		**out = **in
	}
	if in.AllowRegister != nil {
		in, out := &in.AllowRegister, &out.AllowRegister
		*out = new(bool)
		**out = **in
	}
	if in.DefaultRedirectURI != nil {
		in, out := &in.DefaultRedirectURI, &out.DefaultRedirectURI
		*out = new(string)
		**out = **in
	}
	if in.DisableLoginWithEmail != nil {
		in, out := &in.DisableLoginWithEmail, &out.DisableLoginWithEmail
		*out = new(bool)
		**out = **in
	}
	if in.DisableLoginWithPhone != nil {
		in, out := &in.DisableLoginWithPhone, &out.DisableLoginWithPhone
		*out = new(bool)
		**out = **in
	}
	if in.ExternalLoginCheckLifetime != nil {
		in, out := &in.ExternalLoginCheckLifetime, &out.ExternalLoginCheckLifetime
		*out = new(string)
		**out = **in
	}
	if in.ForceMfa != nil {
		in, out := &in.ForceMfa, &out.ForceMfa
		*out = new(bool)
		**out = **in
	}
	if in.ForceMfaLocalOnly != nil {
		in, out := &in.ForceMfaLocalOnly, &out.ForceMfaLocalOnly
		*out = new(bool)
		**out = **in
	}
	if in.HidePasswordReset != nil {
		in, out := &in.HidePasswordReset, &out.HidePasswordReset
		*out = new(bool)
		**out = **in
	}
	if in.Idps != nil {
		in, out := &in.Idps, &out.Idps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IgnoreUnknownUsernames != nil {
		in, out := &in.IgnoreUnknownUsernames, &out.IgnoreUnknownUsernames
		*out = new(bool)
		**out = **in
	}
	if in.MfaInitSkipLifetime != nil {
		in, out := &in.MfaInitSkipLifetime, &out.MfaInitSkipLifetime
		*out = new(string)
		**out = **in
	}
	if in.MultiFactorCheckLifetime != nil {
		in, out := &in.MultiFactorCheckLifetime, &out.MultiFactorCheckLifetime
		*out = new(string)
		**out = **in
	}
	if in.MultiFactors != nil {
		in, out := &in.MultiFactors, &out.MultiFactors
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PasswordCheckLifetime != nil {
		in, out := &in.PasswordCheckLifetime, &out.PasswordCheckLifetime
		*out = new(string)
		**out = **in
	}
	if in.PasswordlessType != nil {
		in, out := &in.PasswordlessType, &out.PasswordlessType
		*out = new(string)
		**out = **in
	}
	if in.SecondFactorCheckLifetime != nil {
		in, out := &in.SecondFactorCheckLifetime, &out.SecondFactorCheckLifetime
		*out = new(string)
		**out = **in
	}
	if in.SecondFactors != nil {
		in, out := &in.SecondFactors, &out.SecondFactors
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.UserLogin != nil {
		in, out := &in.UserLogin, &out.UserLogin
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoginPolicyInitParameters.
func (in *LoginPolicyInitParameters) DeepCopy() *LoginPolicyInitParameters {
	if in == nil {
		return nil
	}
	out := new(LoginPolicyInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoginPolicyList) DeepCopyInto(out *LoginPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoginPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoginPolicyList.
func (in *LoginPolicyList) DeepCopy() *LoginPolicyList {
	if in == nil {
		return nil
	}
	out := new(LoginPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoginPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoginPolicyObservation) DeepCopyInto(out *LoginPolicyObservation) {
	*out = *in
	if in.AllowDomainDiscovery != nil {
		in, out := &in.AllowDomainDiscovery, &out.AllowDomainDiscovery
		*out = new(bool)
		**out = **in
	}
	if in.AllowExternalIdp != nil {
		in, out := &in.AllowExternalIdp, &out.AllowExternalIdp
		*out = new(bool)
		**out = **in
	}
	if in.AllowRegister != nil {
		in, out := &in.AllowRegister, &out.AllowRegister
		*out = new(bool)
		**out = **in
	}
	if in.DefaultRedirectURI != nil {
		in, out := &in.DefaultRedirectURI, &out.DefaultRedirectURI
		*out = new(string)
		**out = **in
	}
	if in.DisableLoginWithEmail != nil {
		in, out := &in.DisableLoginWithEmail, &out.DisableLoginWithEmail
		*out = new(bool)
		**out = **in
	}
	if in.DisableLoginWithPhone != nil {
		in, out := &in.DisableLoginWithPhone, &out.DisableLoginWithPhone
		*out = new(bool)
		**out = **in
	}
	if in.ExternalLoginCheckLifetime != nil {
		in, out := &in.ExternalLoginCheckLifetime, &out.ExternalLoginCheckLifetime
		*out = new(string)
		**out = **in
	}
	if in.ForceMfa != nil {
		in, out := &in.ForceMfa, &out.ForceMfa
		*out = new(bool)
		**out = **in
	}
	if in.ForceMfaLocalOnly != nil {
		in, out := &in.ForceMfaLocalOnly, &out.ForceMfaLocalOnly
		*out = new(bool)
		**out = **in
	}
	if in.HidePasswordReset != nil {
		in, out := &in.HidePasswordReset, &out.HidePasswordReset
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Idps != nil {
		in, out := &in.Idps, &out.Idps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IgnoreUnknownUsernames != nil {
		in, out := &in.IgnoreUnknownUsernames, &out.IgnoreUnknownUsernames
		*out = new(bool)
		**out = **in
	}
	if in.MfaInitSkipLifetime != nil {
		in, out := &in.MfaInitSkipLifetime, &out.MfaInitSkipLifetime
		*out = new(string)
		**out = **in
	}
	if in.MultiFactorCheckLifetime != nil {
		in, out := &in.MultiFactorCheckLifetime, &out.MultiFactorCheckLifetime
		*out = new(string)
		**out = **in
	}
	if in.MultiFactors != nil {
		in, out := &in.MultiFactors, &out.MultiFactors
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.PasswordCheckLifetime != nil {
		in, out := &in.PasswordCheckLifetime, &out.PasswordCheckLifetime
		*out = new(string)
		**out = **in
	}
	if in.PasswordlessType != nil {
		in, out := &in.PasswordlessType, &out.PasswordlessType
		*out = new(string)
		**out = **in
	}
	if in.SecondFactorCheckLifetime != nil {
		in, out := &in.SecondFactorCheckLifetime, &out.SecondFactorCheckLifetime
		*out = new(string)
		**out = **in
	}
	if in.SecondFactors != nil {
		in, out := &in.SecondFactors, &out.SecondFactors
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.UserLogin != nil {
		in, out := &in.UserLogin, &out.UserLogin
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoginPolicyObservation.
func (in *LoginPolicyObservation) DeepCopy() *LoginPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(LoginPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoginPolicyParameters) DeepCopyInto(out *LoginPolicyParameters) {
	*out = *in
	if in.AllowDomainDiscovery != nil {
		in, out := &in.AllowDomainDiscovery, &out.AllowDomainDiscovery
		*out = new(bool)
		**out = **in
	}
	if in.AllowExternalIdp != nil {
		in, out := &in.AllowExternalIdp, &out.AllowExternalIdp
		*out = new(bool)
		**out = **in
	}
	if in.AllowRegister != nil {
		in, out := &in.AllowRegister, &out.AllowRegister
		*out = new(bool)
		**out = **in
	}
	if in.DefaultRedirectURI != nil {
		in, out := &in.DefaultRedirectURI, &out.DefaultRedirectURI
		*out = new(string)
		**out = **in
	}
	if in.DisableLoginWithEmail != nil {
		in, out := &in.DisableLoginWithEmail, &out.DisableLoginWithEmail
		*out = new(bool)
		**out = **in
	}
	if in.DisableLoginWithPhone != nil {
		in, out := &in.DisableLoginWithPhone, &out.DisableLoginWithPhone
		*out = new(bool)
		**out = **in
	}
	if in.ExternalLoginCheckLifetime != nil {
		in, out := &in.ExternalLoginCheckLifetime, &out.ExternalLoginCheckLifetime
		*out = new(string)
		**out = **in
	}
	if in.ForceMfa != nil {
		in, out := &in.ForceMfa, &out.ForceMfa
		*out = new(bool)
		**out = **in
	}
	if in.ForceMfaLocalOnly != nil {
		in, out := &in.ForceMfaLocalOnly, &out.ForceMfaLocalOnly
		*out = new(bool)
		**out = **in
	}
	if in.HidePasswordReset != nil {
		in, out := &in.HidePasswordReset, &out.HidePasswordReset
		*out = new(bool)
		**out = **in
	}
	if in.Idps != nil {
		in, out := &in.Idps, &out.Idps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.IgnoreUnknownUsernames != nil {
		in, out := &in.IgnoreUnknownUsernames, &out.IgnoreUnknownUsernames
		*out = new(bool)
		**out = **in
	}
	if in.MfaInitSkipLifetime != nil {
		in, out := &in.MfaInitSkipLifetime, &out.MfaInitSkipLifetime
		*out = new(string)
		**out = **in
	}
	if in.MultiFactorCheckLifetime != nil {
		in, out := &in.MultiFactorCheckLifetime, &out.MultiFactorCheckLifetime
		*out = new(string)
		**out = **in
	}
	if in.MultiFactors != nil {
		in, out := &in.MultiFactors, &out.MultiFactors
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OrgID != nil {
		in, out := &in.OrgID, &out.OrgID
		*out = new(string)
		**out = **in
	}
	if in.OrgIDRef != nil {
		in, out := &in.OrgIDRef, &out.OrgIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.OrgIDSelector != nil {
		in, out := &in.OrgIDSelector, &out.OrgIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.PasswordCheckLifetime != nil {
		in, out := &in.PasswordCheckLifetime, &out.PasswordCheckLifetime
		*out = new(string)
		**out = **in
	}
	if in.PasswordlessType != nil {
		in, out := &in.PasswordlessType, &out.PasswordlessType
		*out = new(string)
		**out = **in
	}
	if in.SecondFactorCheckLifetime != nil {
		in, out := &in.SecondFactorCheckLifetime, &out.SecondFactorCheckLifetime
		*out = new(string)
		**out = **in
	}
	if in.SecondFactors != nil {
		in, out := &in.SecondFactors, &out.SecondFactors
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.UserLogin != nil {
		in, out := &in.UserLogin, &out.UserLogin
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoginPolicyParameters.
func (in *LoginPolicyParameters) DeepCopy() *LoginPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(LoginPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoginPolicySpec) DeepCopyInto(out *LoginPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoginPolicySpec.
func (in *LoginPolicySpec) DeepCopy() *LoginPolicySpec {
	if in == nil {
		return nil
	}
	out := new(LoginPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoginPolicyStatus) DeepCopyInto(out *LoginPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoginPolicyStatus.
func (in *LoginPolicyStatus) DeepCopy() *LoginPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(LoginPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
