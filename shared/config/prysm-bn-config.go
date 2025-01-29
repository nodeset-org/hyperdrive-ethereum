package config

import (
	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
)

const (
	// Tags
	prysmBnTag string = "rocketpool/prysm:v5.1.2"
)

// Configuration for the Prysm BN
type PrysmBnConfig struct {
	// The max number of P2P peers to connect to
	MaxPeers hdconfig.UintParameter

	// The RPC port for BN / VC connections
	RpcPort hdconfig.UintParameter

	// Toggle for forwarding the RPC API outside of Docker
	OpenRpcPort hdconfig.ChoiceParameter[RpcPortMode] //Parameter[RpcPortMode]

	// The Docker Hub tag for the Prysm BN
	ContainerTag hdconfig.StringParameter

	// Custom command line flags for the BN
	AdditionalFlags hdconfig.StringParameter
}

// Generates a new Prysm BN configuration
func NewPrysmBnConfig() *PrysmBnConfig {
	cfg := &PrysmBnConfig{}

	cfg.MaxPeers.ID = hdconfig.Identifier(ids.MaxPeersID)
	cfg.MaxPeers.Name = "Max Peers"
	cfg.MaxPeers.Description.Default = "The maximum number of peers your client should try to maintain. You can try lowering this if you have a low-resource system or a constrained network."
	cfg.MaxPeers.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	cfg.RpcPort.ID = hdconfig.Identifier(ids.PrysmRpcPortID)
	cfg.RpcPort.Name = "RPC Port"
	cfg.RpcPort.Description.Default = "The port Prysm should run its JSON-RPC API on."
	cfg.RpcPort.AffectedContainers = []string{string(ContainerID_BeaconNode), string(ContainerID_ValidatorClient)}

	// Options for OpenRpcPort
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

	cfg.OpenRpcPort.ID = hdconfig.Identifier(ids.PrysmOpenRpcPortID)
	cfg.OpenRpcPort.Name = "Expose RPC Port"
	cfg.OpenRpcPort.Description.Default = "Expose Prysm's JSON-RPC port to other processes on your machine, or to your local network so other machines can access it too."
	cfg.OpenRpcPort.AffectedContainers = []string{string(ContainerID_BeaconNode)}
	cfg.OpenRpcPort.Options = options

	cfg.ContainerTag.ID = hdconfig.Identifier(ids.ContainerTagID)
	cfg.ContainerTag.Name = "Container Tag"
	cfg.ContainerTag.Description.Default = "The tag name of the Prysm Beacon Node container on Docker Hub you want to use for the Beacon Node."
	cfg.ContainerTag.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	cfg.AdditionalFlags.ID = hdconfig.Identifier(ids.AdditionalFlagsID)
	cfg.AdditionalFlags.Name = "Additional Flags"
	cfg.AdditionalFlags.Description.Default = "Additional custom command line flags you want to pass Prysm's Beacon Node, to take advantage of other settings that aren't covered here."
	cfg.AdditionalFlags.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	return cfg
}

// The title for the config
func (cfg *PrysmBnConfig) GetTitle() string {
	return "Prysm Beacon Node"
}

// Get the parameters for this config
func (cfg *PrysmBnConfig) GetParameters() []hdconfig.IParameter {
	return []hdconfig.IParameter{
		&cfg.MaxPeers,
		&cfg.RpcPort,
		&cfg.OpenRpcPort,
		&cfg.ContainerTag,
		&cfg.AdditionalFlags,
	}
}

// Get the sections underneath this one
func (cfg *PrysmBnConfig) GetSections() map[string]hdconfig.ISection {
	return map[string]hdconfig.ISection{}
}
