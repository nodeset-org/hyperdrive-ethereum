package config

import (
	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	"github.com/nodeset-org/hyperdrive/modules/config"
)

// Configuration for external Execution clients
type ExternalExecutionConfig struct {
	// The selected EC
	ExecutionClient Parameter[ExecutionClient]

	// The URL of the HTTP endpoint
	HttpUrl config.StringParameter

	// The URL of the Websocket endpoint
	WebsocketUrl config.StringParameter
}

// Generates a new ExternalExecutionConfig configuration
func NewExternalExecutionConfig() *ExternalExecutionConfig {
	cfg := &ExternalExecutionConfig{}

	cfg.HttpUrl.ID = config.Identifier(ids.HttpUrlID)
	cfg.HttpUrl.Name = "HTTP URL"
	cfg.HttpUrl.Description.Default = "The URL of the HTTP RPC endpoint for your external Execution client.\nNOTE: If you are running it on the same machine as this node, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead, for example 'http://192.168.1.100:8545'."
	cfg.HttpUrl.AffectedContainers = []string{string(ContainerID_Daemon)}

	cfg.WebsocketUrl.ID = config.Identifier(ids.ExternalEcWebsocketUrlID)
	cfg.WebsocketUrl.Name = "Websocket URL"
	cfg.WebsocketUrl.Description.Default = "The URL of the Websocket RPC endpoint for your external Execution client.\nNOTE: If you are running it on the same machine as this node, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead, for example 'http://192.168.1.100:8546'."
	cfg.WebsocketUrl.AffectedContainers = []string{}

	return cfg

	// return &ExternalExecutionConfig{
	// 	ExecutionClient: Parameter[ExecutionClient]{
	// 		ParameterCommon: &ParameterCommon{
	// 			ID:                 ids.EcID,
	// 			Name:               "Execution Client",
	// 			Description:        "Select which Execution client your external client is.",
	// 			AffectsContainers:  []ContainerID{ContainerID_ValidatorClient},
	// 			CanBeBlank:         false,
	// 			OverwriteOnUpgrade: false,
	// 		},
	// 		Options: []*ParameterOption[ExecutionClient]{
	// 			{
	// 				ParameterOptionCommon: &ParameterOptionCommon{
	// 					Name:        "Geth",
	// 					Description: "Select if your external client is Geth.",
	// 				},
	// 				Value: ExecutionClient_Geth,
	// 			}, {
	// 				ParameterOptionCommon: &ParameterOptionCommon{
	// 					Name:        "Nethermind",
	// 					Description: "Select if your external client is Nethermind.",
	// 				},
	// 				Value: ExecutionClient_Nethermind,
	// 			}, {
	// 				ParameterOptionCommon: &ParameterOptionCommon{
	// 					Name:        "Besu",
	// 					Description: "Select if your external client is Besu.",
	// 				},
	// 				Value: ExecutionClient_Besu,
	// 			}, {
	// 				ParameterOptionCommon: &ParameterOptionCommon{
	// 					Name:        "Reth",
	// 					Description: "Select if your external client is Reth.",
	// 				},
	// 				Value: ExecutionClient_Reth,
	// 			}},
	// 		Default: map[Network]ExecutionClient{
	// 			Network_All: ExecutionClient_Geth},
	// 	},
}

// The title for the config
func (cfg *ExternalExecutionConfig) GetTitle() string {
	return "External Execution Client"
}

// Get the parameters for this config
func (cfg *ExternalExecutionConfig) GetParameters() []config.IParameter {
	return []config.IParameter{
		&cfg.ExecutionClient,
		&cfg.HttpUrl,
		&cfg.WebsocketUrl,
	}
}

// Get the sections underneath this one
func (cfg *ExternalExecutionConfig) GetSubconfigs() map[string]IConfigSection {
	return map[string]IConfigSection{}
}
