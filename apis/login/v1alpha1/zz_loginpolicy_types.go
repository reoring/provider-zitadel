// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type LoginPolicyInitParameters struct {

	// (Boolean) if set to true, the suffix (@domain.com) of an unknown username input on the login screen will be matched against the org domains and will redirect to the registration of that organisation on success.
	// if set to true, the suffix (@domain.com) of an unknown username input on the login screen will be matched against the org domains and will redirect to the registration of that organisation on success.
	AllowDomainDiscovery *bool `json:"allowDomainDiscovery,omitempty" tf:"allow_domain_discovery,omitempty"`

	// (Boolean) defines if a user is allowed to add a defined identity provider. E.g. Google auth
	// defines if a user is allowed to add a defined identity provider. E.g. Google auth
	AllowExternalIdp *bool `json:"allowExternalIdp,omitempty" tf:"allow_external_idp,omitempty"`

	// (Boolean) defines if a person is allowed to register a user on this organisation
	// defines if a person is allowed to register a user on this organisation
	AllowRegister *bool `json:"allowRegister,omitempty" tf:"allow_register,omitempty"`

	// (String) defines where the user will be redirected to if the login is started without app context (e.g. from mail)
	// defines where the user will be redirected to if the login is started without app context (e.g. from mail)
	DefaultRedirectURI *string `json:"defaultRedirectUri,omitempty" tf:"default_redirect_uri,omitempty"`

	// (Boolean) defines if user can additionally (to the loginname) be identified by their verified email address
	// defines if user can additionally (to the loginname) be identified by their verified email address
	DisableLoginWithEmail *bool `json:"disableLoginWithEmail,omitempty" tf:"disable_login_with_email,omitempty"`

	// (Boolean) defines if user can additionally (to the loginname) be identified by their verified phone number
	// defines if user can additionally (to the loginname) be identified by their verified phone number
	DisableLoginWithPhone *bool `json:"disableLoginWithPhone,omitempty" tf:"disable_login_with_phone,omitempty"`

	// (String)
	ExternalLoginCheckLifetime *string `json:"externalLoginCheckLifetime,omitempty" tf:"external_login_check_lifetime,omitempty"`

	// (Boolean) defines if a user MUST use a multi factor to log in
	// defines if a user MUST use a multi factor to log in
	ForceMfa *bool `json:"forceMfa,omitempty" tf:"force_mfa,omitempty"`

	// (Boolean) if activated, ZITADEL only enforces MFA on local authentications. On authentications through MFA, ZITADEL won't prompt for MFA.
	// if activated, ZITADEL only enforces MFA on local authentications. On authentications through MFA, ZITADEL won't prompt for MFA.
	ForceMfaLocalOnly *bool `json:"forceMfaLocalOnly,omitempty" tf:"force_mfa_local_only,omitempty"`

	// (Boolean) defines if password reset link should be shown in the login screen
	// defines if password reset link should be shown in the login screen
	HidePasswordReset *bool `json:"hidePasswordReset,omitempty" tf:"hide_password_reset,omitempty"`

	// (Boolean) defines if unknown username on login screen directly return an error or always display the password screen
	// defines if unknown username on login screen directly return an error or always display the password screen
	IgnoreUnknownUsernames *bool `json:"ignoreUnknownUsernames,omitempty" tf:"ignore_unknown_usernames,omitempty"`

	// (String)
	MfaInitSkipLifetime *string `json:"mfaInitSkipLifetime,omitempty" tf:"mfa_init_skip_lifetime,omitempty"`

	// (String)
	MultiFactorCheckLifetime *string `json:"multiFactorCheckLifetime,omitempty" tf:"multi_factor_check_lifetime,omitempty"`

	// (Set of String) allowed multi factors
	// allowed multi factors
	MultiFactors []*string `json:"multiFactors,omitempty" tf:"multi_factors,omitempty"`

	// (String)
	PasswordCheckLifetime *string `json:"passwordCheckLifetime,omitempty" tf:"password_check_lifetime,omitempty"`

	// (String) defines if passwordless is allowed for users
	// defines if passwordless is allowed for users
	PasswordlessType *string `json:"passwordlessType,omitempty" tf:"passwordless_type,omitempty"`

	// (String)
	SecondFactorCheckLifetime *string `json:"secondFactorCheckLifetime,omitempty" tf:"second_factor_check_lifetime,omitempty"`

	// (Set of String) allowed second factors
	// allowed second factors
	SecondFactors []*string `json:"secondFactors,omitempty" tf:"second_factors,omitempty"`

	// (Boolean) defines if a user is allowed to login with his username and password
	// defines if a user is allowed to login with his username and password
	UserLogin *bool `json:"userLogin,omitempty" tf:"user_login,omitempty"`
}

