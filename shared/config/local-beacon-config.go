package config

import (
	"fmt"

	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
)

// Common parameters shared by all of the Beacon Clients
type LocalBeaconConfig struct {
	// The selected BN
	BeaconNode hdconfig.ChoiceParameter[BeaconNode] //Parameter[BeaconNode]

	// The checkpoint sync URL if used
	CheckpointSyncProvider hdconfig.StringParameter

	// The port to use for gossip traffic
	P2pPort hdconfig.UintParameter

	// The port to expose the HTTP API on
	HttpPort hdconfig.UintParameter

	// Toggle for forwarding the HTTP API port outside of Docker
	OpenHttpPort hdconfig.ChoiceParameter[RpcPortMode] //Parameter[RpcPortMode]

	//Comma separated
	OpenPorts hdconfig.StringParameter

	//Comma separated
	AdditionalDockerNetworks hdconfig.StringParameter

	// Subconfigs
	Lighthouse *LighthouseBnConfig
	Lodestar   *LodestarBnConfig
	Nimbus     *NimbusBnConfig
	Prysm      *PrysmBnConfig
	Teku       *TekuBnConfig
}

type LocalBeaconConfigSettings struct {
	BeaconNode               BeaconNode  `json:"beaconNode"`
	CheckpointSyncProvider   string      `json:"checkpointSyncProvider"`
	P2pPort                  uint64      `json:"p2pPort"`
	HttpPort                 uint64      `json:"httpPort"`
	OpenHttpPort             RpcPortMode `json:"openHttpPort"`
	OpenPorts                []uint64    `json:"openPorts"`
	AdditionalDockerNetworks []string    `json:"additionalDockerNetworks"`

	Lighthouse *LighthouseBnConfigSettings `json:"lighthouse"`
	Lodestar   *LodestarBnConfigSettings   `json:"lodestar"`
	Nimbus     *NimbusBnConfigSettings     `json:"nimbus"`
	Prysm      *PrysmBnConfigSettings      `json:"prysm"`
	Teku       *TekuBnConfigSettings       `json:"teku"`
}

// Create a new LocalBeaconConfig struct
func NewLocalBeaconConfig() *LocalBeaconConfig {
	cfg := &LocalBeaconConfig{}

	cfg.Lighthouse = NewLighthouseBnConfig()
	cfg.Lodestar = NewLodestarBnConfig()
	cfg.Nimbus = NewNimbusBnConfig()
	cfg.Prysm = NewPrysmBnConfig()
	cfg.Teku = NewTekuBnConfig()

	cfg.CheckpointSyncProvider.ID = hdconfig.Identifier(ids.LocalBnCheckpointSyncUrlID)
	cfg.CheckpointSyncProvider.Name = "Checkpoint Sync URL"
	cfg.CheckpointSyncProvider.Description.Default = "If you would like to instantly sync using an existing Beacon node, enter its URL.\n" +
		"Example:  	https://checkpoint-sync.holesky.ethpandaops.io (for the Holesky Testnet).\n" +
		"Leave this blank if you want to sync normally from the start of the chain."
	// cfg.CheckpointSyncProvider.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	cfg.P2pPort.ID = hdconfig.Identifier(ids.P2pPortID)
	cfg.P2pPort.Name = "P2P Port"
	cfg.P2pPort.Description.Default = "The port to use for P2P (blockchain) traffic."
	// cfg.P2pPort.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	cfg.HttpPort.ID = hdconfig.Identifier(ids.HttpPortID)
	cfg.HttpPort.Name = "HTTP API Port"
	cfg.HttpPort.Description.Default = "The port your Beacon Node should run its HTTP API on."
	// cfg.HttpPort.AffectedContainers = []string{string(ContainerID_Daemon), string(ContainerID_BeaconNode), string(ContainerID_ValidatorClient), string(ContainerID_Prometheus)}

	// Options for OpenHttpPort
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

	cfg.OpenHttpPort.ID = hdconfig.Identifier(ids.OpenHttpPortsID)
	cfg.OpenHttpPort.Name = "Expose API Port"
	cfg.OpenHttpPort.Description.Default = "Select an option to expose your Beacon Node's API port to your localhost or external hosts on the network, so other machines can access it too."
	// cfg.OpenHttpPort.AffectedContainers = []string{string(ContainerID_BeaconNode)}
	cfg.OpenHttpPort.Options = options

	// Options for BeaconNode
	optionsBeaconNode := make([]hdconfig.ParameterOption[BeaconNode], 5)
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

	cfg.BeaconNode.ID = hdconfig.Identifier(ids.BnID)
	cfg.BeaconNode.Name = "Beacon Node"
	cfg.BeaconNode.Description.Default = "Select which Beacon Node client you would like to use."
	// cfg.BeaconNode.AffectedContainers = []string{string(ContainerID_Daemon), string(ContainerID_BeaconNode), string(ContainerID_ValidatorClient)}
	cfg.BeaconNode.Options = optionsBeaconNode

	return cfg
}

