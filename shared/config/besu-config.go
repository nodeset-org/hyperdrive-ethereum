package config

import (
	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
)

// Constants
const (
	// Tags
	besuTag string = "hyperledger/besu:24.10.0"
)

// Configuration for Besu
type BesuConfig struct {
	hdconfig.SectionHeader

	// Max number of P2P peers to connect to
	JvmHeapSize hdconfig.UintParameter

	// Max number of P2P peers to connect to
	MaxPeers hdconfig.UintParameter

	// Historical state block regeneration limit
	MaxBackLayers hdconfig.UintParameter

	// The archive mode flag
	ArchiveMode hdconfig.BoolParameter

	// The Docker Hub tag for Besu
	ContainerTag hdconfig.StringParameter

	// Custom command line flags
	AdditionalFlags hdconfig.StringParameter
}

type BesuConfigSettings struct {
	JvmHeapSize     uint64 `json:"jvmHeapSize" yaml:"jvmHeapSize"`
	MaxPeers        uint64 `json:"maxPeers" yaml:"maxPeers"`
	MaxBackLayers   uint64 `json:"maxBackLayers" yaml:"maxBackLayers"`
	ArchiveMode     bool   `json:"archiveMode" yaml:"archiveMode"`
	ContainerTag    string `json:"containerTag" yaml:"containerTag"`
	AdditionalFlags string `json:"additionalFlags" yaml:"additionalFlags"`
}

// Generates a new Besu configuration
func NewBesuConfig() *BesuConfig {
	cfg := &BesuConfig{}
	cfg.ID = hdconfig.Identifier(ids.LocalEcBesuID)

	// TODO: Get these reviewed
	cfg.JvmHeapSize.ID = hdconfig.Identifier(ids.BesuJvmHeapSizeID)
	cfg.JvmHeapSize.Name = "JVM Heap Size"
	cfg.JvmHeapSize.Description.Default = "The max amount of RAM, in MB, that Besu's JVM should limit itself to. Setting this lower will cause Besu to use less RAM, though it will always use more than this limit.\n\nUse 0 for automatic allocation."
	cfg.JvmHeapSize.Default = 0

	cfg.MaxPeers.ID = hdconfig.Identifier(ids.MaxPeersID)
	cfg.MaxPeers.Name = "Max Peers"
	cfg.MaxPeers.Description.Default = "The maximum number of peers Besu should connect to. This can be lowered to improve performance on low-power systems or constrained networks. We recommend keeping it at 12 or higher."
	cfg.MaxPeers.Default = 25

	cfg.MaxBackLayers.ID = hdconfig.Identifier(ids.BesuMaxBackLayersID)
	cfg.MaxBackLayers.Name = "Historical Block Replay Limit"
	cfg.MaxBackLayers.Description.Default = "Besu has the ability to revisit the state of any historical block on the chain by \"replaying\" all of the previous blocks to get back to the target. This limit controls how many blocks you can replay - in other words, how far back Besu can go in time. Normal Execution client processing will be paused while a replay is in progress.\n\n[orange]NOTE: If you try to replay a state from a long time ago, it may take Besu several minutes to rebuild the state!"
	cfg.MaxBackLayers.Default = 512

	cfg.ArchiveMode.ID = hdconfig.Identifier(ids.BesuArchiveModeID)
	cfg.ArchiveMode.Name = "Enable Archive Mode"
	cfg.ArchiveMode.Description.Default = "When enabled, Besu will run in \"archive\" mode which means it can recreate the state of the chain for a previous block. This is required for accessing the state of blocks that are more than about half-an-hour old, which may be a part of things like reward systems."
	cfg.ArchiveMode.Default = false

	cfg.ContainerTag.ID = hdconfig.Identifier(ids.ContainerTagID)
	cfg.ContainerTag.Name = "Container Tag"
	cfg.ContainerTag.Description.Default = "The tag name of the Besu container you want to use on Docker Hub."
	cfg.ContainerTag.Default = besuTag

	cfg.AdditionalFlags.ID = hdconfig.Identifier(ids.AdditionalFlagsID)
	cfg.AdditionalFlags.Name = "Additional Flags"
	cfg.AdditionalFlags.Description.Default = "Additional custom command line flags you want to pass to Besu, to take advantage of other settings that aren't covered here."
	cfg.AdditionalFlags.Default = ""

	return cfg
}

// The title for the config
func (cfg *BesuConfig) GetTitle() string {
	return "Besu"
}

// Get the parameters for this config
func (cfg *BesuConfig) GetParameters() []hdconfig.IParameter {
	return []hdconfig.IParameter{
		&cfg.JvmHeapSize,
		&cfg.MaxPeers,
		&cfg.MaxBackLayers,
		&cfg.ArchiveMode,
		&cfg.ContainerTag,
		&cfg.AdditionalFlags,
	}
}

// Get the sections underneath this one
func (cfg *BesuConfig) GetSections() []hdconfig.ISection {
	return []hdconfig.ISection{}
}

// TODO: Talk to Joe about these funcs required for ISection
func (cfg *BesuConfig) GetDescription() hdconfig.DynamicProperty[string] {
	return hdconfig.DynamicProperty[string]{}
}

func (cfg *BesuConfig) GetDisabled() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}

func (cfg *BesuConfig) GetHidden() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}

func (cfg *BesuConfig) GetID() hdconfig.Identifier {
	return hdconfig.Identifier(ids.LocalEcBesuID)
}

func (cfg *BesuConfig) GetName() string {
	return "Besu"
}
