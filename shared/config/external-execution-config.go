package config

import (
	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
)

// Configuration for external Execution clients
type ExternalExecutionConfig struct {
	hdconfig.SectionHeader

	// The selected EC
	ExecutionClient hdconfig.ChoiceParameter[ExecutionClient] //Parameter[ExecutionClient]

	// The URL of the HTTP endpoint
	HttpUrl hdconfig.StringParameter

	// The URL of the Websocket endpoint
	WebsocketUrl hdconfig.StringParameter
}

type ExternalExecutionConfigSettings struct {
	ExecutionClient ExecutionClient `json:"executionClient" yaml:"executionClient"`
	HttpUrl         string          `json:"httpUrl" yaml:"httpUrl"`
	WebsocketUrl    string          `json:"wsUrl" yaml:"wsUrl"`
}

// Generates a new ExternalExecutionConfig configuration
func NewExternalExecutionConfig() *ExternalExecutionConfig {
	cfg := &ExternalExecutionConfig{}
	cfg.ID = hdconfig.Identifier(ids.ExternalEcId)

	cfg.HttpUrl.ID = hdconfig.Identifier(ids.HttpUrlID)
	cfg.HttpUrl.Name = "HTTP URL"
	cfg.HttpUrl.Description.Default = "The URL of the HTTP RPC endpoint for your external Execution client.\nNOTE: If you are running it on the same machine as this node, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead, for example 'http://192.168.1.100:8545'."
	cfg.HttpUrl.Default = ""

	cfg.WebsocketUrl.ID = hdconfig.Identifier(ids.ExternalEcWebsocketUrlID)
	cfg.WebsocketUrl.Name = "Websocket URL"
	cfg.WebsocketUrl.Description.Default = "The URL of the Websocket RPC endpoint for your external Execution client.\nNOTE: If you are running it on the same machine as this node, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead, for example 'http://192.168.1.100:8546'."
	cfg.WebsocketUrl.Default = ""

	// Options for ExecutionClient
	options := make([]hdconfig.ParameterOption[ExecutionClient], 4)
	options[0].Name = string(ExecutionClient_Geth)
	options[0].Description.Default = "Select if your external client is Geth."
	options[0].Value = ExecutionClient_Geth

	options[1].Name = string(ExecutionClient_Nethermind)
	options[1].Description.Default = "Select if your external client is Nethermind."
	options[1].Value = ExecutionClient_Nethermind

	options[2].Name = string(ExecutionClient_Besu)
	options[2].Description.Default = "Select if your external client is Besu."
	options[2].Value = ExecutionClient_Besu

	options[3].Name = string(ExecutionClient_Reth)
	options[3].Description.Default = "Select if your external client is Reth."
	options[3].Value = ExecutionClient_Reth

	cfg.ExecutionClient.ID = hdconfig.Identifier(ids.EcID)
	cfg.ExecutionClient.Name = "Execution Client"
	cfg.ExecutionClient.Description.Default = "Select which Execution client your external client is."
	cfg.ExecutionClient.Options = options
	cfg.ExecutionClient.Default = ExecutionClient_Geth

	return cfg
}

// The title for the config
func (cfg *ExternalExecutionConfig) GetTitle() string {
	return "External Execution Client"
}

// Get the parameters for this config
func (cfg *ExternalExecutionConfig) GetParameters() []hdconfig.IParameter {
	return []hdconfig.IParameter{
		&cfg.ExecutionClient,
		&cfg.HttpUrl,
		&cfg.WebsocketUrl,
	}
}

func (cfg ExternalExecutionConfig) GetSections() []hdconfig.ISection {
	return []hdconfig.ISection{}
}
