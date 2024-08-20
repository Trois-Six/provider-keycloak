package openid

import "github.com/crossplane/upjet/pkg/config"

const (
	shortGroup = "openid"
)

// Configure configures group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("keycloak_openid_client", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_openid_client_scope", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_openid_user_attribute_protocol_mapper", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_openid_user_property_protocol_mapper", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_openid_group_membership_protocol_mapper", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_openid_full_name_protocol_mapper", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_openid_hardcoded_claim_protocol_mapper", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_openid_audience_protocol_mapper", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_openid_audience_resolve_protocol_mapper", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_openid_hardcoded_role_protocol_mapper", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_openid_user_realm_role_protocol_mapper", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_openid_user_client_role_protocol_mapper", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_openid_user_session_note_protocol_mapper", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_openid_script_protocol_mapper", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}
	})
	p.AddResourceConfigurator("keycloak_openid_client_default_scopes", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		r.References["client_id"] = config.Reference{
			TerraformName: "keycloak_openid_client",
		}
	})
	p.AddResourceConfigurator("keycloak_openid_client_optional_scopes", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		r.References["client_id"] = config.Reference{
			TerraformName: "keycloak_openid_client",
		}
	})
	p.AddResourceConfigurator("keycloak_openid_client_authorization_resource", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		//TODO
		// r.References["resource_server_id"] = config.Reference{
		// 	TerraformName: "XXXX",
		// }
	})
	p.AddResourceConfigurator("keycloak_openid_client_authorization_scope", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		//TODO
		// r.References["resource_server_id"] = config.Reference{
		// 	TerraformName: "XXXX",
		// }
	})
	p.AddResourceConfigurator("keycloak_openid_client_authorization_permission", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		//TODO
		// r.References["resource_server_id"] = config.Reference{
		// 	TerraformName: "XXXX",
		// }
	})
	p.AddResourceConfigurator("keycloak_openid_client_group_policy", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		//TODO
		// r.References["resource_server_id"] = config.Reference{
		// 	TerraformName: "XXXX",
		// }
	})
	p.AddResourceConfigurator("keycloak_openid_client_role_policy", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		//TODO
		// r.References["resource_server_id"] = config.Reference{
		// 	TerraformName: "XXXX",
		// }
	})
	p.AddResourceConfigurator("keycloak_openid_client_aggregate_policy", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		//TODO
		// r.References["resource_server_id"] = config.Reference{
		// 	TerraformName: "XXXX",
		// }
	})
	p.AddResourceConfigurator("keycloak_openid_client_js_policy", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		//TODO
		// r.References["resource_server_id"] = config.Reference{
		// 	TerraformName: "XXXX",
		// }
	})
	p.AddResourceConfigurator("keycloak_openid_client_time_policy", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		//TODO
		// r.References["resource_server_id"] = config.Reference{
		// 	TerraformName: "XXXX",
		// }
	})
	p.AddResourceConfigurator("keycloak_openid_client_user_policy", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		//TODO
		// r.References["resource_server_id"] = config.Reference{
		// 	TerraformName: "XXXX",
		// }
	})
	p.AddResourceConfigurator("keycloak_openid_client_client_policy", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		//TODO
		// r.References["resource_server_id"] = config.Reference{
		// 	TerraformName: "XXXX",
		// }
	})
	p.AddResourceConfigurator("keycloak_openid_client_service_account_role", func(r *config.Resource) {
		r.References["service_account_user_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		//TODO: check that it's correct
		r.References["service_account_user_id"] = config.Reference{
			TerraformName: "keycloak_openid_client",
		}

		r.References["client_id"] = config.Reference{
			TerraformName: "keycloak_openid_client",
		}

		r.References["role"] = config.Reference{
			TerraformName: "keycloak_role",
		}

	})
	p.AddResourceConfigurator("keycloak_openid_client_service_account_realm_role", func(r *config.Resource) {
		r.References["service_account_user_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		//TODO: check that it's correct
		r.References["service_account_user_id"] = config.Reference{
			TerraformName: "keycloak_openid_client",
		}

		r.References["role"] = config.Reference{
			TerraformName: "keycloak_role",
		}
	})
	p.AddResourceConfigurator("keycloak_openid_client_permissions", func(r *config.Resource) {
		r.References["realm_id"] = config.Reference{
			TerraformName: "keycloak_realm",
		}

		r.References["client_id"] = config.Reference{
			TerraformName: "keycloak_openid_client",
		}
	})
}
