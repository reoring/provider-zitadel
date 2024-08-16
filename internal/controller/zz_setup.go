/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	action "github.com/reoring/provider-zitadel/internal/controller/action/action"
	applicationoidc "github.com/reoring/provider-zitadel/internal/controller/application/applicationoidc"
	loginpolicy "github.com/reoring/provider-zitadel/internal/controller/login/loginpolicy"
	org "github.com/reoring/provider-zitadel/internal/controller/org/org"
	orgidpgithub "github.com/reoring/provider-zitadel/internal/controller/org/orgidpgithub"
	project "github.com/reoring/provider-zitadel/internal/controller/project/project"
	providerconfig "github.com/reoring/provider-zitadel/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		action.Setup,
		applicationoidc.Setup,
		loginpolicy.Setup,
		org.Setup,
		orgidpgithub.Setup,
		project.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
