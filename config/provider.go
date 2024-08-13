/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	zitadel_action "github.com/reoring/provider-zitadel/config/action"
	zitadel_application_oidc "github.com/reoring/provider-zitadel/config/application_oidc"
	zitadel_login_policy "github.com/reoring/provider-zitadel/config/login_policy"
	zitadel_org "github.com/reoring/provider-zitadel/config/org"
	zitadel_org_idp_github "github.com/reoring/provider-zitadel/config/org_idp_github"
	zitadel_project "github.com/reoring/provider-zitadel/config/project"
)

const (
	resourcePrefix = "zitadel"
	modulePath     = "github.com/reoring/provider-zitadel"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("reoring.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		zitadel_action.Configure,
		zitadel_org.Configure,
		zitadel_org_idp_github.Configure,
		zitadel_application_oidc.Configure,
		zitadel_project.Configure,
		zitadel_login_policy.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