type LoginPolicyObservation struct {

	// (Boolean) if set to true, the suffix (@domain.com) of an unknown username input on the login screen will be matched against the org domains and will redirect to the registration of that organisation on success.
	// if set to true, the suffix (@domain.com) of an unknown username input on the login screen will be matched against the org domains and will redirect to the registration of that organisation on success.
	AllowDomainDiscovery *bool `json:"allowDomainDiscovery,omitempty" tf:"allow_domain_discovery,omitempty"`

	// (Boolean) defines if a user is allowed to add a defined identity provider. E.g. Google auth
	// defines if a user is allowed to add a defined identity provider. E.g. Google auth
	AllowExternalIdp *bool `json:"allowExternalIdp,omitempty" tf:"allow_external_idp,omitempty"`

	// (Boolean) defines if a person is allowed to register a user on this organisation
	// defines if a person is allowed to register a user on this organisation
	AllowRegister *bool `json:"allowRegister,omitempty" tf:"allow_register,omitempty"`

	// (String) defines where the user will be redirected to if the login is started without app context (e.g. from mail)
	// defines where the user will be redirected to if the login is started without app context (e.g. from mail)
	DefaultRedirectURI *string `json:"defaultRedirectUri,omitempty" tf:"default_redirect_uri,omitempty"`

	// (Boolean) defines if user can additionally (to the loginname) be identified by their verified email address
	// defines if user can additionally (to the loginname) be identified by their verified email address
	DisableLoginWithEmail *bool `json:"disableLoginWithEmail,omitempty" tf:"disable_login_with_email,omitempty"`

	// (Boolean) defines if user can additionally (to the loginname) be identified by their verified phone number
	// defines if user can additionally (to the loginname) be identified by their verified phone number
	DisableLoginWithPhone *bool `json:"disableLoginWithPhone,omitempty" tf:"disable_login_with_phone,omitempty"`

	// (String)
	ExternalLoginCheckLifetime *string `json:"externalLoginCheckLifetime,omitempty" tf:"external_login_check_lifetime,omitempty"`

	// (Boolean) defines if a user MUST use a multi factor to log in
	// defines if a user MUST use a multi factor to log in
	ForceMfa *bool `json:"forceMfa,omitempty" tf:"force_mfa,omitempty"`

	// (Boolean) if activated, ZITADEL only enforces MFA on local authentications. On authentications through MFA, ZITADEL won't prompt for MFA.
	// if activated, ZITADEL only enforces MFA on local authentications. On authentications through MFA, ZITADEL won't prompt for MFA.
	ForceMfaLocalOnly *bool `json:"forceMfaLocalOnly,omitempty" tf:"force_mfa_local_only,omitempty"`

	// (Boolean) defines if password reset link should be shown in the login screen
	// defines if password reset link should be shown in the login screen
	HidePasswordReset *bool `json:"hidePasswordReset,omitempty" tf:"hide_password_reset,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Set of String) allowed idps to login or register
	// allowed idps to login or register
	Idps []*string `json:"idps,omitempty" tf:"idps,omitempty"`

	// (Boolean) defines if unknown username on login screen directly return an error or always display the password screen
	// defines if unknown username on login screen directly return an error or always display the password screen
	IgnoreUnknownUsernames *bool `json:"ignoreUnknownUsernames,omitempty" tf:"ignore_unknown_usernames,omitempty"`

	// (String)
	MfaInitSkipLifetime *string `json:"mfaInitSkipLifetime,omitempty" tf:"mfa_init_skip_lifetime,omitempty"`

	// (String)
	MultiFactorCheckLifetime *string `json:"multiFactorCheckLifetime,omitempty" tf:"multi_factor_check_lifetime,omitempty"`

	// (Set of String) allowed multi factors
	// allowed multi factors
	MultiFactors []*string `json:"multiFactors,omitempty" tf:"multi_factors,omitempty"`

	// (String) ID of the organization
	// ID of the organization
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// (String)
	PasswordCheckLifetime *string `json:"passwordCheckLifetime,omitempty" tf:"password_check_lifetime,omitempty"`

	// (String) defines if passwordless is allowed for users
	// defines if passwordless is allowed for users
	PasswordlessType *string `json:"passwordlessType,omitempty" tf:"passwordless_type,omitempty"`

	// (String)
	SecondFactorCheckLifetime *string `json:"secondFactorCheckLifetime,omitempty" tf:"second_factor_check_lifetime,omitempty"`

	// (Set of String) allowed second factors
	// allowed second factors
	SecondFactors []*string `json:"secondFactors,omitempty" tf:"second_factors,omitempty"`

	// (Boolean) defines if a user is allowed to login with his username and password
	// defines if a user is allowed to login with his username and password
	UserLogin *bool `json:"userLogin,omitempty" tf:"user_login,omitempty"`
}

type LoginPolicyParameters struct {

	// (Boolean) if set to true, the suffix (@domain.com) of an unknown username input on the login screen will be matched against the org domains and will redirect to the registration of that organisation on success.
	// if set to true, the suffix (@domain.com) of an unknown username input on the login screen will be matched against the org domains and will redirect to the registration of that organisation on success.
	// +kubebuilder:validation:Optional
	AllowDomainDiscovery *bool `json:"allowDomainDiscovery,omitempty" tf:"allow_domain_discovery,omitempty"`

	// (Boolean) defines if a user is allowed to add a defined identity provider. E.g. Google auth
	// defines if a user is allowed to add a defined identity provider. E.g. Google auth
	// +kubebuilder:validation:Optional
	AllowExternalIdp *bool `json:"allowExternalIdp,omitempty" tf:"allow_external_idp,omitempty"`

	// (Boolean) defines if a person is allowed to register a user on this organisation
	// defines if a person is allowed to register a user on this organisation
	// +kubebuilder:validation:Optional
	AllowRegister *bool `json:"allowRegister,omitempty" tf:"allow_register,omitempty"`

	// (String) defines where the user will be redirected to if the login is started without app context (e.g. from mail)
	// defines where the user will be redirected to if the login is started without app context (e.g. from mail)
	// +kubebuilder:validation:Optional
	DefaultRedirectURI *string `json:"defaultRedirectUri,omitempty" tf:"default_redirect_uri,omitempty"`

	// (Boolean) defines if user can additionally (to the loginname) be identified by their verified email address
	// defines if user can additionally (to the loginname) be identified by their verified email address
	// +kubebuilder:validation:Optional
	DisableLoginWithEmail *bool `json:"disableLoginWithEmail,omitempty" tf:"disable_login_with_email,omitempty"`

	// (Boolean) defines if user can additionally (to the loginname) be identified by their verified phone number
	// defines if user can additionally (to the loginname) be identified by their verified phone number
	// +kubebuilder:validation:Optional
	DisableLoginWithPhone *bool `json:"disableLoginWithPhone,omitempty" tf:"disable_login_with_phone,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	ExternalLoginCheckLifetime *string `json:"externalLoginCheckLifetime,omitempty" tf:"external_login_check_lifetime,omitempty"`

	// (Boolean) defines if a user MUST use a multi factor to log in
	// defines if a user MUST use a multi factor to log in
	// +kubebuilder:validation:Optional
	ForceMfa *bool `json:"forceMfa,omitempty" tf:"force_mfa,omitempty"`

	// (Boolean) if activated, ZITADEL only enforces MFA on local authentications. On authentications through MFA, ZITADEL won't prompt for MFA.
	// if activated, ZITADEL only enforces MFA on local authentications. On authentications through MFA, ZITADEL won't prompt for MFA.
	// +kubebuilder:validation:Optional
	ForceMfaLocalOnly *bool `json:"forceMfaLocalOnly,omitempty" tf:"force_mfa_local_only,omitempty"`

	// (Boolean) defines if password reset link should be shown in the login screen
	// defines if password reset link should be shown in the login screen
	// +kubebuilder:validation:Optional
	HidePasswordReset *bool `json:"hidePasswordReset,omitempty" tf:"hide_password_reset,omitempty"`

	// (Set of String) allowed idps to login or register
	// allowed idps to login or register
	// +crossplane:generate:reference:type=github.com/reoring/provider-zitadel/apis/org/v1alpha1.OrgIDPGithub
	// +kubebuilder:validation:Optional
	Idps []*string `json:"idps,omitempty" tf:"idps,omitempty"`

	// References to OrgIDPGithub in org to populate idps.
	// +kubebuilder:validation:Optional
	IdpsRefs []v1.Reference `json:"idpsRefs,omitempty" tf:"-"`

	// Selector for a list of OrgIDPGithub in org to populate idps.
	// +kubebuilder:validation:Optional
	IdpsSelector *v1.Selector `json:"idpsSelector,omitempty" tf:"-"`

	// (Boolean) defines if unknown username on login screen directly return an error or always display the password screen
	// defines if unknown username on login screen directly return an error or always display the password screen
	// +kubebuilder:validation:Optional
	IgnoreUnknownUsernames *bool `json:"ignoreUnknownUsernames,omitempty" tf:"ignore_unknown_usernames,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	MfaInitSkipLifetime *string `json:"mfaInitSkipLifetime,omitempty" tf:"mfa_init_skip_lifetime,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	MultiFactorCheckLifetime *string `json:"multiFactorCheckLifetime,omitempty" tf:"multi_factor_check_lifetime,omitempty"`

	// (Set of String) allowed multi factors
	// allowed multi factors
	// +kubebuilder:validation:Optional
	MultiFactors []*string `json:"multiFactors,omitempty" tf:"multi_factors,omitempty"`

	// (String) ID of the organization
	// ID of the organization
	// +crossplane:generate:reference:type=github.com/reoring/provider-zitadel/apis/org/v1alpha1.Org
	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Reference to a Org in org to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDRef *v1.Reference `json:"orgIdRef,omitempty" tf:"-"`

	// Selector for a Org in org to populate orgId.
	// +kubebuilder:validation:Optional
	OrgIDSelector *v1.Selector `json:"orgIdSelector,omitempty" tf:"-"`

	// (String)
	// +kubebuilder:validation:Optional
	PasswordCheckLifetime *string `json:"passwordCheckLifetime,omitempty" tf:"password_check_lifetime,omitempty"`

	// (String) defines if passwordless is allowed for users
	// defines if passwordless is allowed for users
	// +kubebuilder:validation:Optional
	PasswordlessType *string `json:"passwordlessType,omitempty" tf:"passwordless_type,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	SecondFactorCheckLifetime *string `json:"secondFactorCheckLifetime,omitempty" tf:"second_factor_check_lifetime,omitempty"`

	// (Set of String) allowed second factors
	// allowed second factors
	// +kubebuilder:validation:Optional
	SecondFactors []*string `json:"secondFactors,omitempty" tf:"second_factors,omitempty"`

	// (Boolean) defines if a user is allowed to login with his username and password
	// defines if a user is allowed to login with his username and password
	// +kubebuilder:validation:Optional
	UserLogin *bool `json:"userLogin,omitempty" tf:"user_login,omitempty"`
}

// LoginPolicySpec defines the desired state of LoginPolicy
type LoginPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoginPolicyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider LoginPolicyInitParameters `json:"initProvider,omitempty"`
}

// LoginPolicyStatus defines the observed state of LoginPolicy.
type LoginPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoginPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoginPolicy is the Schema for the LoginPolicys API. Resource representing the custom login policy of an organization.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zitadel}
type LoginPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.allowExternalIdp) || (has(self.initProvider) && has(self.initProvider.allowExternalIdp))",message="spec.forProvider.allowExternalIdp is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.allowRegister) || (has(self.initProvider) && has(self.initProvider.allowRegister))",message="spec.forProvider.allowRegister is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.defaultRedirectUri) || (has(self.initProvider) && has(self.initProvider.defaultRedirectUri))",message="spec.forProvider.defaultRedirectUri is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.externalLoginCheckLifetime) || (has(self.initProvider) && has(self.initProvider.externalLoginCheckLifetime))",message="spec.forProvider.externalLoginCheckLifetime is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.forceMfa) || (has(self.initProvider) && has(self.initProvider.forceMfa))",message="spec.forProvider.forceMfa is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.forceMfaLocalOnly) || (has(self.initProvider) && has(self.initProvider.forceMfaLocalOnly))",message="spec.forProvider.forceMfaLocalOnly is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hidePasswordReset) || (has(self.initProvider) && has(self.initProvider.hidePasswordReset))",message="spec.forProvider.hidePasswordReset is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ignoreUnknownUsernames) || (has(self.initProvider) && has(self.initProvider.ignoreUnknownUsernames))",message="spec.forProvider.ignoreUnknownUsernames is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.mfaInitSkipLifetime) || (has(self.initProvider) && has(self.initProvider.mfaInitSkipLifetime))",message="spec.forProvider.mfaInitSkipLifetime is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.multiFactorCheckLifetime) || (has(self.initProvider) && has(self.initProvider.multiFactorCheckLifetime))",message="spec.forProvider.multiFactorCheckLifetime is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordCheckLifetime) || (has(self.initProvider) && has(self.initProvider.passwordCheckLifetime))",message="spec.forProvider.passwordCheckLifetime is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordlessType) || (has(self.initProvider) && has(self.initProvider.passwordlessType))",message="spec.forProvider.passwordlessType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.secondFactorCheckLifetime) || (has(self.initProvider) && has(self.initProvider.secondFactorCheckLifetime))",message="spec.forProvider.secondFactorCheckLifetime is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.userLogin) || (has(self.initProvider) && has(self.initProvider.userLogin))",message="spec.forProvider.userLogin is a required parameter"
	Spec   LoginPolicySpec   `json:"spec"`
	Status LoginPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoginPolicyList contains a list of LoginPolicys
type LoginPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoginPolicy `json:"items"`
}

// Repository type metadata.
var (
	LoginPolicy_Kind             = "LoginPolicy"
	LoginPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoginPolicy_Kind}.String()
	LoginPolicy_KindAPIVersion   = LoginPolicy_Kind + "." + CRDGroupVersion.String()
	LoginPolicy_GroupVersionKind = CRDGroupVersion.WithKind(LoginPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&LoginPolicy{}, &LoginPolicyList{})
}
