package group

import "github.com/crossplane/upjet/pkg/config"

// Configure configures group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("keycloak_group", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_group_memberships", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		r.References["group_id"] = config.Reference{
			TerraformName: "keycloak_group",
		}

		r.References["members"] = config.Reference{
			TerraformName: "keycloak_user",
		}
	})
	p.AddResourceConfigurator("keycloak_default_groups", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_group_roles", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		r.References["group_id"] = config.Reference{
			TerraformName: "keycloak_group",
		}

		r.References["role_dis"] = config.Reference{
			TerraformName: "keycloak_role",
		}
	})
	p.AddResourceConfigurator("keycloak_group_permissions", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		r.References["group_id"] = config.Reference{
			TerraformName: "keycloak_group",
		}
	})
}
