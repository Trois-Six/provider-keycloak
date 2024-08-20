package user

import "github.com/crossplane/upjet/pkg/config"

// Configure configures group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("keycloak_user", func(r *config.Resource) {
		r.References["service_account_user_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_user_roles", func(r *config.Resource) {
	})
	p.AddResourceConfigurator("keycloak_default_roles", func(r *config.Resource) {
	})
	p.AddResourceConfigurator("keycloak_users_permissions", func(r *config.Resource) {
	})
	p.AddResourceConfigurator("keycloak_user_groups", func(r *config.Resource) {
	})
}
