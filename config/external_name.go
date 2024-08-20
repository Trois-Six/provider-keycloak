/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"keycloak_group":               config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .external_name }}"), // {{realm_id}}/{{group_id}}
	"keycloak_group_memberships":   config.IdentifierFromProvider,                                                             // import not supported
	"keycloak_default_groups":      config.ParameterAsIdentifier("realm_id"),                                                  // {{realm_id}}
	"keycloak_group_roles":         config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .external_name }}"), // {{realm_id}}/{{group_id}}
	"keycloak_group_permissions":   config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .external_name }}"), // {{realm_id}}/{{group_id}}
	"keycloak_openid_client":       config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .external_name }}"), // {{realm_id}}/{{client_keycloak_id}}
	"keycloak_openid_client_scope": config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .external_name }}"), // {{realm_id}}/{{client_scope_id}}
	// TODO: need to detect if it's a client or a client scope
	// "keycloak_openid_user_attribute_protocol_mapper":    config.IdentifierFromProvider,                                                                                                  // import not supported
	// "keycloak_openid_user_property_protocol_mapper":     config.IdentifierFromProvider,                                                                                                  // import not supported
	// "keycloak_openid_group_membership_protocol_mapper":  config.IdentifierFromProvider,                                                                                                  // import not supported
	// "keycloak_openid_full_name_protocol_mapper":         config.IdentifierFromProvider,                                                                                                  // import not supported
	// "keycloak_openid_hardcoded_claim_protocol_mapper":   config.IdentifierFromProvider,                                                                                                  // import not supported
	// "keycloak_openid_audience_protocol_mapper":          config.IdentifierFromProvider,                                                                                                  // import not supported
	// "keycloak_openid_audience_resolve_protocol_mapper":  config.IdentifierFromProvider,                                                                                                  // import not supported
	// "keycloak_openid_hardcoded_role_protocol_mapper":    config.IdentifierFromProvider,                                                                                                  // import not supported
	// "keycloak_openid_user_realm_role_protocol_mapper":   config.IdentifierFromProvider,                                                                                                  // import not supported
	// "keycloak_openid_user_client_role_protocol_mapper":  config.IdentifierFromProvider,                                                                                                  // import not supported
	// "keycloak_openid_user_session_note_protocol_mapper": config.IdentifierFromProvider,                                                                                                  // import not supported
	// "keycloak_openid_script_protocol_mapper":            config.IdentifierFromProvider,                                                                                                  // import not supported
	"keycloak_openid_client_default_scopes":             config.IdentifierFromProvider,                                                                                                                                   // import not supported
	"keycloak_openid_client_optional_scopes":            config.IdentifierFromProvider,                                                                                                                                   // import not supported
	"keycloak_openid_client_authorization_resource":     config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .parameters.resource_server_id}}/{{ .external_name }}"),                                   // {{realm_id}}/{{resource_server_id}}/{{authorization_resource_id}}
	"keycloak_openid_client_authorization_scope":        config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .parameters.resource_server_id}}/{{ .external_name }}"),                                   // {{realm_id}}/{{resource_server_id}}/{{authorization_scope_id}}
	"keycloak_openid_client_authorization_permission":   config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .parameters.resource_server_id}}/{{ .external_name }}"),                                   // {{realm_id}}/{{resource_server_id}}/{{permission_id}}
	"keycloak_openid_client_group_policy":               config.IdentifierFromProvider,                                                                                                                                   // import not supported
	"keycloak_openid_client_role_policy":                config.IdentifierFromProvider,                                                                                                                                   // import not supported
	"keycloak_openid_client_aggregate_policy":           config.IdentifierFromProvider,                                                                                                                                   // import not supported
	"keycloak_openid_client_js_policy":                  config.IdentifierFromProvider,                                                                                                                                   // import not supported
	"keycloak_openid_client_time_policy":                config.IdentifierFromProvider,                                                                                                                                   // import not supported
	"keycloak_openid_client_user_policy":                config.IdentifierFromProvider,                                                                                                                                   // import not supported
	"keycloak_openid_client_client_policy":              config.IdentifierFromProvider,                                                                                                                                   // import not supported
	"keycloak_openid_client_service_account_role":       config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .parameters.service_account_user_id }}/{{ .parameters.client_id }}/{{ .external_name }}"), // {{realm_id}}/{{service_account_user_id}}/{{client_id}}/{{role_id}}
	"keycloak_openid_client_service_account_realm_role": config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .parameters.service_account_user_id }}/{{ .external_name }}"),                             // {{realm_id}}/{{service_account_user_id}}/{{role_id}}
	"keycloak_openid_client_permissions":                config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .parameters.client_id }}"),                                                                // {{realm_id}}/{{client_id}}
	"keycloak_realm":                                    config.NameAsIdentifier,                                                                                                                                         // {{realm}}
	"keycloak_realm_events":                             config.IdentifierFromProvider,                                                                                                                                   // import not supported
	"keycloak_realm_keystore_aes_generated":             config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .external_name }}"),                                                                       // {{realm_id}}/{{keystore_id}}
	"keycloak_realm_keystore_ecdsa_generated":           config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .external_name }}"),                                                                       // {{realm_id}}/{{keystore_id}}
	"keycloak_realm_keystore_hmac_generated":            config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .external_name }}"),                                                                       // {{realm_id}}/{{keystore_id}}
	"keycloak_realm_keystore_java_keystore":             config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .external_name }}"),                                                                       // {{realm_id}}/{{keystore_id}}
	"keycloak_realm_keystore_rsa":                       config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .external_name }}"),                                                                       // {{realm_id}}/{{keystore_id}}
	"keycloak_realm_keystore_rsa_generated":             config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .external_name }}"),                                                                       // {{realm_id}}/{{keystore_id}}
	"keycloak_realm_user_profile":                       config.IdentifierFromProvider,                                                                                                                                   // import not supported
	"keycloak_role":                                     config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .parameters.role_id }}"),                                                                  // {{realm_id}}/{role_id}}
	"keycloak_user":                                     config.TemplatedStringAsIdentifier("", "{{ .parameters.realm_id }}/{{ .parameters.user_id }}"),                                                                  // {{realm_id}}/{{user_id}}
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
