package config

import (
	"runtime"

	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	"github.com/nodeset-org/hyperdrive/modules/config"
	"github.com/pbnjay/memory"
	nmcconfig "github.com/rocket-pool/node-manager-core/config"
)

// Constants
const (
	rethTag string = "ghcr.io/paradigmxyz/reth:v1.1.0"
)

// Configuration for Reth
type RethConfig struct {
	// Size of Reth's Cache
	CacheSize config.UintParameter

	// Max number of P2P peers that can connect to this node
	MaxInboundPeers config.UintParameter

	// Max number of P2P peers to this node can connect to
	MaxOutboundPeers config.UintParameter

	// The Docker Hub tag for Reth
	ContainerTag config.StringParameter

	// Custom command line flags
	AdditionalFlags config.StringParameter
}

// Generates a new Reth configuration
func NewRethConfig() *RethConfig {
	cfg := &RethConfig{}

	cfg.CacheSize.ID = config.Identifier(ids.CacheSizeID)
	cfg.CacheSize.Name = "Cache Size"
	cfg.CacheSize.Description.Default = "The amount of RAM (in MB) you want Reth's cache to use. Larger values mean your disk space usage will increase slower, and you will have to prune less frequently. The default is based on how much total RAM your system has but you can adjust it manually."
	cfg.CacheSize.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	cfg.MaxInboundPeers.ID = config.Identifier(ids.RethMaxInboundPeersID)
	cfg.MaxInboundPeers.Name = "Max Inbound Peers"
	cfg.MaxInboundPeers.Description.Default = "The maximum number of inbound peers that should be allowed to connect to Reth (peers that request to connect to your node). This can be lowered to improve performance on low-power systems or constrained networks. Inbound peers requires you to have properly forwarded ports. We recommend keeping the sum of this and max outbound peers at 12 or higher."
	cfg.MaxInboundPeers.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	cfg.MaxOutboundPeers.ID = config.Identifier(ids.RethMaxOutboundPeersID)
	cfg.MaxOutboundPeers.Name = "Max Outbound Peers"
	cfg.MaxOutboundPeers.Description.Default = "The maximum number of outbound peers that Reth can connect to (peers that your node requests to connect to). This can be lowered to improve performance on low-power systems or constrained networks. Outbound peers do not require proper port forwarding, but are slower to accumulate than inbound peers. We recommend keeping the sum of this and max outbound peers at 12 or higher."
	cfg.MaxOutboundPeers.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	cfg.ContainerTag.ID = config.Identifier(ids.ContainerTagID)
	cfg.ContainerTag.Name = "Container Tag"
	cfg.ContainerTag.Description.Default = "The tag name of the Reth container you want to use."
	cfg.ContainerTag.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	cfg.AdditionalFlags.ID = config.Identifier(ids.AdditionalFlagsID)
	cfg.AdditionalFlags.Name = "Additional Flags"
	cfg.AdditionalFlags.Description.Default = "Additional custom command line flags you want to pass to Reth, to take advantage of other settings that aren't covered here."
	cfg.AdditionalFlags.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	return cfg
}

// Get the title for the config
func (cfg *RethConfig) GetTitle() string {
	return "Reth"
}

// Get the config.Parameters for this config
func (cfg *RethConfig) GetParameters() []config.IParameter {
	return []config.IParameter{
		&cfg.CacheSize,
		&cfg.MaxInboundPeers,
		&cfg.MaxOutboundPeers,
		&cfg.ContainerTag,
		&cfg.AdditionalFlags,
	}
}

// Get the sections underneath this one
func (cfg *RethConfig) GetSubconfigs() map[string]nmcconfig.IConfigSection {
	return map[string]nmcconfig.IConfigSection{}
}

// Calculate the recommended size for Reth's cache based on the amount of system RAM
func calculateRethCache() uint64 {
	totalMemoryGB := memory.TotalMemory() / 1024 / 1024 / 1024

	if totalMemoryGB == 0 {
		return 0
	} else if totalMemoryGB < 9 {
		return 256
	} else if totalMemoryGB < 13 {
		return 2048
	} else if totalMemoryGB < 17 {
		return 4096
	} else if totalMemoryGB < 25 {
		return 8192
	} else if totalMemoryGB < 33 {
		return 12288
	} else {
		return 16384
	}
}

// Calculate the default number of Reth peers
func calculateRethPeers() uint16 {
	if runtime.GOARCH == "arm64" {
		return 12
	}
	return 25
}

// TODO (HN)
func (cfg *RethConfig) GetDescription() config.DynamicProperty[string] {
	return config.DynamicProperty[string]{}
}

// TODO (HN)
func (cfg *RethConfig) GetDisabled() config.DynamicProperty[bool] {
	return config.DynamicProperty[bool]{}
}

// TODO (HN)
func (cfg *RethConfig) GetHidden() config.DynamicProperty[bool] {
	return config.DynamicProperty[bool]{}
}

// TODO (HN)
func (cfg *RethConfig) GetID() config.Identifier {
	return config.Identifier("")
}

// TODO (HN)
func (cfg *RethConfig) GetName() string {
	return ""
}

// TODO (HN)
func (cfg *RethConfig) GetSections() []config.ISection {
	return []config.ISection{}
}