// The title for the config
func (cfg *LocalBeaconConfig) GetTitle() string {
	return "Local Beacon Node"
}

// Get the parameters for this config
func (cfg *LocalBeaconConfig) GetParameters() []hdconfig.IParameter {
	return []hdconfig.IParameter{
		&cfg.BeaconNode,
		&cfg.CheckpointSyncProvider,
		&cfg.P2pPort,
		&cfg.HttpPort,
		&cfg.OpenHttpPort,
	}
}

// Get the sections underneath this one
func (cfg *LocalBeaconConfig) GetSections() map[string]hdconfig.ISection {
	return map[string]hdconfig.ISection{
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
func (cfg *LocalBeaconConfigSettings) GetOpenApiPortMapping() []string {
	bnOpenPorts := make([]string, 0)

	// Handle the standard HTTP API port
	apiPortMode := RpcPortMode(cfg.OpenHttpPort)
	if apiPortMode.IsOpen() {
		apiPort := uint64(cfg.HttpPort)
		bnOpenPorts = append(bnOpenPorts, apiPortMode.DockerPortMapping(apiPort))
	}

	// Handle Prysm's RPC port
	if cfg.BeaconNode == BeaconNode_Prysm {
		prysmRpcPortMode := RpcPortMode(cfg.Prysm.OpenRpcPort)
		if prysmRpcPortMode.IsOpen() {
			prysmRpcPort := uint64(cfg.Prysm.RpcPort)
			bnOpenPorts = append(bnOpenPorts, prysmRpcPortMode.DockerPortMapping(prysmRpcPort))
		}
	}
	return bnOpenPorts
}

// Gets the max peers of the selected EC
func (cfg *LocalBeaconConfigSettings) GetMaxPeers() uint64 {
	switch cfg.BeaconNode {
	case BeaconNode_Lighthouse:
		return cfg.Lighthouse.MaxPeers
	case BeaconNode_Lodestar:
		return cfg.Lodestar.MaxPeers
	case BeaconNode_Nimbus:
		return cfg.Nimbus.MaxPeers
	case BeaconNode_Prysm:
		return cfg.Prysm.MaxPeers
	case BeaconNode_Teku:
		return cfg.Teku.MaxPeers
	default:
		panic(fmt.Sprintf("Unknown Beacon Node %s", string(cfg.BeaconNode)))
	}
}

// Get the container tag of the selected BN
func (cfg *LocalBeaconConfigSettings) GetContainerTag() string {
	switch cfg.BeaconNode {
	case BeaconNode_Lighthouse:
		return cfg.Lighthouse.ContainerTag
	case BeaconNode_Lodestar:
		return cfg.Lodestar.ContainerTag
	case BeaconNode_Nimbus:
		return cfg.Nimbus.ContainerTag
	case BeaconNode_Prysm:
		return cfg.Prysm.ContainerTag
	case BeaconNode_Teku:
		return cfg.Teku.ContainerTag
	default:
		panic(fmt.Sprintf("Unknown Beacon Node %s", string(cfg.BeaconNode)))
	}
}

// Gets the additional flags of the selected BN
func (cfg *LocalBeaconConfigSettings) GetAdditionalFlags() string {
	switch cfg.BeaconNode {
	case BeaconNode_Lighthouse:
		return cfg.Lighthouse.AdditionalFlags
	case BeaconNode_Lodestar:
		return cfg.Lodestar.AdditionalFlags
	case BeaconNode_Nimbus:
		return cfg.Nimbus.AdditionalFlags
	case BeaconNode_Prysm:
		return cfg.Prysm.AdditionalFlags
	case BeaconNode_Teku:
		return cfg.Teku.AdditionalFlags
	default:
		panic(fmt.Sprintf("Unknown Beacon Node %s", string(cfg.BeaconNode)))
	}
}
