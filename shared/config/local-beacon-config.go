package config

import (
	"fmt"

	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	"github.com/nodeset-org/hyperdrive/modules/config"
)

// Common parameters shared by all of the Beacon Clients
type LocalBeaconConfig struct {
	// The selected BN
	BeaconNode config.ChoiceParameter[BeaconNode] //Parameter[BeaconNode]

	// The checkpoint sync URL if used
	CheckpointSyncProvider config.StringParameter

	// The port to use for gossip traffic
	P2pPort config.UintParameter

	// The port to expose the HTTP API on
	HttpPort config.UintParameter

	// Toggle for forwarding the HTTP API port outside of Docker
	OpenHttpPort config.ChoiceParameter[RpcPortMode] //Parameter[RpcPortMode]

	// Subconfigs
	Lighthouse *LighthouseBnConfig
	Lodestar   *LodestarBnConfig
	Nimbus     *NimbusBnConfig
	Prysm      *PrysmBnConfig
	Teku       *TekuBnConfig
}

// Create a new LocalBeaconConfig struct
func NewLocalBeaconConfig() *LocalBeaconConfig {
	cfg := &LocalBeaconConfig{}

	cfg.Lighthouse = NewLighthouseBnConfig()
	cfg.Lodestar = NewLodestarBnConfig()
	cfg.Nimbus = NewNimbusBnConfig()
	cfg.Prysm = NewPrysmBnConfig()
	cfg.Teku = NewTekuBnConfig()

	cfg.CheckpointSyncProvider.ID = config.Identifier(ids.LocalBnCheckpointSyncUrlID)
	cfg.CheckpointSyncProvider.Name = "Checkpoint Sync URL"
	cfg.CheckpointSyncProvider.Description.Default = "If you would like to instantly sync using an existing Beacon node, enter its URL.\n" +
		"Example:  	https://checkpoint-sync.holesky.ethpandaops.io (for the Holesky Testnet).\n" +
		"Leave this blank if you want to sync normally from the start of the chain."
	cfg.CheckpointSyncProvider.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	cfg.P2pPort.ID = config.Identifier(ids.P2pPortID)
	cfg.P2pPort.Name = "P2P Port"
	cfg.P2pPort.Description.Default = "The port to use for P2P (blockchain) traffic."
	cfg.P2pPort.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	cfg.HttpPort.ID = config.Identifier(ids.HttpPortID)
	cfg.HttpPort.Name = "HTTP API Port"
	cfg.HttpPort.Description.Default = "The port your Beacon Node should run its HTTP API on."
	cfg.HttpPort.AffectedContainers = []string{string(ContainerID_Daemon), string(ContainerID_BeaconNode), string(ContainerID_ValidatorClient), string(ContainerID_Prometheus)}

	// Options for OpenHttpPort
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

	cfg.OpenHttpPort.ID = config.Identifier(ids.OpenHttpPortsID)
	cfg.OpenHttpPort.Name = "Expose API Port"
	cfg.OpenHttpPort.Description.Default = "Select an option to expose your Beacon Node's API port to your localhost or external hosts on the network, so other machines can access it too."
	cfg.OpenHttpPort.AffectedContainers = []string{string(ContainerID_BeaconNode)}
	cfg.OpenHttpPort.Options = options

	// Options for BeaconNode
	optionsBeaconNode := make([]config.ParameterOption[BeaconNode], 5)
	optionsBeaconNode[0].Name = "Lighthouse"
	optionsBeaconNode[0].Description.Default = "Lighthouse is a Beacon Node with a heavy focus on speed and security. The team behind it, Sigma Prime, is an information security and software engineering firm who have funded Lighthouse along with the Ethereum Foundation, Consensys, and private individuals. Lighthouse is built in Rust and offered under an Apache 2.0 License."
	optionsBeaconNode[0].Value = BeaconNode_Lighthouse

	optionsBeaconNode[1].Name = "Lodestar"
	optionsBeaconNode[1].Description.Default = "Lodestar is the fifth open-source Ethereum Beacon Node. It is written in Typescript maintained by ChainSafe Systems. Lodestar, their flagship product, is a production-capable Beacon Chain and Validator Client uniquely situated as the go-to for researchers and developers for rapid prototyping and browser usage."
	optionsBeaconNode[1].Value = BeaconNode_Lodestar

	optionsBeaconNode[2].Name = "Nimbus"
	optionsBeaconNode[2].Description.Default = "Nimbus is a Beacon Node implementation that strives to be as lightweight as possible in terms of resources used. This allows it to perform well on embedded systems, resource-restricted devices -- including Raspberry Pis and mobile devices -- and multi-purpose servers."
	optionsBeaconNode[2].Value = BeaconNode_Nimbus

	optionsBeaconNode[3].Name = "Prysm"
	optionsBeaconNode[3].Description.Default = "Prysm is a Go implementation of Ethereum Consensus protocol with a focus on usability, security, and reliability. Prysm is developed by Prysmatic Labs, a company with the sole focus on the development of their client. Prysm is written in Go and released under a GPL-3.0 license."
	optionsBeaconNode[3].Value = BeaconNode_Prysm

	optionsBeaconNode[4].Name = "Teku"
	optionsBeaconNode[4].Description.Default = "PegaSys Teku (formerly known as Artemis) is a Java-based Ethereum 2.0 client designed & built to meet institutional needs and security requirements. PegaSys is an arm of ConsenSys dedicated to building enterprise-ready clients and tools for interacting with the core Ethereum platform. Teku is Apache 2 licensed and written in Java, a language notable for its maturity & ubiquity."
	optionsBeaconNode[4].Value = BeaconNode_Teku

	cfg.BeaconNode.ID = config.Identifier(ids.BnID)
	cfg.BeaconNode.Name = "Beacon Node"
	cfg.BeaconNode.Description.Default = "Select which Beacon Node client you would like to use."
	cfg.BeaconNode.AffectedContainers = []string{string(ContainerID_Daemon), string(ContainerID_BeaconNode), string(ContainerID_ValidatorClient)}
	cfg.BeaconNode.Options = optionsBeaconNode

	return cfg
}

