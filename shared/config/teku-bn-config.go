package config

import (
	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
	"github.com/pbnjay/memory"
)

const (
	// Tags
	tekuBnTag string = "consensys/teku:24.10.3"
)

// Configuration for Teku
type TekuBnConfig struct {
	hdconfig.SectionHeader

	// Max number of P2P peers to connect to
	JvmHeapSize hdconfig.UintParameter

	// The max number of P2P peers to connect to
	MaxPeers hdconfig.UintParameter

	// The archive mode flag
	ArchiveMode hdconfig.BoolParameter

	// The Docker Hub tag for the Teku BN
	ContainerTag hdconfig.StringParameter

	// Custom command line flags for the BN
	AdditionalFlags hdconfig.StringParameter
}

type TekuBnConfigSettings struct {
	JvmHeapSize     uint64 `json:"jvmHeapSize" yaml:"jvmHeapSize"`
	MaxPeers        uint64 `json:"maxPeers" yaml:"maxPeers"`
	ArchiveMode     bool   `json:"archiveMode" yaml:"archiveMode"`
	ContainerTag    string `json:"containerTag" yaml:"containerTag"`
	AdditionalFlags string `json:"additionalFlags" yaml:"additionalFlags"`
}

// Generates a new Teku BN configuration
func NewTekuBnConfig() *TekuBnConfig {
	cfg := &TekuBnConfig{}
	cfg.ID = hdconfig.Identifier(ids.LocalBnTekuID)
	cfg.Hidden.Default = true
	cfg.Hidden.Template = "{{if eq .GetValue \"localBeaconClient/beaconNode\" \"teku\"}}false{{else}}true{{end}}"

	cfg.JvmHeapSize.ID = hdconfig.Identifier(ids.TekuJvmHeapSizeID)
	cfg.JvmHeapSize.Name = "JVM Heap Size"
	// cfg.JvmHeapSize.Hidden.Default = true
	cfg.JvmHeapSize.Description.Default = "The max amount of RAM, in MB, that Teku's JVM should limit itself to. Setting this lower will cause Teku to use less RAM, though it will always use more than this limit.\n\nUse 0 for automatic allocation."
	cfg.JvmHeapSize.Default = getTekuHeapSize()

	cfg.MaxPeers.ID = hdconfig.Identifier(ids.MaxPeersID)
	cfg.MaxPeers.Name = "Max Peers"
	// cfg.MaxPeers.Hidden.Default = true
	cfg.MaxPeers.Description.Default = "The maximum number of peers your client should try to maintain. You can try lowering this if you have a low-resource system or a constrained network."
	cfg.MaxPeers.Default = 100

	cfg.ArchiveMode.ID = hdconfig.Identifier(ids.TekuArchiveModeID)
	cfg.ArchiveMode.Name = "Enable Archive Mode"
	// cfg.ArchiveMode.Hidden.Default = true
	cfg.ArchiveMode.Description.Default = "When enabled, Teku will run in \"archive\" mode which means it can recreate the state of the Beacon chain for a previous block. This is required for manually generating the Merkle rewards tree.\n\nIf you are sure you will never be manually generating a tree, you can disable archive mode."
	cfg.ArchiveMode.Default = false

	cfg.ContainerTag.ID = hdconfig.Identifier(ids.ContainerTagID)
	cfg.ContainerTag.Name = "Container Tag"
	// cfg.ContainerTag.Hidden.Default = true
	cfg.ContainerTag.Description.Default = "The tag name of the Teku container on Docker Hub you want to use for the Beacon Node."
	cfg.ContainerTag.Default = tekuBnTag

	cfg.AdditionalFlags.ID = hdconfig.Identifier(ids.AdditionalFlagsID)
	cfg.AdditionalFlags.Name = "Additional Flags"
	// cfg.AdditionalFlags.Hidden.Default = true
	cfg.AdditionalFlags.Description.Default = "Additional custom command line flags you want to pass Teku's Beacon Node, to take advantage of other settings that aren't covered here."
	cfg.AdditionalFlags.Default = ""

	return cfg
}

// Get the title for the config
func (cfg *TekuBnConfig) GetTitle() string {
	return "Teku"
}

// Get the parameters for this config
func (cfg *TekuBnConfig) GetParameters() []hdconfig.IParameter {
	return []hdconfig.IParameter{
		&cfg.JvmHeapSize,
		&cfg.MaxPeers,
		&cfg.ArchiveMode,
		&cfg.ContainerTag,
		&cfg.AdditionalFlags,
	}
}

// Get the sections underneath this one
func (cfg *TekuBnConfig) GetSections() []hdconfig.ISection {
	return []hdconfig.ISection{}
}

// Get the recommended heap size for Teku
func getTekuHeapSize() uint64 {
	totalMemoryGB := memory.TotalMemory() / 1024 / 1024 / 1024
	if totalMemoryGB < 9 {
		return 2048
	}
	return 0
}

// TODO: Talk to Joe about these funcs required for ISection
func (cfg *TekuBnConfig) GetDescription() hdconfig.DynamicProperty[string] {
	return hdconfig.DynamicProperty[string]{}
}

func (cfg *TekuBnConfig) GetDisabled() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}

func (cfg *TekuBnConfig) GetHidden() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}

func (cfg *TekuBnConfig) GetID() hdconfig.Identifier {
	return hdconfig.Identifier(ids.LocalBnTekuID)
}

func (cfg *TekuBnConfig) GetName() string {
	return "Teku"
}
