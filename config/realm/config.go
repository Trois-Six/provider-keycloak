package realm

import "github.com/crossplane/upjet/pkg/config"

// Configure configures realm
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("keycloak_realm_events", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_realm_keystore_aes_generated", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_realm_keystore_ecdsa_generated", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_realm_keystore_hmac_generated", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_realm_keystore_java_keystore", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_realm_keystore_rsa", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_realm_keystore_rsa_generated", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_realm_user_profile", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
}
