package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/sdk/plugin"
	"github.com/sethvargo/vault-secrets-gen/gen"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{})

	defer func() {
		if r := recover(); r != nil {
			logger.Error("plugin paniced", "error", r)
			os.Exit(1)
		}
	}()

	meta := &api.PluginAPIClientMeta{}

	flags := meta.FlagSet()
	if err := flags.Parse(os.Args[1:]); err != nil {
		logger.Error("failed to parse flags: %s", err)
		os.Exit(1)
	}

	tlsConfig := meta.GetTLSConfig()
	tlsProviderFunc := api.VaultPluginTLSProvider(tlsConfig)

	if err := plugin.Serve(&plugin.ServeOpts{
		BackendFactoryFunc: gen.Factory,
		TLSProviderFunc:    tlsProviderFunc,
	}); err != nil {
		logger.Error("plugin shutting down", "error", err)
		os.Exit(1)
	}
}
