package config

import (
	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
)

const (
	// Tags
	lighthouseBnTag string = "sigp/lighthouse:v5.3.0"

	lighthouseDefaultMaxPeers    uint = 100
	lighthouseDefaultP2pQuicPort uint = 8001
)

// Configuration for the Lighthouse BN
type LighthouseBnConfig struct {
	hdconfig.SectionHeader

	// The port to use for gossip traffic using the QUIC protocol
	P2pQuicPort hdconfig.UintParameter

	// The max number of P2P peers to connect to
	MaxPeers hdconfig.UintParameter

	// The Docker Hub tag for Lighthouse BN
	ContainerTag hdconfig.StringParameter

	// Custom command line flags for the BN
	AdditionalFlags hdconfig.StringParameter
}

type LighthouseBnConfigSettings struct {
	P2pQuicPort     uint64 `json:"p2pQuicPort" yaml:"p2pQuicPort"`
	MaxPeers        uint64 `json:"maxPeers" yaml:"maxPeers"`
	ContainerTag    string `json:"containerTag" yaml:"containerTag"`
	AdditionalFlags string `json:"additionalFlags" yaml:"additionalFlags"`
}

// Generates a new Lighthouse BN configuration
func NewLighthouseBnConfig() *LighthouseBnConfig {
	cfg := &LighthouseBnConfig{}
	cfg.ID = hdconfig.Identifier(ids.LocalBnLighthouseID)
	cfg.Name = "Lighthouse"
	cfg.Description.Default = "Lighthouse"

	// cfg.Hidden.Default = true
	// cfg.Hidden.Template = "{{if eq .GetValue \"localBeaconClient/beaconNode\" \"lighthouse\"}}false{{else}}true{{end}}"

	cfg.MaxPeers.ID = hdconfig.Identifier(ids.MaxPeersID)
	cfg.MaxPeers.Name = "Max Peers"
	cfg.MaxPeers.Description.Default = "The maximum number of peers your client should try to maintain. You can try lowering this if you have a low-resource system or a constrained network."
	cfg.MaxPeers.Default = uint64(lighthouseDefaultMaxPeers)

	cfg.ContainerTag.ID = hdconfig.Identifier(ids.ContainerTagID)
	cfg.ContainerTag.Name = "Container Tag"
	cfg.ContainerTag.Description.Default = "The tag name of the Lighthouse container from Docker Hub you want to use for the Beacon Node."
	cfg.ContainerTag.Default = lighthouseBnTag

	cfg.AdditionalFlags.ID = hdconfig.Identifier(ids.AdditionalFlagsID)
	cfg.AdditionalFlags.Name = "Additional Flags"
	cfg.AdditionalFlags.Description.Default = "Additional custom command line flags you want to pass Lighthouse's Beacon Node, to take advantage of other settings that aren't covered here."
	cfg.AdditionalFlags.Default = ""

	cfg.P2pQuicPort.ID = hdconfig.Identifier(ids.P2pQuicPortID)
	cfg.P2pQuicPort.Name = "P2pQuicPort"
	cfg.P2pQuicPort.Description.Default = "The port to use for P2P (blockchain) traffic using the QUIC protocol."
	cfg.P2pQuicPort.Default = uint64(lighthouseDefaultP2pQuicPort)
	return cfg
}

// Get the parameters for this config
func (cfg *LighthouseBnConfig) GetParameters() []hdconfig.IParameter {
	return []hdconfig.IParameter{
		&cfg.MaxPeers,
		&cfg.P2pQuicPort,
		&cfg.ContainerTag,
		&cfg.AdditionalFlags,
	}
}

// Get the sections underneath this one
func (cfg *LighthouseBnConfig) GetSubconfigs() map[string]hdconfig.ISection {
	return map[string]hdconfig.ISection{}
}

func (cfg *LighthouseBnConfig) GetSections() []hdconfig.ISection {
	return []hdconfig.ISection{}
}

func (cfg *LighthouseBnConfig) GetDescription() hdconfig.DynamicProperty[string] {
	return hdconfig.DynamicProperty[string]{}
}

func (cfg *LighthouseBnConfig) GetDisabled() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}

func (cfg *LighthouseBnConfig) GetHidden() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}
