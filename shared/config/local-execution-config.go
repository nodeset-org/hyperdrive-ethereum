package config

import (
	"fmt"

	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	"github.com/nodeset-org/hyperdrive/modules/config"
)

// Configuration for the Execution client
type LocalExecutionConfig struct {
	// The selected EC
	ExecutionClient config.ChoiceParameter[ExecutionClient] //Parameter[ExecutionClient]

	// The HTTP API port
	HttpPort config.UintParameter

	// The Websocket API port
	WebsocketPort config.UintParameter

	// The Engine API port
	EnginePort config.UintParameter

	// Toggle for forwarding the HTTP API port outside of Docker
	OpenApiPorts config.ChoiceParameter[RpcPortMode] //Parameter[RpcPortMode]

	// P2P traffic port
	P2pPort config.UintParameter

	// Subconfigs
	Geth       *GethConfig
	Nethermind *NethermindConfig
	Besu       *BesuConfig
	Reth       *RethConfig
}

// Create a new LocalExecutionConfig struct
func NewLocalExecutionConfig() *LocalExecutionConfig {
	cfg := &LocalExecutionConfig{}

	//TODO: Confirm these

	// Options for ExecutionClient
	optionsEc := make([]config.ParameterOption[ExecutionClient], 4)
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

	cfg.ExecutionClient.ID = config.Identifier(ids.EcID)
	cfg.ExecutionClient.Name = "Execution Client"
	cfg.ExecutionClient.Description.Default = "Select which Execution client you would like to run."
	cfg.ExecutionClient.AffectedContainers = []string{string(ContainerID_ExecutionClient), string(ContainerID_ValidatorClient)}
	cfg.ExecutionClient.Options = optionsEc

	cfg.HttpPort.ID = config.Identifier(ids.HttpPortID)
	cfg.HttpPort.Name = "HTTP API Port"
	cfg.HttpPort.Description.Default = "The port your Execution client should use for its HTTP API endpoint (also known as HTTP RPC API endpoint)."
	cfg.HttpPort.AffectedContainers = []string{string(ContainerID_Daemon), string(ContainerID_ExecutionClient), string(ContainerID_BeaconNode)}

	cfg.WebsocketPort.ID = config.Identifier(ids.LocalEcWebsocketPortID)
	cfg.WebsocketPort.Name = "Websocket API Port"
	cfg.WebsocketPort.Description.Default = "The port your Execution client should use for its Websocket API endpoint (also known as Websocket RPC API endpoint)."
	cfg.WebsocketPort.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	cfg.EnginePort.ID = config.Identifier(ids.LocalEcEnginePortID)
	cfg.EnginePort.Name = "Engine API Port"
	cfg.EnginePort.Description.Default = "The port your Execution client should use for its Engine API endpoint (the endpoint the Beacon Node will connect to post-merge)."
	cfg.EnginePort.AffectedContainers = []string{string(ContainerID_ExecutionClient), string(ContainerID_BeaconNode)}

	// Options for OpenApiPorts
	options := make([]config.ParameterOption[RpcPortMode], 3)
	options[0].Name = string(RpcPortMode_Closed)
	options[0].Description.Default = "Do not expose the RPC port outside of the Docker container."
	options[0].Value = RpcPortMode_Closed

	options[1].Name = string(RpcPortMode_OpenLocalhost)
	options[1].Description.Default = "Expose the RPC port to other processes on your machine."
	options[1].Value = RpcPortMode_OpenLocalhost

	options[2].Name = string(RpcPortMode_OpenExternal)
	options[2].Description.Default = "Expose the RPC port to other machines on your local network."
	options[2].Value = RpcPortMode_OpenExternal

	cfg.OpenApiPorts.ID = config.Identifier(ids.LocalEcOpenApiPortsID)
	cfg.OpenApiPorts.Name = "Expose API Ports"
	cfg.OpenApiPorts.Description.Default = "Expose the HTTP and Websocket API ports to other processes on your machine, or to your local network so other machines can access your Execution Client's API endpoints."
	cfg.OpenApiPorts.AffectedContainers = []string{string(ContainerID_ExecutionClient)}
	cfg.OpenApiPorts.Options = options

	cfg.P2pPort.ID = config.Identifier(ids.P2pPortID)
	cfg.P2pPort.Name = "P2P Port"
	cfg.P2pPort.Description.Default = "The port the Execution Client should use for P2P (blockchain) traffic to communicate with other nodes."
	cfg.P2pPort.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	// Create the subconfigs
	cfg.Geth = NewGethConfig()
	cfg.Nethermind = NewNethermindConfig()
	cfg.Besu = NewBesuConfig()
	cfg.Reth = NewRethConfig()

	return cfg
}