// The title for the config
func (cfg *LocalBeaconConfig) GetTitle() string {
	return "Local Beacon Node"
}

// Get the parameters for this config
func (cfg *LocalBeaconConfig) GetParameters() []config.IParameter {
	return []config.IParameter{
		&cfg.BeaconNode,
		&cfg.CheckpointSyncProvider,
		&cfg.P2pPort,
		&cfg.HttpPort,
		&cfg.OpenHttpPort,
	}
}

// Get the sections underneath this one
func (cfg *LocalBeaconConfig) GetSubconfigs() map[string]config.ISection {
	return map[string]config.ISection{
		ids.LocalBnLighthouseID: cfg.Lighthouse,
		ids.LocalBnLodestarID:   cfg.Lodestar,
		ids.LocalBnNimbusID:     cfg.Nimbus,
		ids.LocalBnPrysmID:      cfg.Prysm,
		ids.LocalBnTekuID:       cfg.Teku,
	}
}

// ==================
// === Templating ===
// ==================

// Get the Docker mapping for the selected API port mode
func (cfg *LocalBeaconConfig) GetOpenApiPortMapping() []string {
	bnOpenPorts := make([]string, 0)

	// Handle the standard HTTP API port
	apiPortMode := cfg.OpenHttpPort.Value
	if apiPortMode.IsOpen() {
		apiPort := cfg.HttpPort.Value
		bnOpenPorts = append(bnOpenPorts, apiPortMode.DockerPortMapping(apiPort))
	}

	// Handle Prysm's RPC port
	if cfg.BeaconNode.Value == BeaconNode_Prysm {
		prysmRpcPortMode := cfg.Prysm.OpenRpcPort.Value
		if prysmRpcPortMode.IsOpen() {
			prysmRpcPort := cfg.Prysm.RpcPort.Value
			bnOpenPorts = append(bnOpenPorts, prysmRpcPortMode.DockerPortMapping(prysmRpcPort))
		}
	}
	return bnOpenPorts
}

// Gets the max peers of the selected EC
func (cfg *LocalBeaconConfig) GetMaxPeers() uint64 {
	switch cfg.BeaconNode.Value {
	case BeaconNode_Lighthouse:
		return cfg.Lighthouse.MaxPeers.Value
	case BeaconNode_Lodestar:
		return cfg.Lodestar.MaxPeers.Value
	case BeaconNode_Nimbus:
		return cfg.Nimbus.MaxPeers.Value
	case BeaconNode_Prysm:
		return cfg.Prysm.MaxPeers.Value
	case BeaconNode_Teku:
		return cfg.Teku.MaxPeers.Value
	default:
		panic(fmt.Sprintf("Unknown Beacon Node %s", string(cfg.BeaconNode.Value)))
	}
}

// Get the container tag of the selected BN
func (cfg *LocalBeaconConfig) GetContainerTag() string {
	switch cfg.BeaconNode.Value {
	case BeaconNode_Lighthouse:
		return cfg.Lighthouse.ContainerTag.Value
	case BeaconNode_Lodestar:
		return cfg.Lodestar.ContainerTag.Value
	case BeaconNode_Nimbus:
		return cfg.Nimbus.ContainerTag.Value
	case BeaconNode_Prysm:
		return cfg.Prysm.ContainerTag.Value
	case BeaconNode_Teku:
		return cfg.Teku.ContainerTag.Value
	default:
		panic(fmt.Sprintf("Unknown Beacon Node %s", string(cfg.BeaconNode.Value)))
	}
}

// Gets the additional flags of the selected BN
func (cfg *LocalBeaconConfig) GetAdditionalFlags() string {
	switch cfg.BeaconNode.Value {
	case BeaconNode_Lighthouse:
		return cfg.Lighthouse.AdditionalFlags.Value
	case BeaconNode_Lodestar:
		return cfg.Lodestar.AdditionalFlags.Value
	case BeaconNode_Nimbus:
		return cfg.Nimbus.AdditionalFlags.Value
	case BeaconNode_Prysm:
		return cfg.Prysm.AdditionalFlags.Value
	case BeaconNode_Teku:
		return cfg.Teku.AdditionalFlags.Value
	default:
		panic(fmt.Sprintf("Unknown Beacon Node %s", string(cfg.BeaconNode.Value)))
	}
}
