package config

import (
	"fmt"
	"runtime"

	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
)

// Constants
const (
	// Tags
	gethTag string = "ethereum/client-go:v1.14.11"
)

// Configuration for Geth
type GethConfig struct {
	hdconfig.SectionHeader

	// Max number of P2P peers to connect to
	MaxPeers hdconfig.UintParameter

	// Number of seconds EVM calls can run before timing out
	EvmTimeout hdconfig.UintParameter

	// The archive mode flag
	ArchiveMode hdconfig.BoolParameter

	// The Docker Hub tag for Geth
	ContainerTag hdconfig.StringParameter

	// Custom command line flags
	AdditionalFlags hdconfig.StringParameter
}

type GethConfigSettings struct {
	MaxPeers        uint64 `json:"maxPeers" yaml:"maxPeers"`
	EvmTimeout      uint64 `json:"evmTimeout" yaml:"evmTimeout"`
	ArchiveMode     bool   `json:"archiveMode" yaml:"archiveMode"`
	ContainerTag    string `json:"containerTag" yaml:"containerTag"`
	AdditionalFlags string `json:"additionalFlags" yaml:"additionalFlags"`
}

// Generates a new Geth configuration
func NewGethConfig() *GethConfig {
	cfg := &GethConfig{}
	cfg.ID = hdconfig.Identifier(ids.LocalEcGethID)
	cfg.Name = "Geth"
	cfg.Description.Default = "Geth is the official Go implementation of an Ethereum client. It is known for its high performance and reliability."

	cfg.MaxPeers.ID = hdconfig.Identifier(ids.MaxPeersID)
	cfg.MaxPeers.Name = "Max Peers"
	cfg.MaxPeers.Description.Default = "The maximum number of peers Geth should connect to. This can be lowered to improve performance on low-power systems or constrained Networks. We recommend keeping it at 12 or higher."
	cfg.MaxPeers.Default = uint64(calculateGethPeers())

	cfg.EvmTimeout.ID = hdconfig.Identifier(ids.GethEvmTimeoutID)
	cfg.EvmTimeout.Name = "EVM Timeout"
	cfg.EvmTimeout.Description.Default = "The number of seconds an Execution Client API call is allowed to run before Geth times out and aborts it. Increase this if you see a lot of timeout errors in your logs."
	cfg.EvmTimeout.Default = 5

	cfg.ArchiveMode.ID = hdconfig.Identifier(ids.GethArchiveModeID)
	cfg.ArchiveMode.Name = "Enable Archive Mode"
	cfg.ArchiveMode.Description.Default = "When enabled, Geth will run in \"archive\" mode which means it can recreate the state of the chain for a previous block. This is required for manually generating the Merkle rewards tree.\n\nArchive mode takes several TB of disk space, so only enable it if you need it and can support it."
	cfg.ArchiveMode.Default = false

	cfg.ContainerTag.ID = hdconfig.Identifier(ids.ContainerTagID)
	cfg.ContainerTag.Name = "Container Tag"
	cfg.ContainerTag.Description.Default = "The tag name of the Geth container you want to use on Docker Hub."
	cfg.ContainerTag.Default = gethTag

	cfg.AdditionalFlags.ID = hdconfig.Identifier(ids.AdditionalFlagsID)
	cfg.AdditionalFlags.Name = "Additional Flags"
	cfg.AdditionalFlags.Description.Default = "Additional custom command line flags you want to pass to Geth, to take advantage of other settings that aren't covered here."
	cfg.AdditionalFlags.Default = ""

	return cfg
}

// Get the parameters for this config
func (cfg *GethConfig) GetParameters() []hdconfig.IParameter {
	return []hdconfig.IParameter{
		&cfg.MaxPeers,
		&cfg.EvmTimeout,
		&cfg.ArchiveMode,
		&cfg.ContainerTag,
		&cfg.AdditionalFlags,
	}
}

// Get the sections underneath this one
func (cfg *GethConfig) GetSections() []hdconfig.ISection {
	return []hdconfig.ISection{}
}

// Calculate the default number of Geth peers
func calculateGethPeers() uint16 {
	switch runtime.GOARCH {
	case "arm64":
		return 25
	case "amd64":
		return 50
	default:
		panic(fmt.Sprintf("unsupported architecture %s", runtime.GOARCH))
	}
}

func (cfg *GethConfig) GetDescription() hdconfig.DynamicProperty[string] {
	return hdconfig.DynamicProperty[string]{}
}

func (cfg *GethConfig) GetDisabled() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}

func (cfg *GethConfig) GetHidden() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}
