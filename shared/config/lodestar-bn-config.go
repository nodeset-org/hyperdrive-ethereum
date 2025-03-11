package config

import (
	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
)

const (
	lodestarBnTag string = "chainsafe/lodestar:v1.22.0"

	lodestarMaxPeers uint = 100
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
	cfg.Name = "Lodestar"
	cfg.Description.Default = "Lodestar is a full-featured Ethereum 2.0 client written in TypeScript. It is known for its ease of use and high performance."
	// cfg.Hidden.Default = true
	// cfg.Hidden.Template = "{{if eq .GetValue \"localBeaconClient/beaconNode\" \"lodestar\"}}false{{else}}true{{end}}"

	cfg.MaxPeers.ID = hdconfig.Identifier(ids.MaxPeersID)
	cfg.MaxPeers.Name = "Max Peers"
	cfg.MaxPeers.Description.Default = "The maximum number of peers your client should try to maintain. You can try lowering this if you have a low-resource system or a constrained network."
	cfg.MaxPeers.Default = uint64(lodestarMaxPeers)

	cfg.ContainerTag.ID = hdconfig.Identifier(ids.ContainerTagID)
	cfg.ContainerTag.Name = "Container Tag"
	cfg.ContainerTag.Description.Default = "The tag name of the Lodestar container from Docker Hub you want to use for the Beacon Node."
	cfg.ContainerTag.Default = lodestarBnTag

	cfg.AdditionalFlags.ID = hdconfig.Identifier(ids.AdditionalFlagsID)
	cfg.AdditionalFlags.Name = "Additional Flags"
	cfg.AdditionalFlags.Description.Default = "Additional custom command line flags you want to pass Lodestar's Beacon Client, to take advantage of other settings that aren't covered here."
	cfg.AdditionalFlags.Default = ""

	return cfg
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

func (cfg *LodestarBnConfig) GetDescription() hdconfig.DynamicProperty[string] {
	return hdconfig.DynamicProperty[string]{}
}

func (cfg *LodestarBnConfig) GetDisabled() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}

func (cfg *LodestarBnConfig) GetHidden() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}