// Get the title for the config
func (cfg *LocalExecutionConfig) GetTitle() string {
	return "Local Execution Client"
}

// Get the parameters for this config
func (cfg *LocalExecutionConfig) GetParameters() []config.IParameter {
	return []config.IParameter{
		&cfg.ExecutionClient,
		&cfg.HttpPort,
		&cfg.WebsocketPort,
		&cfg.EnginePort,
		&cfg.OpenApiPorts,
		&cfg.P2pPort,
	}
}

// Get the sections underneath this one
func (cfg *LocalExecutionConfig) GetSubconfigs() map[string]IConfigSection {
	return map[string]IConfigSection{
		ids.LocalEcBesuID:       cfg.Besu,
		ids.LocalEcGethID:       cfg.Geth,
		ids.LocalEcNethermindID: cfg.Nethermind,
		ids.LocalEcRethID:       cfg.Reth,
	}
}

// ==================
// === Templating ===
// ==================

// Get the Docker mapping for the selected API port mode
func (cfg *LocalExecutionConfig) GetOpenApiPortMapping() string {
	rpcMode := cfg.OpenApiPorts.Value
	if !rpcMode.IsOpen() {
		return ""
	}
	httpMapping := rpcMode.DockerPortMapping(cfg.HttpPort.Value)
	wsMapping := rpcMode.DockerPortMapping(cfg.WebsocketPort.Value)
	return fmt.Sprintf(", \"%s\", \"%s\"", httpMapping, wsMapping)
}

// Gets the max peers of the selected EC
// Note that Reth treats the max peer count specially
func (cfg *LocalExecutionConfig) GetMaxPeers() uint64 {
	switch cfg.ExecutionClient.Value {
	case ExecutionClient_Geth:
		return cfg.Geth.MaxPeers.Value
	case ExecutionClient_Nethermind:
		return cfg.Nethermind.MaxPeers.Value
	case ExecutionClient_Besu:
		return cfg.Besu.MaxPeers.Value
	case ExecutionClient_Reth:
		return cfg.Reth.MaxInboundPeers.Value + cfg.Reth.MaxOutboundPeers.Value
	default:
		panic(fmt.Sprintf("Unknown Execution Client %s", string(cfg.ExecutionClient.Value)))
	}
}

// Get the container tag of the selected EC
func (cfg *LocalExecutionConfig) GetContainerTag() string {
	switch cfg.ExecutionClient.Value {
	case ExecutionClient_Geth:
		return cfg.Geth.ContainerTag.Value
	case ExecutionClient_Nethermind:
		return cfg.Nethermind.ContainerTag.Value
	case ExecutionClient_Besu:
		return cfg.Besu.ContainerTag.Value
	case ExecutionClient_Reth:
		return cfg.Reth.ContainerTag.Value
	default:
		panic(fmt.Sprintf("Unknown Execution Client %s", string(cfg.ExecutionClient.Value)))
	}
}

// Gets the additional flags of the selected EC
func (cfg *LocalExecutionConfig) GetAdditionalFlags() string {
	switch cfg.ExecutionClient.Value {
	case ExecutionClient_Geth:
		return cfg.Geth.AdditionalFlags.Value
	case ExecutionClient_Nethermind:
		return cfg.Nethermind.AdditionalFlags.Value
	case ExecutionClient_Besu:
		return cfg.Besu.AdditionalFlags.Value
	case ExecutionClient_Reth:
		return cfg.Reth.AdditionalFlags.Value
	default:
		panic(fmt.Sprintf("Unknown Execution Client %s", string(cfg.ExecutionClient.Value)))
	}
}
