// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Mozilla Public License 2.0

package omsecrets

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/mitchellh/mapstructure"

	client "fybrik.io/openmetadata-connector/datacatalog-go-client"
)

type OpenMetadataSecretsReader struct {
	client *client.APIClient
}

const FybrikAccessKeyString = "access_key"
const FybrikSecretKeyString = "secret_key"

func extractSecretsFromConfig(config map[string]interface{}) (map[string]interface{}, error) {
	type SecurityConfig struct {
		AccessKey string `mapstructure:"awsAccessKeyId"`
		SecretKey string `mapstructure:"awsSecretAccessKey"`
	}
	type ConfigSource struct {
		SecurityConfig map[string]interface{} `mapstructure:",securityConfig"`
	}

	type Config struct {
		ConfigSource map[string]interface{} `mapstructure:",configSource"`
	}

	var c Config
	err := mapstructure.Decode(config, &c)
	if err != nil {
		return nil, err
	}

	var configSource ConfigSource
	err = mapstructure.Decode(c.ConfigSource, &configSource)
	if err != nil {
		return nil, err
	}

	var securityConfig SecurityConfig
	err = mapstructure.Decode(configSource.SecurityConfig, &securityConfig)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		FybrikAccessKeyString: securityConfig.AccessKey,
		FybrikSecretKeyString: securityConfig.SecretKey,
	}, nil
}

// GetSecret returns the content of openmetadata secret.
func (s *OpenMetadataSecretsReader) GetSecret(ctx context.Context, secretName string, log hclog.Logger) (map[string]interface{}, error) {
	databaseService, _, err := s.client.DatabaseServiceApi.GetDatabaseServiceByFQN(ctx, secretName).Execute()
	if err != nil {
		return nil, err
	}
	config := databaseService.Connection.GetConfig()
	return extractSecretsFromConfig(config)
}
