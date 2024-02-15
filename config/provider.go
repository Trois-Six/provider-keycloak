/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/trois-six/provider-keycloak/config/realm"
	"github.com/trois-six/provider-keycloak/config/realmKeystoreAESGenerated"
)

const (
	resourcePrefix = "keycloak"
	modulePath     = "github.com/trois-six/provider-keycloak"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("keycloak.crd.alt"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		realm.Configure,
		realmKeystoreAESGenerated.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
