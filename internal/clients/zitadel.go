/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/reoring/provider-zitadel/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig              = "no providerConfigRef provided"
	errGetProviderConfig             = "cannot get referenced ProviderConfig"
	errTrackUsage                    = "cannot track ProviderConfig usage"
	errExtractCredentials            = "cannot extract credentials"
	errUnmarshalCredentials          = "cannot unmarshal github credentials as JSON"
	errProviderConfigurationBuilder  = "cannot build configuration for terraform provider block"
	errTerraformProviderMissingOwner = "github provider app_auth needs owner key to be set"

	// provider config variables
	keyDomain                = "domain"
	keyInsecure              = "insecure"
	keyPort                  = "port"
	keyJwtProfileFile        = "jwt_profile_file"
)

type zitadelConfig struct {
	Domain             *string `json:"domain,omitempty"`
	Insecure           *string `json:"insecure,omitempty"`
	Port               *string `json:"port,omitempty"`
	JwtProfileFile     *string `json:"jwt_profile_file,omitempty"`
}

func terraformProviderConfigurationBuilder(creds zitadelConfig) (terraform.ProviderConfiguration, error) {

	cnf := terraform.ProviderConfiguration{}
	
	if creds.Domain != nil {
		cnf[keyDomain] = *creds.Domain
	}

	if creds.Insecure != nil {
		cnf[keyInsecure] = *creds.Insecure
	}
	
	if creds.Port != nil {
		cnf[keyPort] = *creds.Port
	}

	if creds.JwtProfileFile != nil {
		cnf[keyJwtProfileFile] = *creds.JwtProfileFile
	}

	return cnf, nil
}

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}

		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}

		creds := zitadelConfig{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		ps.Configuration, err = terraformProviderConfigurationBuilder(creds)
		if err != nil {
			return ps, errors.Wrap(err, errProviderConfigurationBuilder)
		}

		return ps, nil

	}
}
