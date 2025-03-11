package config

import (
	"fmt"

	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
)

const (
	defaultLocalExecutionHttpPort      uint64 = 8545
	defaultLocalExecutionWebsocketPort uint64 = 8546
	defaultLocalExecutionEnginePort    uint64 = 8551
	defaultLocalExecutionP2pPort       uint64 = 30303
)

// Configuration for the Execution client
type LocalExecutionConfig struct {
	hdconfig.SectionHeader
	// The selected EC
	ExecutionClient hdconfig.ChoiceParameter[ExecutionClient] //Parameter[ExecutionClient]

	// The Hostname
	HostName hdconfig.StringParameter

	// The HTTP API port
	HttpPort hdconfig.UintParameter

	// The Websocket API port
	WebsocketPort hdconfig.UintParameter

	// The Engine API port
	EnginePort hdconfig.UintParameter

	// Toggle for forwarding the HTTP API port outside of Docker
	OpenApiPorts hdconfig.ChoiceParameter[RpcPortMode] //Parameter[RpcPortMode]

	// P2P traffic port
	P2pPort hdconfig.UintParameter

	//Comma separated
	AdditionalDockerNetworks hdconfig.StringParameter

	// Subconfigs
	Geth       *GethConfig
	Nethermind *NethermindConfig
	Besu       *BesuConfig
	Reth       *RethConfig
}

type LocalExecutionConfigSettings struct {
	ExecutionClient          ExecutionClient `json:"executionClient" yaml:"executionClient"`
	HostName                 string          `json:"hostName" yaml:"hostName"`
	HttpPort                 uint64          `json:"httpPort" yaml:"httpPort"`
	WebsocketPort            uint64          `json:"wsPort" yaml:"wsPort"`
	EnginePort               uint64          `json:"enginePort" yaml:"enginePort"`
	OpenApiPorts             RpcPortMode     `json:"openApiPorts" yaml:"openApiPorts"`
	P2pPort                  uint64          `json:"p2pPort" yaml:"p2pPort"`
	AdditionalDockerNetworks string          `json:"additionalDockerNetworks" yaml:"additionalDockerNetworks"`

	Geth       *GethConfigSettings       `json:"geth" yaml:"geth"`
	Nethermind *NethermindConfigSettings `json:"nethermind" yaml:"nethermind"`
	Besu       *BesuConfigSettings       `json:"besu" yaml:"besu"`
	Reth       *RethConfigSettings       `json:"reth" yaml:"reth"`
}

// Create a new LocalExecutionConfig struct
func NewLocalExecutionConfig() *LocalExecutionConfig {
	cfg := &LocalExecutionConfig{}
	cfg.ID = hdconfig.Identifier(ids.EcID)

	//TODO: Confirm these

	// Options for ExecutionClient
	optionsEc := make([]hdconfig.ParameterOption[ExecutionClient], 4)
	optionsEc[0].Name = string(ExecutionClient_Geth)
	optionsEc[0].Description.Default = "Select if your external client is Geth."
	optionsEc[0].Value = ExecutionClient_Geth

	optionsEc[1].Name = string(ExecutionClient_Nethermind)
	optionsEc[1].Description.Default = "Select if your external client is Nethermind."
	optionsEc[1].Value = ExecutionClient_Nethermind

	optionsEc[2].Name = string(ExecutionClient_Besu)
	optionsEc[2].Description.Default = "Select if your external client is Besu."
	optionsEc[2].Value = ExecutionClient_Besu

	optionsEc[3].Name = string(ExecutionClient_Reth)
	optionsEc[3].Description.Default = "Select if your external client is Reth."
	optionsEc[3].Value = ExecutionClient_Reth

	cfg.ExecutionClient.ID = hdconfig.Identifier(ids.EcID)
	cfg.ExecutionClient.Name = "Execution Client"
	cfg.ExecutionClient.Description.Default = "Select which Execution client you would like to run."
	cfg.ExecutionClient.Options = optionsEc
	cfg.ExecutionClient.Default = ExecutionClient_Geth

	cfg.HttpPort.ID = hdconfig.Identifier(ids.HttpPortID)
	cfg.HttpPort.Name = "HTTP API Port"
	cfg.HttpPort.Description.Default = "The port your Execution client should use for its HTTP API endpoint (also known as HTTP RPC API endpoint)."
	cfg.HttpPort.Default = defaultLocalExecutionHttpPort

	cfg.WebsocketPort.ID = hdconfig.Identifier(ids.LocalEcWebsocketPortID)
	cfg.WebsocketPort.Name = "Websocket API Port"
	cfg.WebsocketPort.Description.Default = "The port your Execution client should use for its Websocket API endpoint (also known as Websocket RPC API endpoint)."
	cfg.WebsocketPort.Default = defaultLocalExecutionWebsocketPort

	cfg.EnginePort.ID = hdconfig.Identifier(ids.LocalEcEnginePortID)
	cfg.EnginePort.Name = "Engine API Port"
	cfg.EnginePort.Description.Default = "The port your Execution client should use for its Engine API endpoint (the endpoint the Beacon Node will connect to post-merge)."
	cfg.EnginePort.Default = defaultLocalExecutionEnginePort

	// Options for OpenApiPorts
	options := make([]hdconfig.ParameterOption[RpcPortMode], 3)
	options[0].Name = string(RpcPortMode_Closed)
	options[0].Description.Default = "Do not expose the RPC port outside of the Docker container."
	options[0].Value = RpcPortMode_Closed

	options[1].Name = string(RpcPortMode_OpenLocalhost)
	options[1].Description.Default = "Expose the RPC port to other processes on your machine."
	options[1].Value = RpcPortMode_OpenLocalhost

	options[2].Name = string(RpcPortMode_OpenExternal)
	options[2].Description.Default = "Expose the RPC port to other machines on your local network."
	options[2].Value = RpcPortMode_OpenExternal

	cfg.OpenApiPorts.ID = hdconfig.Identifier(ids.LocalEcOpenApiPortsID)
	cfg.OpenApiPorts.Name = "Expose API Ports"
	cfg.OpenApiPorts.Description.Default = "Expose the HTTP and Websocket API ports to other processes on your machine, or to your local network so other machines can access your Execution Client's API endpoints."
	cfg.OpenApiPorts.Options = options
	cfg.OpenApiPorts.Default = RpcPortMode_Closed

	cfg.P2pPort.ID = hdconfig.Identifier(ids.P2pPortID)
	cfg.P2pPort.Name = "P2P Port"
	cfg.P2pPort.Description.Default = "The port the Execution Client should use for P2P (blockchain) traffic to communicate with other nodes."
	cfg.P2pPort.Default = defaultLocalExecutionP2pPort

	// Create the subconfigs
	cfg.Geth = NewGethConfig()
	cfg.Nethermind = NewNethermindConfig()
	cfg.Besu = NewBesuConfig()
	cfg.Reth = NewRethConfig()

	return cfg
}

