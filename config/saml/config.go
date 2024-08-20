package saml

import "github.com/crossplane/upjet/pkg/config"

const (
	shortGroup = "saml"
)

// Configure configures group
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("keycloak_saml_client", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("keycloak_saml_client_scope", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("keycloak_saml_client_default_scopes", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("keycloak_saml_user_attribute_protocol_mapper", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("keycloak_saml_user_property_protocol_mapper", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("keycloak_saml_script_protocol_mapper", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("keycloak_saml_identity_provider", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
}
