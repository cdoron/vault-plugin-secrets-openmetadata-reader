// Copyright 2022 IBM Corp.
// SPDX-License-Identifier: Mozilla Public License 2.0

package omclient

import (
	"context"

	client "fybrik.io/openmetadata-connector/datacatalog-go-client"
	"github.com/mitchellh/mapstructure"
)

type OMClient struct {
}

// Return WKClient object
func NewOMClient() *client.APIClient {
	conf := client.Configuration{Servers: client.ServerConfigurations{
		client.ServerConfiguration{
			URL:         "http://localhost:8585/api",
			Description: "Endpoint URL",
		},
	},
	}

	return client.NewAPIClient(&conf)
}

const FybrikAccessKeyString = "access_key"
const FybrikSecretKeyString = "secret_key"

func (o *OMClient) ExtractSecretsFromConfig(config map[string]interface{}) (map[string]interface{}, error) {
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

func (o *OMClient) GetConnectionInformation(ctx context.Context, connectionName string) (*client.DatabaseService, error) {
	c := NewOMClient()
	databaseService, _, err := c.DatabaseServiceApi.GetDatabaseServiceByFQN(ctx, connectionName).Execute()
	if err != nil {
		return nil, err
	}
	return databaseService, nil
}
