/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/trois-six/provider-keycloak/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal keycloak credentials as JSON"
	// Terraform Provider configuration block keys
	keyUrl                   = "url"
	keyBasePath              = "base_path"
	keyClientId              = "client_id"
	keyClientSecret          = "client_secret"
	keyUsername              = "username"
	keyPassword              = "password"
	keyRealm                 = "realm"
	keyInitialLogin          = "initial_login"
	keyClientTimeout         = "client_timeout"
	keyTLSInsecureSkipVerify = "tls_insecure_skip_verify"
	keyRootCACertificate     = "root_ca_certificate"
	keyRedHatSSO             = "red_hat_sso"
	keyAdditionalHeaders     = "additional_headers"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]any{}
		if v, ok := creds[keyUrl]; ok {
			ps.Configuration[keyUrl] = v
		}
		if v, ok := creds[keyBasePath]; ok {
			ps.Configuration[keyBasePath] = v
		}
		if v, ok := creds[keyClientId]; ok {
			ps.Configuration[keyClientId] = v
		}
		if v, ok := creds[keyClientSecret]; ok {
			ps.Configuration[keyClientSecret] = v
		}
		if v, ok := creds[keyUsername]; ok {
			ps.Configuration[keyUsername] = v
		}
		if v, ok := creds[keyPassword]; ok {
			ps.Configuration[keyPassword] = v
		}
		if v, ok := creds[keyRealm]; ok {
			ps.Configuration[keyRealm] = v
		}
		if v, ok := creds[keyInitialLogin]; ok {
			ps.Configuration[keyInitialLogin] = v
		}
		if v, ok := creds[keyClientTimeout]; ok {
			ps.Configuration[keyClientTimeout] = v
		}
		if v, ok := creds[keyInitialLogin]; ok {
			ps.Configuration[keyInitialLogin] = v
		}
		if v, ok := creds[keyTLSInsecureSkipVerify]; ok {
			ps.Configuration[keyTLSInsecureSkipVerify] = v
		}
		if v, ok := creds[keyRootCACertificate]; ok {
			ps.Configuration[keyRootCACertificate] = v
		}
		if v, ok := creds[keyRedHatSSO]; ok {
			ps.Configuration[keyRedHatSSO] = v
		}
		if v, ok := creds[keyAdditionalHeaders]; ok {
			ps.Configuration[keyAdditionalHeaders] = v
		}

		return ps, nil
	}
}
