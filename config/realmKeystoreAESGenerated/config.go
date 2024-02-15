package realmKeystoreAESGenerated

import "github.com/crossplane/upjet/pkg/config"

// Configure configures realm keystore
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("keycloak_realm_keystore_aes_generated", func(r *config.Resource) {
		r.Kind = "RealmKeyStoreAESGenerated"

		r.References["realm"] = config.Reference{
			Type: "github.com/trois-six/provider-keycloak/apis/realm/v1alpha1.Realm",
		}
	})
}
