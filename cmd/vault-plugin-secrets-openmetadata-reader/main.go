package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	omsecrets "github.com/cdoron/vault-plugin-secrets-openmetadata-reader"
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/sdk/plugin"
)

func main() {
	// Boilerplate code to get started.
	// Please see https://www.hashicorp.com/blog/building-a-vault-secure-plugin for more info.
	apiClientMeta := &api.PluginAPIClientMeta{}
	flags := apiClientMeta.FlagSet()
	err := flags.Parse(os.Args[1:])
	if err != nil {
		logger := hclog.New(&hclog.LoggerOptions{})

		logger.Error("Error in flags.Parse", err)
		os.Exit(1)
	}

	tlsConfig := apiClientMeta.GetTLSConfig()
	tlsProviderFunc := api.VaultPluginTLSProvider(tlsConfig)

	err = plugin.Serve(&plugin.ServeOpts{
		BackendFactoryFunc: omsecrets.Factory,
		TLSProviderFunc:    tlsProviderFunc,
	})
	if err != nil {
		logger := hclog.New(&hclog.LoggerOptions{})

		logger.Error("plugin shutting down", "error", err)
		os.Exit(1)
	}
}
