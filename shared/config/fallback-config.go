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

type FallbackConfigSettings struct {
	UseFallbackClients bool   `json:"useFallbackClients" yaml:"useFallbackClients"`
	EcHttpUrl          string `json:"ecHttpUrl" yaml:"ecHttpUrl"`
	BnHttpUrl          string `json:"bnHttpUrl" yaml:"bnHttpUrl"`
	PrysmRpcUrl        string `json:"prysmRpcUrl" yaml:"prysmRpcUrl"`
}

// Generates a new FallbackConfig configuration
func NewFallbackConfig() *FallbackConfig {
	cfg := &FallbackConfig{}

	// Use Fallback Clients
	cfg.UseFallbackClients.ID = hdconfig.Identifier(ids.FallbackUseFallbackClientsID)
	cfg.UseFallbackClients.Name = "Use Fallback Clients"
	cfg.UseFallbackClients.Description.Default = "Enable this if you would like to specify a fallback Execution and Beacon Node, which will temporarily be used by your node and Validator Client(s) if your primary Execution / Beacon Node pair ever go offline (e.g. if you switch, prune, or resync your clients)."
	cfg.UseFallbackClients.Default = false

	cfg.EcHttpUrl.ID = hdconfig.Identifier(ids.FallbackEcHttpUrlID)
	cfg.EcHttpUrl.Name = "Execution Client URL"
	cfg.EcHttpUrl.Description.Default = "The URL of the HTTP API endpoint for your fallback Execution client.\n\nNOTE: If you are running it on the same machine as your node, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead."
	cfg.EcHttpUrl.Default = ""

	cfg.BnHttpUrl.ID = hdconfig.Identifier(ids.FallbackBnHttpUrlID)
	cfg.BnHttpUrl.Name = "Beacon Node URL"
	cfg.BnHttpUrl.Description.Default = "The URL of the HTTP Beacon API endpoint for your fallback Beacon Node.\n\nNOTE: If you are running it on the same machine as your node, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead."
	cfg.BnHttpUrl.Default = ""

	cfg.PrysmRpcUrl.ID = hdconfig.Identifier(ids.FallbackPrysmRpcUrlID)
	cfg.PrysmRpcUrl.Name = "Prysm RPC URL"
	cfg.PrysmRpcUrl.Description.Default = "**Only used if you have a Prysm Validator Client.**\n\nThe URL of Prysm's gRPC API endpoint for your fallback Beacon Node. Prysm's Validator Client will need this in order to connect to it.\nNOTE: If you are running it on the same machine as your node, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead."
	cfg.PrysmRpcUrl.Default = ""

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

func (cfg *FallbackConfig) GetDescription() hdconfig.DynamicProperty[string] {
	return hdconfig.DynamicProperty[string]{}
}

func (cfg *FallbackConfig) GetDisabled() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}

func (cfg *FallbackConfig) GetHidden() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}

func (cfg *FallbackConfig) GetID() hdconfig.Identifier {
	return hdconfig.Identifier(ids.FallbackID)
}

func (cfg *FallbackConfig) GetName() string {
	return "Fallback"
}

func (cfg *FallbackConfig) GetSections() []hdconfig.ISection {
	return []hdconfig.ISection{}
}
