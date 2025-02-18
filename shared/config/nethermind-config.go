package config

import (
	"fmt"
	"runtime"

	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
	"github.com/pbnjay/memory"
)

// Constants
const (
	// Tags
	nethermindTag string = "nethermind/nethermind:1.29.1"
)

// Configuration for Nethermind
type NethermindConfig struct {
	// Nethermind's cache memory hint
	CacheSize hdconfig.UintParameter

	// Max number of P2P peers to connect to
	MaxPeers hdconfig.UintParameter

	// Nethermind's memory for in-memory pruning
	PruneMemSize hdconfig.UintParameter

	// Nethermind's memory budget for full pruning
	FullPruneMemoryBudget hdconfig.UintParameter

	// Nethermind's remaining disk space to trigger a pruning
	FullPruningThresholdMb hdconfig.UintParameter

	// Additional modules to enable on the primary JSON RPC endpoint
	AdditionalModules hdconfig.StringParameter

	// Additional JSON RPC URLs
	AdditionalUrls hdconfig.StringParameter

	// The Docker Hub tag for Nethermind
	ContainerTag hdconfig.StringParameter

	// Custom command line flags
	AdditionalFlags hdconfig.StringParameter
}

type NethermindConfigSettings struct {
	CacheSize              uint64 `json:"cacheSize"`
	MaxPeers               uint64 `json:"maxPeers"`
	PruneMemSize           uint64 `json:"pruneMemSize"`
	FullPruneMemoryBudget  uint64 `json:"fullPruneMemoryBudget"`
	FullPruningThresholdMb uint64 `json:"fullPruningThresholdMb"`
	AdditionalModules      string `json:"additionalModules"`
	AdditionalUrls         string `json:"additionalUrls"`
	ContainerTag           string `json:"containerTag"`
	AdditionalFlags        string `json:"additionalFlags"`
}

// Generates a new Nethermind configuration
func NewNethermindConfig() *NethermindConfig {
	cfg := &NethermindConfig{}

	cfg.CacheSize.ID = hdconfig.Identifier(ids.CacheSizeID)
	cfg.CacheSize.Name = "Cache (Memory Hint) Size"
	cfg.CacheSize.Description.Default = "The amount of RAM (in MB) you want to suggest for Nethermind's cache. While there is no guarantee that Nethermind will stay under this limit, lower values are preferred for machines with less RAM.\n\nThe default value for this will be calculated dynamically based on your system's available RAM, but you can adjust it manually."
	// cfg.CacheSize.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	cfg.MaxPeers.ID = hdconfig.Identifier(ids.MaxPeersID)
	cfg.MaxPeers.Name = "Max Peers"
	cfg.MaxPeers.Description.Default = "The maximum number of peers Nethermind should connect to. This can be lowered to improve performance on low-power systems or constrained Networks. We recommend keeping it at 12 or higher."
	// cfg.MaxPeers.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	cfg.PruneMemSize.ID = hdconfig.Identifier(ids.NethermindPruneMemSizeID)
	cfg.PruneMemSize.Name = "In-Memory Pruning Cache Size"
	cfg.PruneMemSize.Description.Default = "The amount of RAM (in MB) you want to dedicate to Nethermind for its in-memory pruning system. Higher values mean less writes to your SSD and slower overall database growth.\n\nThe default value for this will be calculated dynamically based on your system's available RAM, but you can adjust it manually."
	// cfg.PruneMemSize.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	cfg.FullPruneMemoryBudget.ID = hdconfig.Identifier(ids.NethermindFullPruneMemoryBudgetID)
	cfg.FullPruneMemoryBudget.Name = "Full Prune Memory Budget Size"
	cfg.FullPruneMemoryBudget.Description.Default = "The amount of RAM (in MB) you want to dedicate to Nethermind for its full pruning system. Higher values mean less writes to your SSD and faster pruning times.\n\nThe default value for this will be calculated dynamically based on your system's available RAM, but you can adjust it manually."
	// cfg.FullPruneMemoryBudget.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	cfg.FullPruningThresholdMb.ID = hdconfig.Identifier(ids.NethermindFullPruningThresholdMbID)
	cfg.FullPruningThresholdMb.Name = "Prune Threshold (MB)"
	cfg.FullPruningThresholdMb.Description.Default = "When the volume free space (in MB) hits this level, Nethermind will automatically start full pruning to reclaim disk space."
	// cfg.FullPruningThresholdMb.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	cfg.AdditionalModules.ID = hdconfig.Identifier(ids.NethermindAdditionalModulesID)
	cfg.AdditionalModules.Name = "Additional Modules"
	cfg.AdditionalModules.Description.Default = "Additional modules you want to add to the primary JSON-RPC route. The defaults are Eth,Net,Personal,Web3. You can add any additional ones you need here; separate multiple modules with commas, and do not use spaces."
	// cfg.AdditionalModules.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	cfg.AdditionalUrls.ID = hdconfig.Identifier(ids.NethermindAdditionalUrlsID)
	cfg.AdditionalUrls.Name = "Additional URLs"
	cfg.AdditionalUrls.Description.Default = "Additional JSON-RPC URLs you want to run alongside the primary URL. These will be added to the \"--JsonRpc.AdditionalRpcUrls\" argument. Wrap each additional URL in quotes, and separate multiple URLs with commas (no spaces). Please consult the Nethermind documentation for more information on this flag, its intended usage, and its expected formatting.\n\nFor advanced users only."
	// cfg.AdditionalUrls.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	cfg.ContainerTag.ID = hdconfig.Identifier(ids.ContainerTagID)
	cfg.ContainerTag.Name = "Container Tag"
	cfg.ContainerTag.Description.Default = "The tag name of the Nethermind container you want to use on Docker Hub."
	// cfg.ContainerTag.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	cfg.AdditionalFlags.ID = hdconfig.Identifier(ids.AdditionalFlagsID)
	cfg.AdditionalFlags.Name = "Additional Flags"
	cfg.AdditionalFlags.Description.Default = "Additional custom command line flags you want to pass to Nethermind, to take advantage of other settings that aren't covered here."
	// cfg.AdditionalFlags.AffectedContainers = []string{string(ContainerID_ExecutionClient)}

	return cfg
}

