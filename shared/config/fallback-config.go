package config

import (
	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	"github.com/nodeset-org/hyperdrive/modules/config"
)

// Fallback configuration
type FallbackConfig struct {
	// Flag for enabling fallback clients
	UseFallbackClients config.BoolParameter

	// The URL of the Execution Client HTTP endpoint
	EcHttpUrl config.StringParameter

	// The URL of the Beacon Node HTTP endpoint
	BnHttpUrl config.StringParameter

	// The URL of the Prysm gRPC endpoint (only needed if using Prysm VCs)
	PrysmRpcUrl config.StringParameter
}

// Generates a new FallbackConfig configuration
func NewFallbackConfig() *FallbackConfig {
	cfg := &FallbackConfig{}

	// TODO: Get these reviewed
	// Use Fallback Clients
	cfg.UseFallbackClients.ID = config.Identifier(ids.FallbackUseFallbackClientsID)
	cfg.UseFallbackClients.Name = "Use Fallback Clients"
	cfg.UseFallbackClients.Description.Default = "Enable this if you would like to specify a fallback Execution and Beacon Node, which will temporarily be used by your node and Validator Client(s) if your primary Execution / Beacon Node pair ever go offline (e.g. if you switch, prune, or resync your clients)."
	cfg.UseFallbackClients.AffectedContainers = []string{string(ContainerID_Daemon)}

	cfg.EcHttpUrl.ID = config.Identifier(ids.FallbackEcHttpUrlID)
	cfg.EcHttpUrl.Name = "Execution Client URL"
	cfg.EcHttpUrl.Description.Default = "The URL of the HTTP API endpoint for your fallback Execution client.\n\nNOTE: If you are running it on the same machine as your node, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead."
	cfg.EcHttpUrl.AffectedContainers = []string{string(ContainerID_Daemon)}

	cfg.BnHttpUrl.ID = config.Identifier(ids.FallbackBnHttpUrlID)
	cfg.BnHttpUrl.Name = "Beacon Node URL"
	cfg.BnHttpUrl.Description.Default = "The URL of the HTTP Beacon API endpoint for your fallback Beacon Node.\n\nNOTE: If you are running it on the same machine as your node, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead."
	cfg.BnHttpUrl.AffectedContainers = []string{string(ContainerID_Daemon)}

	cfg.PrysmRpcUrl.ID = config.Identifier(ids.FallbackPrysmRpcUrlID)
	cfg.PrysmRpcUrl.Name = "Prysm RPC URL"
	cfg.PrysmRpcUrl.Description.Default = "**Only used if you have a Prysm Validator Client.**\n\nThe URL of Prysm's gRPC API endpoint for your fallback Beacon Node. Prysm's Validator Client will need this in order to connect to it.\nNOTE: If you are running it on the same machine as your node, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead."
	cfg.PrysmRpcUrl.AffectedContainers = []string{string(ContainerID_Daemon)}

	return cfg
}

// The title for the config
func (cfg *FallbackConfig) GetTitle() string {
	return "Fallback Clients"
}

// Get the Parameters for this config
func (cfg *FallbackConfig) GetParameters() []config.IParameter {
	return []config.IParameter{
		&cfg.UseFallbackClients,
		&cfg.EcHttpUrl,
		&cfg.BnHttpUrl,
		&cfg.PrysmRpcUrl,
	}
}

// Get the sections underneath this one
func (cfg *FallbackConfig) GetSubconfigs() map[string]IConfigSection {
	return map[string]IConfigSection{}
}
