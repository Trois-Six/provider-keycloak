/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Realms can be imported using their name.
	"keycloak_realm": config.ParameterAsIdentifier("realm"),
	// This resource currently does not support importing.
	"keycloak_realm_events": config.IdentifierFromProvider,
	// Realm keys can be imported using realm name and keystore id, you can find it in web UI.
	"keycloak_realm_keystore_aes_generated": config.TemplatedStringAsIdentifier("keystore", "{{ .parameters.realm_id }}:{{ .external_name }}"),
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
	// This resource currently does not support importing.
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
