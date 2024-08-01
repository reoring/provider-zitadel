// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	action "github.com/reoring/provider-zitadel/internal/controller/action/action"
	oidc "github.com/reoring/provider-zitadel/internal/controller/application_oidc/oidc"
	org "github.com/reoring/provider-zitadel/internal/controller/org/org"
	idpgithub "github.com/reoring/provider-zitadel/internal/controller/org_idp_github/idpgithub"
	project "github.com/reoring/provider-zitadel/internal/controller/project/project"
	providerconfig "github.com/reoring/provider-zitadel/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		action.Setup,
		oidc.Setup,
		org.Setup,
		idpgithub.Setup,
		project.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
