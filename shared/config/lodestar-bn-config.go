package config

import (
	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	"github.com/nodeset-org/hyperdrive/modules/config"
)

const (
	lodestarBnTag string = "chainsafe/lodestar:v1.22.0"
)

// Configuration for the Lodestar BN
type LodestarBnConfig struct {
	// The max number of P2P peers to connect to
	MaxPeers config.UintParameter

	// The Docker Hub tag for Lodestar BN
	ContainerTag config.StringParameter

	// Custom command line flags for the BN
	AdditionalFlags config.StringParameter
}

// Generates a new Lodestar BN configuration
func NewLodestarBnConfig() *LodestarBnConfig {
	cfg := &LodestarBnConfig{}
	cfg.MaxPeers.ID = config.Identifier(ids.MaxPeersID)
	cfg.MaxPeers.Name = "Max Peers"
	cfg.MaxPeers.Description.Default = "The maximum number of peers your client should try to maintain. You can try lowering this if you have a low-resource system or a constrained network."
	cfg.MaxPeers.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	cfg.ContainerTag.ID = config.Identifier(ids.ContainerTagID)
	cfg.ContainerTag.Name = "Container Tag"
	cfg.ContainerTag.Description.Default = "The tag name of the Lodestar container from Docker Hub you want to use for the Beacon Node."
	cfg.ContainerTag.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	cfg.AdditionalFlags.ID = config.Identifier(ids.AdditionalFlagsID)
	cfg.AdditionalFlags.Name = "Additional Flags"
	cfg.AdditionalFlags.Description.Default = "Additional custom command line flags you want to pass Lodestar's Beacon Client, to take advantage of other settings that aren't covered here."
	cfg.AdditionalFlags.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	return cfg
}

// The title for the config
func (cfg *LodestarBnConfig) GetTitle() string {
	return "Lodestar Beacon Node"
}

// Get the parameters for this config
func (cfg *LodestarBnConfig) GetParameters() []IParameter {
	return []IParameter{
		&cfg.MaxPeers,
		&cfg.ContainerTag,
		&cfg.AdditionalFlags,
	}
}

// Get the sections underneath this one
func (cfg *LodestarBnConfig) GetSubconfigs() map[string]IConfigSection {
	return map[string]IConfigSection{}
}
