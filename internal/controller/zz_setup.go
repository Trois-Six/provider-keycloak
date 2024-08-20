// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	groups "github.com/trois-six/provider-keycloak/internal/controller/default/groups"
	memberships "github.com/trois-six/provider-keycloak/internal/controller/group/memberships"
	permissions "github.com/trois-six/provider-keycloak/internal/controller/group/permissions"
	roles "github.com/trois-six/provider-keycloak/internal/controller/group/roles"
	group "github.com/trois-six/provider-keycloak/internal/controller/keycloak/group"
	realm "github.com/trois-six/provider-keycloak/internal/controller/keycloak/realm"
	role "github.com/trois-six/provider-keycloak/internal/controller/keycloak/role"
	user "github.com/trois-six/provider-keycloak/internal/controller/keycloak/user"
	client "github.com/trois-six/provider-keycloak/internal/controller/openid/client"
	clientaggregatepolicy "github.com/trois-six/provider-keycloak/internal/controller/openid/clientaggregatepolicy"
	clientauthorizationpermission "github.com/trois-six/provider-keycloak/internal/controller/openid/clientauthorizationpermission"
	clientauthorizationresource "github.com/trois-six/provider-keycloak/internal/controller/openid/clientauthorizationresource"
	clientauthorizationscope "github.com/trois-six/provider-keycloak/internal/controller/openid/clientauthorizationscope"
	clientclientpolicy "github.com/trois-six/provider-keycloak/internal/controller/openid/clientclientpolicy"
	clientdefaultscopes "github.com/trois-six/provider-keycloak/internal/controller/openid/clientdefaultscopes"
	clientgrouppolicy "github.com/trois-six/provider-keycloak/internal/controller/openid/clientgrouppolicy"
	clientjspolicy "github.com/trois-six/provider-keycloak/internal/controller/openid/clientjspolicy"
	clientoptionalscopes "github.com/trois-six/provider-keycloak/internal/controller/openid/clientoptionalscopes"
	clientpermissions "github.com/trois-six/provider-keycloak/internal/controller/openid/clientpermissions"
	clientrolepolicy "github.com/trois-six/provider-keycloak/internal/controller/openid/clientrolepolicy"
	clientscope "github.com/trois-six/provider-keycloak/internal/controller/openid/clientscope"
	clientserviceaccountrealmrole "github.com/trois-six/provider-keycloak/internal/controller/openid/clientserviceaccountrealmrole"
	clientserviceaccountrole "github.com/trois-six/provider-keycloak/internal/controller/openid/clientserviceaccountrole"
	clienttimepolicy "github.com/trois-six/provider-keycloak/internal/controller/openid/clienttimepolicy"
	clientuserpolicy "github.com/trois-six/provider-keycloak/internal/controller/openid/clientuserpolicy"
	providerconfig "github.com/trois-six/provider-keycloak/internal/controller/providerconfig"
	events "github.com/trois-six/provider-keycloak/internal/controller/realm/events"
	keystoreaesgenerated "github.com/trois-six/provider-keycloak/internal/controller/realm/keystoreaesgenerated"
	keystoreecdsagenerated "github.com/trois-six/provider-keycloak/internal/controller/realm/keystoreecdsagenerated"
	keystorehmacgenerated "github.com/trois-six/provider-keycloak/internal/controller/realm/keystorehmacgenerated"
	keystorejavakeystore "github.com/trois-six/provider-keycloak/internal/controller/realm/keystorejavakeystore"
	keystorersa "github.com/trois-six/provider-keycloak/internal/controller/realm/keystorersa"
	keystorersagenerated "github.com/trois-six/provider-keycloak/internal/controller/realm/keystorersagenerated"
	userprofile "github.com/trois-six/provider-keycloak/internal/controller/realm/userprofile"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		groups.Setup,
		memberships.Setup,
		permissions.Setup,
		roles.Setup,
		group.Setup,
		realm.Setup,
		role.Setup,
		user.Setup,
		client.Setup,
		clientaggregatepolicy.Setup,
		clientauthorizationpermission.Setup,
		clientauthorizationresource.Setup,
		clientauthorizationscope.Setup,
		clientclientpolicy.Setup,
		clientdefaultscopes.Setup,
		clientgrouppolicy.Setup,
		clientjspolicy.Setup,
		clientoptionalscopes.Setup,
		clientpermissions.Setup,
		clientrolepolicy.Setup,
		clientscope.Setup,
		clientserviceaccountrealmrole.Setup,
		clientserviceaccountrole.Setup,
		clienttimepolicy.Setup,
		clientuserpolicy.Setup,
		providerconfig.Setup,
		events.Setup,
		keystoreaesgenerated.Setup,
		keystoreecdsagenerated.Setup,
		keystorehmacgenerated.Setup,
		keystorejavakeystore.Setup,
		keystorersa.Setup,
		keystorersagenerated.Setup,
		userprofile.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
