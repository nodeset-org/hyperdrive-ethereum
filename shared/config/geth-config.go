package config

import (
	"fmt"
	"runtime"

	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	"github.com/nodeset-org/hyperdrive/modules/config"
	nmcconfig "github.com/rocket-pool/node-manager-core/config"
)

// Constants
const (
	// Tags
	gethTag string = "ethereum/client-go:v1.14.11"
)

// Configuration for Geth
type GethConfig struct {
	// Max number of P2P peers to connect to
	MaxPeers config.UintParameter

	// Number of seconds EVM calls can run before timing out
	EvmTimeout config.UintParameter

	// The archive mode flag
	ArchiveMode config.BoolParameter

	// The Docker Hub tag for Geth
	ContainerTag config.StringParameter

	// Custom command line flags
	AdditionalFlags config.StringParameter
}

// Generates a new Geth configuration
func NewGethConfig() *GethConfig {
	cfg := &GethConfig{}

	cfg.MaxPeers.ID = config.Identifier(ids.MaxPeersID)
	cfg.MaxPeers.Name = "Max Peers"
	cfg.MaxPeers.Description.Default = "The maximum number of peers Geth should connect to. This can be lowered to improve performance on low-power systems or constrained Networks. We recommend keeping it at 12 or higher."
	cfg.MaxPeers.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	cfg.EvmTimeout.ID = config.Identifier(ids.GethEvmTimeoutID)
	cfg.EvmTimeout.Name = "EVM Timeout"
	cfg.EvmTimeout.Description.Default = "The number of seconds an Execution Client API call is allowed to run before Geth times out and aborts it. Increase this if you see a lot of timeout errors in your logs."
	cfg.EvmTimeout.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	cfg.ArchiveMode.ID = config.Identifier(ids.GethArchiveModeID)
	cfg.ArchiveMode.Name = "Enable Archive Mode"
	cfg.ArchiveMode.Description.Default = "When enabled, Geth will run in \"archive\" mode which means it can recreate the state of the chain for a previous block. This is required for manually generating the Merkle rewards tree.\n\nArchive mode takes several TB of disk space, so only enable it if you need it and can support it."
	cfg.ArchiveMode.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	cfg.ContainerTag.ID = config.Identifier(ids.ContainerTagID)
	cfg.ContainerTag.Name = "Container Tag"
	cfg.ContainerTag.Description.Default = "The tag name of the Geth container you want to use on Docker Hub."
	cfg.ContainerTag.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	cfg.AdditionalFlags.ID = config.Identifier(ids.AdditionalFlagsID)
	cfg.AdditionalFlags.Name = "Additional Flags"
	cfg.AdditionalFlags.Description.Default = "Additional custom command line flags you want to pass to Geth, to take advantage of other settings that aren't covered here."
	cfg.AdditionalFlags.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	return cfg
}

// Get the title for the config
func (cfg *GethConfig) GetTitle() string {
	return "Geth"
}

// Get the parameters for this config
func (cfg *GethConfig) GetParameters() []config.IParameter {
	return []config.IParameter{
		&cfg.MaxPeers,
		&cfg.EvmTimeout,
		&cfg.ArchiveMode,
		&cfg.ContainerTag,
		&cfg.AdditionalFlags,
	}
}

// Get the sections underneath this one
func (cfg *GethConfig) GetSubconfigs() map[string]nmcconfig.IConfigSection {
	return map[string]nmcconfig.IConfigSection{}
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

// TODO (HN)
func (cfg *GethConfig) GetDescription() config.DynamicProperty[string] {
	return config.DynamicProperty[string]{}
}

// TODO (HN)
func (cfg *GethConfig) GetDisabled() config.DynamicProperty[bool] {
	return config.DynamicProperty[bool]{}
}

// TODO (HN)
func (cfg *GethConfig) GetHidden() config.DynamicProperty[bool] {
	return config.DynamicProperty[bool]{}
}

// TODO (HN)
func (cfg *GethConfig) GetID() config.Identifier {
	return config.Identifier("")
}

// TODO (HN)
func (cfg *GethConfig) GetName() string {
	return ""
}

// TODO (HN)
func (cfg *GethConfig) GetSections() []config.ISection {
	return []config.ISection{}
}
