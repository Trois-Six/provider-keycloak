// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	realm "github.com/trois-six/provider-keycloak/internal/controller/keycloak/realm"
	providerconfig "github.com/trois-six/provider-keycloak/internal/controller/providerconfig"
	events "github.com/trois-six/provider-keycloak/internal/controller/realm/events"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		realm.Setup,
		providerconfig.Setup,
		events.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
