package config

import (
	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
)

const (
	lodestarBnTag string = "chainsafe/lodestar:v1.22.0"
)

// Configuration for the Lodestar BN
type LodestarBnConfig struct {
	hdconfig.SectionHeader

	// The max number of P2P peers to connect to
	MaxPeers hdconfig.UintParameter

	// The Docker Hub tag for Lodestar BN
	ContainerTag hdconfig.StringParameter

	// Custom command line flags for the BN
	AdditionalFlags hdconfig.StringParameter
}

type LodestarBnConfigSettings struct {
	MaxPeers        uint64 `json:"maxPeers" yaml:"maxPeers"`
	ContainerTag    string `json:"containerTag" yaml:"containerTag"`
	AdditionalFlags string `json:"additionalFlags" yaml:"additionalFlags"`
}

// Generates a new Lodestar BN configuration
func NewLodestarBnConfig() *LodestarBnConfig {
	cfg := &LodestarBnConfig{}
	cfg.ID = hdconfig.Identifier(ids.LocalBnLodestarID)
	cfg.MaxPeers.ID = hdconfig.Identifier(ids.MaxPeersID)
	cfg.MaxPeers.Name = "Max Peers"
	cfg.MaxPeers.Description.Default = "The maximum number of peers your client should try to maintain. You can try lowering this if you have a low-resource system or a constrained network."
	// cfg.MaxPeers.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	cfg.ContainerTag.ID = hdconfig.Identifier(ids.ContainerTagID)
	cfg.ContainerTag.Name = "Container Tag"
	cfg.ContainerTag.Description.Default = "The tag name of the Lodestar container from Docker Hub you want to use for the Beacon Node."
	// cfg.ContainerTag.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	cfg.AdditionalFlags.ID = hdconfig.Identifier(ids.AdditionalFlagsID)
	cfg.AdditionalFlags.Name = "Additional Flags"
	cfg.AdditionalFlags.Description.Default = "Additional custom command line flags you want to pass Lodestar's Beacon Client, to take advantage of other settings that aren't covered here."
	// cfg.AdditionalFlags.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	return cfg
}

// The title for the config
func (cfg *LodestarBnConfig) GetTitle() string {
	return "Lodestar"
}

// Get the parameters for this config
func (cfg *LodestarBnConfig) GetParameters() []hdconfig.IParameter {
	return []hdconfig.IParameter{
		&cfg.MaxPeers,
		&cfg.ContainerTag,
		&cfg.AdditionalFlags,
	}
}

// Get the sections underneath this one
func (cfg *LodestarBnConfig) GetSections() []hdconfig.ISection {
	return []hdconfig.ISection{}
}

// TODO: Talk to Joe about these funcs required for ISection
func (cfg *LodestarBnConfig) GetDescription() hdconfig.DynamicProperty[string] {
	return hdconfig.DynamicProperty[string]{}
}

func (cfg *LodestarBnConfig) GetDisabled() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}

func (cfg *LodestarBnConfig) GetHidden() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}

func (cfg *LodestarBnConfig) GetID() hdconfig.Identifier {
	return hdconfig.Identifier(ids.LocalBnLodestarID)
}

func (cfg *LodestarBnConfig) GetName() string {
	return "Lodestar"
}
