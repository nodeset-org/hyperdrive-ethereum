package config

import (
	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	"github.com/nodeset-org/hyperdrive/modules/config"
	nmcconfig "github.com/rocket-pool/node-manager-core/config"
)

const (
	// Tags
	lighthouseBnTag string = "sigp/lighthouse:v5.3.0"
)

// Configuration for the Lighthouse BN
type LighthouseBnConfig struct {
	// The port to use for gossip traffic using the QUIC protocol
	P2pQuicPort config.UintParameter

	// The max number of P2P peers to connect to
	MaxPeers config.UintParameter

	// The Docker Hub tag for Lighthouse BN
	ContainerTag config.StringParameter

	// Custom command line flags for the BN
	AdditionalFlags config.StringParameter
}

// Generates a new Lighthouse BN configuration
func NewLighthouseBnConfig() *LighthouseBnConfig {
	cfg := &LighthouseBnConfig{}

	cfg.MaxPeers.ID = config.Identifier(ids.MaxPeersID)
	cfg.MaxPeers.Name = "Max Peers"
	cfg.MaxPeers.Description.Default = "The maximum number of peers your client should try to maintain. You can try lowering this if you have a low-resource system or a constrained network."
	cfg.MaxPeers.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	cfg.ContainerTag.ID = config.Identifier(ids.ContainerTagID)
	cfg.ContainerTag.Name = "Container Tag"
	cfg.ContainerTag.Description.Default = "The tag name of the Lighthouse container from Docker Hub you want to use for the Beacon Node."
	cfg.ContainerTag.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	cfg.AdditionalFlags.ID = config.Identifier(ids.AdditionalFlagsID)
	cfg.AdditionalFlags.Name = "Additional Flags"
	cfg.AdditionalFlags.Description.Default = "Additional custom command line flags you want to pass Lighthouse's Beacon Node, to take advantage of other settings that aren't covered here."
	cfg.AdditionalFlags.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	cfg.P2pQuicPort.ID = config.Identifier(ids.P2pQuicPortID)
	cfg.P2pQuicPort.Name = "P2pQuicPort"
	cfg.P2pQuicPort.Description.Default = "TODO"
	cfg.P2pQuicPort.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	return cfg
}

// The title for the config
func (cfg *LighthouseBnConfig) GetTitle() string {
	return "Lighthouse Beacon Node"
}

// Get the parameters for this config
func (cfg *LighthouseBnConfig) GetParameters() []config.IParameter {
	return []config.IParameter{
		&cfg.MaxPeers,
		&cfg.P2pQuicPort,
		&cfg.ContainerTag,
		&cfg.AdditionalFlags,
	}
}

// Get the sections underneath this one
func (cfg *LighthouseBnConfig) GetSubconfigs() map[string]nmcconfig.IConfigSection {
	return map[string]nmcconfig.IConfigSection{}
}

// TODO (HN)
func (cfg *LighthouseBnConfig) GetDescription() config.DynamicProperty[string] {
	return config.DynamicProperty[string]{}
}

// TODO (HN)
func (cfg *LighthouseBnConfig) GetDisabled() config.DynamicProperty[bool] {
	return config.DynamicProperty[bool]{}
}

// TODO (HN)
func (cfg *LighthouseBnConfig) GetHidden() config.DynamicProperty[bool] {
	return config.DynamicProperty[bool]{}
}

// TODO (HN)
func (cfg *LighthouseBnConfig) GetID() config.Identifier {
	return config.Identifier("")
}

// TODO (HN)
func (cfg *LighthouseBnConfig) GetName() string {
	return ""
}

// TODO (HN)
func (cfg *LighthouseBnConfig) GetSections() []config.ISection {
	return []config.ISection{}
}
