package config

import (
	"fmt"
	"runtime"

	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
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
	MaxPeers hdconfig.UintParameter

	// The Docker Hub tag for the BN
	ContainerTag hdconfig.StringParameter

	// The pruning mode to use in the BN
	PruningMode hdconfig.ChoiceParameter[Nimbus_PruningMode]

	// Custom command line flags for the BN
	AdditionalFlags hdconfig.StringParameter
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
	cfg.ID = hdconfig.Identifier(ids.LocalBnNimbusID)
	cfg.Hidden.Default = true
	cfg.Hidden.Template = "{{if eq .GetValue \"localBeaconClient/beaconNode\" \"nimbus\"}}false{{else}}true{{end}}"

	cfg.MaxPeers.ID = hdconfig.Identifier(ids.MaxPeersID)
	cfg.MaxPeers.Name = "Max Peers"
	cfg.MaxPeers.Description.Default = "The maximum number of peers your client should try to maintain. You can try lowering this if you have a low-resource system or a constrained network."
	cfg.MaxPeers.Default = uint64(getNimbusDefaultPeers())

	options := make([]hdconfig.ParameterOption[Nimbus_PruningMode], 2)
	options[0].Name = string(Nimbus_PruningMode_Archive)
	options[0].Description.Default = "Archive mode stores all historical data, which can be useful for debugging or auditing purposes."
	options[0].Value = Nimbus_PruningMode_Archive

	options[1].Name = string(Nimbus_PruningMode_Pruned)
	options[1].Description.Default = "Prune mode stores only the most recent data, which can save disk space."
	options[1].Value = Nimbus_PruningMode_Pruned

	cfg.PruningMode.ID = hdconfig.Identifier(ids.NimbusPruningModeID)
	cfg.PruningMode.Name = "Pruning Mode"
	cfg.PruningMode.Description.Default = "Choose how Nimbus will prune its database. Highlight each option to learn more about it."
	cfg.PruningMode.Options = options
	cfg.PruningMode.Default = Nimbus_PruningMode_Pruned

	cfg.ContainerTag.ID = hdconfig.Identifier(ids.ContainerTagID)
	cfg.ContainerTag.Name = "Container Tag"
	cfg.ContainerTag.Description.Default = "The tag name of the Nimbus Beacon Node container you want to use on Docker Hub."
	cfg.ContainerTag.Default = nimbusBnTag

	cfg.AdditionalFlags.ID = hdconfig.Identifier(ids.AdditionalFlagsID)
	cfg.AdditionalFlags.Name = "Additional Flags"
	cfg.AdditionalFlags.Description.Default = "Additional custom command line flags you want to pass Nimbus's Beacon Client, to take advantage of other settings that aren't covered here."
	cfg.AdditionalFlags.Default = ""

	return cfg
}

// Get the title for the config
func (cfg *NimbusBnConfig) GetTitle() string {
	return "Nimbus Beacon Node"
}

// Get the parameters for this config
func (cfg *NimbusBnConfig) GetParameters() []hdconfig.IParameter {
	return []hdconfig.IParameter{
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
