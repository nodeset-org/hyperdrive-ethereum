package config

import (
	"fmt"
	"runtime"

	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	"github.com/nodeset-org/hyperdrive/modules/config"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
)

const (
	// Tags
	nimbusBnTag string = "statusim/nimbus-eth2:multiarch-v24.10.0"
)

// Nimbus's pruning mode
type Nimbus_PruningMode string

const (
	Nimbus_PruningMode_Archive Nimbus_PruningMode = "archive"
	Nimbus_PruningMode_Pruned  Nimbus_PruningMode = "prune"
)

// Configuration for Nimbus
type NimbusBnConfig struct {
	hdconfig.SectionHeader
	// The max number of P2P peers to connect to
	MaxPeers config.UintParameter

	// The Docker Hub tag for the BN
	ContainerTag config.StringParameter

	// The pruning mode to use in the BN
	PruningMode config.StringParameter //Parameter[Nimbus_PruningMode]

	// Custom command line flags for the BN
	AdditionalFlags config.StringParameter
}

type NimbusBnConfigSettings struct {
	MaxPeers        uint64 `json:"maxPeers" yaml:"maxPeers"`
	ContainerTag    string `json:"containerTag" yaml:"containerTag"`
	PruningMode     string `json:"pruningMode" yaml:"pruningMode"`
	AdditionalFlags string `json:"additionalFlags" yaml:"additionalFlags"`
}

// Generates a new Nimbus configuration
func NewNimbusBnConfig() *NimbusBnConfig {
	cfg := &NimbusBnConfig{}
	cfg.ID = config.Identifier(ids.LocalBnNimbusID)

	cfg.MaxPeers.ID = config.Identifier(ids.MaxPeersID)
	cfg.MaxPeers.Name = "Max Peers"
	cfg.MaxPeers.Description.Default = "The maximum number of peers your client should try to maintain. You can try lowering this if you have a low-resource system or a constrained network."
	// cfg.MaxPeers.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	cfg.PruningMode.ID = config.Identifier(ids.NimbusPruningModeID)
	cfg.PruningMode.Name = "Pruning Mode"
	cfg.PruningMode.Description.Default = "Choose how Nimbus will prune its database. Highlight each option to learn more about it."
	// cfg.PruningMode.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	cfg.PruningMode.ID = config.Identifier(ids.NimbusPruningModeID)
	cfg.PruningMode.Name = "Pruning Mode"
	cfg.PruningMode.Description.Default = "Choose how Nimbus will prune its database. Highlight each option to learn more about it."
	// cfg.PruningMode.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	cfg.ContainerTag.ID = config.Identifier(ids.ContainerTagID)
	cfg.ContainerTag.Name = "Container Tag"
	cfg.ContainerTag.Description.Default = "The tag name of the Nimbus Beacon Node container you want to use on Docker Hub."
	// cfg.ContainerTag.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	cfg.AdditionalFlags.ID = config.Identifier(ids.AdditionalFlagsID)
	cfg.AdditionalFlags.Name = "Additional Flags"
	cfg.AdditionalFlags.Description.Default = "Additional custom command line flags you want to pass Nimbus's Beacon Client, to take advantage of other settings that aren't covered here."
	// cfg.AdditionalFlags.AffectedContainers = []string{string(ContainerID_BeaconNode)}

	return cfg
}

// Get the title for the config
func (cfg *NimbusBnConfig) GetTitle() string {
	return "Nimbus Beacon Node"
}

// Get the parameters for this config
func (cfg *NimbusBnConfig) GetParameters() []config.IParameter {
	return []config.IParameter{
		&cfg.MaxPeers,
		&cfg.ContainerTag,
		&cfg.PruningMode,
		&cfg.AdditionalFlags,
	}
}

// Get the sections underneath this one
func (cfg *NimbusBnConfig) GetSections() []hdconfig.ISection {
	return []hdconfig.ISection{}
}

// Get the default number of peers
func getNimbusDefaultPeers() uint16 {
	switch runtime.GOARCH {
	case "arm64":
		return 100
	case "amd64":
		return 160
	default:
		panic(fmt.Sprintf("unsupported architecture %s", runtime.GOARCH))
	}
}

// TODO: Talk to Joe about these funcs required for ISection
func (cfg *NimbusBnConfig) GetDescription() hdconfig.DynamicProperty[string] {
	return hdconfig.DynamicProperty[string]{}
}

func (cfg *NimbusBnConfig) GetDisabled() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}

func (cfg *NimbusBnConfig) GetHidden() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}

func (cfg *NimbusBnConfig) GetID() hdconfig.Identifier {
	return hdconfig.Identifier(ids.LocalBnNimbusID)
}

func (cfg *NimbusBnConfig) GetName() string {
	return "Nimbus"
}
