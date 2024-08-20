package role

import "github.com/crossplane/upjet/pkg/config"

// Configure configures group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("keycloak_role", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		r.References["client_id"] = config.Reference{
			TerraformName: "keycloak_openid_client",
		}
	})
}
