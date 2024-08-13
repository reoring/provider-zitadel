package zitadel_org_idp_github

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_org_idp_github", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "org"
		r.Kind = "OrgIDPGithub"
		r.ExternalName = config.IdentifierFromProvider
		r.References["org_id"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}