// Get the title for the config
func (cfg *NethermindConfig) GetTitle() string {
	return "Nethermind"
}

// Get the parameters for this config
func (cfg *NethermindConfig) GetParameters() []hdconfig.IParameter {
	return []hdconfig.IParameter{
		&cfg.CacheSize,
		&cfg.MaxPeers,
		&cfg.PruneMemSize,
		&cfg.FullPruneMemoryBudget,
		&cfg.FullPruningThresholdMb,
		&cfg.AdditionalModules,
		&cfg.AdditionalUrls,
		&cfg.ContainerTag,
		&cfg.AdditionalFlags,
	}
}

// Get the sections underneath this one
func (cfg *NethermindConfig) GetSections() []hdconfig.ISection {
	return []hdconfig.ISection{}
}

// Calculate the recommended size for Nethermind's cache based on the amount of system RAM
func calculateNethermindCache() uint64 {
	totalMemoryGB := memory.TotalMemory() / 1024 / 1024 / 1024

	if totalMemoryGB == 0 {
		return 0
	} else if totalMemoryGB < 9 {
		return 512
	} else if totalMemoryGB < 13 {
		return 512
	} else if totalMemoryGB < 17 {
		return 1024
	} else if totalMemoryGB < 25 {
		return 1024
	} else if totalMemoryGB < 33 {
		return 1024
	} else {
		return 2048
	}
}

// Calculate the recommended size for Nethermind's in-memory pruning based on the amount of system RAM
func calculateNethermindPruneMemSize() uint64 {
	totalMemoryGB := memory.TotalMemory() / 1024 / 1024 / 1024

	if totalMemoryGB == 0 {
		return 0
	} else if totalMemoryGB < 9 {
		return 512
	} else if totalMemoryGB < 13 {
		return 512
	} else if totalMemoryGB < 17 {
		return 1024
	} else if totalMemoryGB < 25 {
		return 1024
	} else if totalMemoryGB < 33 {
		return 1024
	} else {
		return 1024
	}
}

// Calculate the recommended size for Nethermind's full pruning based on the amount of system RAM
func calculateNethermindFullPruneMemBudget() uint64 {
	totalMemoryGB := memory.TotalMemory() / 1024 / 1024 / 1024

	if totalMemoryGB == 0 {
		return 0
	} else if totalMemoryGB < 9 {
		return 1024
	} else if totalMemoryGB < 17 {
		return 1024
	} else if totalMemoryGB < 25 {
		return 1024
	} else if totalMemoryGB < 33 {
		return 2048
	} else {
		return 4096
	}
}

// Calculate the default number of Nethermind peers
func calculateNethermindPeers() uint16 {
	switch runtime.GOARCH {
	case "arm64":
		return 25
	case "amd64":
		return 50
	default:
		panic(fmt.Sprintf("unsupported architecture %s", runtime.GOARCH))
	}
}

// TODO: Talk to Joe about these funcs required for ISection
func (cfg *NethermindConfig) GetDescription() hdconfig.DynamicProperty[string] {
	return hdconfig.DynamicProperty[string]{}
}

func (cfg *NethermindConfig) GetDisabled() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}

func (cfg *NethermindConfig) GetHidden() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}

func (cfg *NethermindConfig) GetID() hdconfig.Identifier {
	return hdconfig.Identifier("")
}

func (cfg *NethermindConfig) GetName() string {
	return ""
}
