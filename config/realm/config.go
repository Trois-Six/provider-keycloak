package realm

import "github.com/crossplane/upjet/pkg/config"

// Configure configures realm
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("keycloak_realm", func(r *config.Resource) {
		r.Kind = "Realm"
	})
}