// Get the parameters for this config
func (cfg *LocalExecutionConfig) GetParameters() []hdconfig.IParameter {
	return []hdconfig.IParameter{
		&cfg.ExecutionClient,
		&cfg.HttpPort,
		&cfg.WebsocketPort,
		&cfg.EnginePort,
		&cfg.OpenApiPorts,
		&cfg.P2pPort,
	}
}

// Get the sections underneath this one
func (cfg *LocalExecutionConfig) GetSections() []hdconfig.ISection {
	return []hdconfig.ISection{
		cfg.Besu,
		cfg.Geth,
		cfg.Nethermind,
		cfg.Reth,
	}
}

// ==================
// === Templating ===
// ==================

// Get the Docker mapping for the selected API port mode
func (cfg *LocalExecutionConfigSettings) GetOpenApiPortMapping() string {
	rpcMode := RpcPortMode(cfg.OpenApiPorts)
	if !rpcMode.IsOpen() {
		return ""
	}
	httpMapping := rpcMode.DockerPortMapping(uint64(cfg.HttpPort))
	wsMapping := rpcMode.DockerPortMapping(uint64(cfg.WebsocketPort))
	return fmt.Sprintf(", \"%s\", \"%s\"", httpMapping, wsMapping)
}

// Gets the max peers of the selected EC
// Note that Reth treats the max peer count specially
func (cfg *LocalExecutionConfigSettings) GetMaxPeers() uint64 {
	switch cfg.ExecutionClient {
	case ExecutionClient_Geth:
		return cfg.Geth.MaxPeers
	case ExecutionClient_Nethermind:
		return cfg.Nethermind.MaxPeers
	case ExecutionClient_Besu:
		return cfg.Besu.MaxPeers
	case ExecutionClient_Reth:
		return cfg.Reth.MaxInboundPeers + cfg.Reth.MaxOutboundPeers
	default:
		panic(fmt.Sprintf("Unknown Execution Client %s", string(cfg.ExecutionClient)))
	}
}

// Get the container tag of the selected EC
func (cfg *LocalExecutionConfigSettings) GetContainerTag() string {
	switch cfg.ExecutionClient {
	case ExecutionClient_Geth:
		return cfg.Geth.ContainerTag
	case ExecutionClient_Nethermind:
		return cfg.Nethermind.ContainerTag
	case ExecutionClient_Besu:
		return cfg.Besu.ContainerTag
	case ExecutionClient_Reth:
		return cfg.Reth.ContainerTag
	default:
		panic(fmt.Sprintf("Unknown Execution Client %s", string(cfg.ExecutionClient)))
	}
}

// Gets the additional flags of the selected EC
func (cfg *LocalExecutionConfigSettings) GetAdditionalFlags() string {
	switch cfg.ExecutionClient {
	case ExecutionClient_Geth:
		return cfg.Geth.AdditionalFlags
	case ExecutionClient_Nethermind:
		return cfg.Nethermind.AdditionalFlags
	case ExecutionClient_Besu:
		return cfg.Besu.AdditionalFlags
	case ExecutionClient_Reth:
		return cfg.Reth.AdditionalFlags
	default:
		panic(fmt.Sprintf("Unknown Execution Client %s", string(cfg.ExecutionClient)))
	}
}
