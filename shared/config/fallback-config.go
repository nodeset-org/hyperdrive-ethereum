package config

import (
	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
)

// Fallback configuration
type FallbackConfig struct {
	// Flag for enabling fallback clients
	UseFallbackClients hdconfig.BoolParameter

	// The URL of the Execution Client HTTP endpoint
	EcHttpUrl hdconfig.StringParameter

	// The URL of the Beacon Node HTTP endpoint
	BnHttpUrl hdconfig.StringParameter

	// The URL of the Prysm gRPC endpoint (only needed if using Prysm VCs)
	PrysmRpcUrl hdconfig.StringParameter
}

// Generates a new FallbackConfig configuration
func NewFallbackConfig() *FallbackConfig {
	cfg := &FallbackConfig{}

	// TODO: Get these reviewed
	// Use Fallback Clients
	cfg.UseFallbackClients.ID = hdconfig.Identifier(ids.FallbackUseFallbackClientsID)
	cfg.UseFallbackClients.Name = "Use Fallback Clients"
	cfg.UseFallbackClients.Description.Default = "Enable this if you would like to specify a fallback Execution and Beacon Node, which will temporarily be used by your node and Validator Client(s) if your primary Execution / Beacon Node pair ever go offline (e.g. if you switch, prune, or resync your clients)."
	cfg.UseFallbackClients.AffectedContainers = []string{string(ContainerID_Daemon)}

	cfg.EcHttpUrl.ID = hdconfig.Identifier(ids.FallbackEcHttpUrlID)
	cfg.EcHttpUrl.Name = "Execution Client URL"
	cfg.EcHttpUrl.Description.Default = "The URL of the HTTP API endpoint for your fallback Execution client.\n\nNOTE: If you are running it on the same machine as your node, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead."
	cfg.EcHttpUrl.AffectedContainers = []string{string(ContainerID_Daemon)}

	cfg.BnHttpUrl.ID = hdconfig.Identifier(ids.FallbackBnHttpUrlID)
	cfg.BnHttpUrl.Name = "Beacon Node URL"
	cfg.BnHttpUrl.Description.Default = "The URL of the HTTP Beacon API endpoint for your fallback Beacon Node.\n\nNOTE: If you are running it on the same machine as your node, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead."
	cfg.BnHttpUrl.AffectedContainers = []string{string(ContainerID_Daemon)}

	cfg.PrysmRpcUrl.ID = hdconfig.Identifier(ids.FallbackPrysmRpcUrlID)
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
func (cfg *FallbackConfig) GetParameters() []hdconfig.IParameter {
	return []hdconfig.IParameter{
		&cfg.UseFallbackClients,
		&cfg.EcHttpUrl,
		&cfg.BnHttpUrl,
		&cfg.PrysmRpcUrl,
	}
}

// Get the sections underneath this one
func (cfg *FallbackConfig) GetSections() map[string]hdconfig.ISection {
	return map[string]hdconfig.ISection{}
}
