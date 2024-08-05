package zitadel_project

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("zitadel_project", func(r *config.Resource) {
		r.ShortGroup = "project"

		r.ExternalName = config.IdentifierFromProvider

		r.References["org"] = config.Reference{
			TerraformName: "zitadel_org",
		}
	})
}
