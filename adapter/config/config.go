package config

import (
	"fmt"

	"github.com/nodeset-org/hyperdrive-ethereum/shared"
	sharedconfig "github.com/nodeset-org/hyperdrive-ethereum/shared/config"
	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
)

type PortMode string

const (
	PortMode_Closed    PortMode = "closed"
	PortMode_Localhost PortMode = "localhost"
	PortMode_External  PortMode = "external"
)

const (
	// Tags
	hyperdriveTag string = "nodeset/hyperdrive-ethereum:v" + shared.HyperdriveEthereumVersion

	// Private parameter names
	versionName          string = "version"
	moduleEnabledMapName string = "modules"

	// Defaults
	DefaultProjectName   string = "hyperdrive-ethereum"
	DefaultApiPort       uint16 = 8080
	DefaultEnableIPv6    bool   = false
	DefaultClientTimeout uint16 = 30

	// TODO: Discuss these values
	DefaultAutoTxMaxFee       float64 = 1000.0
	DefaultMaxPriorityFee     float64 = 1000.0
	DefaultAutoTxGasThreshold float64 = 100000
)

type HyperdriveEthereumConfig struct {
	ProjectName        hdconfig.StringParameter
	ApiPort            hdconfig.UintParameter
	AutoTxMaxFee       hdconfig.FloatParameter
	MaxPriorityFee     hdconfig.FloatParameter
	AutoTxGasThreshold hdconfig.FloatParameter
	EnableIPv6         hdconfig.BoolParameter

	Network    hdconfig.ChoiceParameter[sharedconfig.Network] // hdconfig.Parameter[config.Network]
	ClientMode hdconfig.ChoiceParameter[ClientMode]           // hdconfig.Parameter[config.ClientMode]

	// Execution client settings
	LocalExecutionClient    *sharedconfig.LocalExecutionConfig
	ExternalExecutionClient *sharedconfig.ExternalExecutionConfig

	// Beacon node settings
	LocalBeaconClient    *sharedconfig.LocalBeaconConfig
	ExternalBeaconClient *sharedconfig.ExternalBeaconConfig
	// Fallback clients

	Fallback *sharedconfig.FallbackConfig
	// Metrics
	// Metrics *config.MetricsConfig
	// MEV-Boost
	// MevBoost *MevBoostConfig

	// The Docker Hub tag for the daemon container
	ContainerTag hdconfig.StringParameter

	// Internal fields
	Version hdconfig.StringParameter

	ServerConfig *ServerConfig
	// DockerConfig *Docker

	// Server settings
	IsNew      hdconfig.BoolParameter
	ExternalIp hdconfig.StringParameter
}

type HyperdriveEthereumConfigSettings struct {
	ProjectName        string  `json:"projectName" yaml:"projectName"`
	ApiPort            uint    `json:"apiPort" yaml:"apiPort"`
	AutoTxMaxFee       float64 `json:"autoTxMaxFee" yaml:"autoTxMaxFee"`
	MaxPriorityFee     float64 `json:"maxPriorityFee" yaml:"maxPriorityFee"`
	AutoTxGasThreshold float64 `json:"autoTxGasThreshold" yaml:"autoTxGasThreshold"`
	EnableIPv6         bool    `json:"enableIPv6" yaml:"enableIPv6"`

	Network    sharedconfig.Network `json:"network" yaml:"network"`
	ClientMode ClientMode           `json:"clientMode" yaml:"clientMode"`

	LocalExecutionClient    *sharedconfig.LocalExecutionConfigSettings    `json:"localExecutionClient" yaml:"localExecutionClient"`
	ExternalExecutionClient *sharedconfig.ExternalExecutionConfigSettings `json:"externalExecutionClient" yaml:"externalExecutionClient"`

	LocalBeaconClient    *sharedconfig.LocalBeaconConfigSettings    `json:"localBeaconClient" yaml:"localBeaconClient"`
	ExternalBeaconClient *sharedconfig.ExternalBeaconConfigSettings `json:"externalBeaconClient" yaml:"externalBeaconClient"`

	Fallback *sharedconfig.FallbackConfigSettings `json:"fallback" yaml:"fallback"`

	ContainerTag string `json:"containerTag" yaml:"containerTag"`

	Version string `json:"version" yaml:"version"`

	ServerConfig *ServerConfigSettings `json:"server" yaml:"server"`
	// DockerConfig *DockerSettings       `json:"dockerConfig"`

	IsNew      bool   `json:"isNew" yaml:"isNew"`
	ExternalIp string `json:"externalIp" yaml:"externalIp"`
}

func NewHyperdriveEthereumConfig() *HyperdriveEthereumConfig {
	cfg := &HyperdriveEthereumConfig{}

	// API Port
	cfg.ApiPort.ID = hdconfig.Identifier(ids.ApiPortID)
	cfg.ApiPort.Name = "Service API Port"
	cfg.ApiPort.Description.Default = "The port that Hyperdrive's API server should run on within the internal Docker network. Note this is bound to the local machine only; it cannot be accessed by other machines."
	cfg.ApiPort.Default = uint64(DefaultApiPort)

	// Container Tag
	cfg.ContainerTag.ID = hdconfig.Identifier(ids.ContainerTagID)
	cfg.ContainerTag.Name = "Service Container Tag"
	cfg.ContainerTag.Description.Default = "The tag name of the Hyperdrive Daemon image to use."
	cfg.ContainerTag.OverwriteOnUpgrade = true
	cfg.ContainerTag.Default = hyperdriveTag

	// AutoTxMaxFee
	cfg.AutoTxMaxFee.ID = hdconfig.Identifier(ids.AutoTxMaxFeeID)
	cfg.AutoTxMaxFee.Name = "Auto TX Max Fee"
	cfg.AutoTxMaxFee.Description.Default = "Set this if you want all of Hyperdrive's automatic transactions to use this specific max fee value (in gwei), which is the most you'd be willing to pay (*including the priority fee*).\n\nA value of 0 will use the suggested max fee based on the current network conditions.\n\nAny other value will ignore the network suggestion and use this value instead."
	cfg.AutoTxMaxFee.Default = DefaultAutoTxMaxFee

	// MaxPriorityFee
	cfg.MaxPriorityFee.ID = hdconfig.Identifier(ids.MaxPriorityFeeID)
	cfg.MaxPriorityFee.Name = "Max Priority Fee"
	cfg.MaxPriorityFee.Description.Default = "The default value for the priority fee (in gwei) for all of your transactions, including automatic ones. This describes how much you're willing to pay *above the network's current base fee* - the higher this is, the more ETH you give to the validators for including your transaction, which generally means it will be included in a block faster (as long as your max fee is sufficiently high to cover the current network conditions).\n\nMust be larger than 0."
	cfg.MaxPriorityFee.Default = DefaultMaxPriorityFee

	// AutoTxGasThreshold
	cfg.AutoTxGasThreshold.ID = hdconfig.Identifier(ids.AutoTxGasThresholdID)
	cfg.AutoTxGasThreshold.Name = "Auto TX Gas Threshold"
	cfg.AutoTxGasThreshold.Description.Default = "The threshold (in gwei) that the recommended network gas price must be under in order for automated transactions to be submitted when due. A value of 0 will disable non-essential automatic transactions.\n\nNOTE: If Auto TX Max Fee is set, this setting will be ignored."
	cfg.AutoTxGasThreshold.Default = DefaultAutoTxGasThreshold

	// Create the subconfigs
	cfg.LocalBeaconClient = sharedconfig.NewLocalBeaconConfig()
	cfg.LocalExecutionClient = sharedconfig.NewLocalExecutionConfig()
	cfg.ExternalBeaconClient = sharedconfig.NewExternalBeaconConfig()
	cfg.ExternalExecutionClient = sharedconfig.NewExternalExecutionConfig()
	cfg.Fallback = sharedconfig.NewFallbackConfig()
	// cfg.Metrics = NewMetricsConfig()
	// cfg.MevBoost = NewMevBoostConfig(cfg)

	cfg.ServerConfig = NewServerConfig()

	return cfg
}

func (cfg HyperdriveEthereumConfig) GetParameters() []hdconfig.IParameter {
	return []hdconfig.IParameter{
		&cfg.ApiPort,
		&cfg.ContainerTag,
		&cfg.AutoTxMaxFee,
		&cfg.MaxPriorityFee,
		&cfg.AutoTxGasThreshold,
	}
}

func (cfg HyperdriveEthereumConfig) GetSections() []hdconfig.ISection {
	return []hdconfig.ISection{
		cfg.ServerConfig,
	}
}

func CreateInstanceFromNativeConfig(native *sharedconfig.NativeHyperdriveEthereumSettings) *HyperdriveEthereumConfigSettings {
	instance := &HyperdriveEthereumConfigSettings{
		ApiPort:            native.ApiPort,
		AutoTxMaxFee:       native.AutoTxMaxFee,
		MaxPriorityFee:     native.MaxPriorityFee,
		AutoTxGasThreshold: native.AutoTxGasThreshold,
		Network:            native.Network,
		// ClientMode:               native.ClientMode,
		ContainerTag: native.ContainerTag,
		ServerConfig: &ServerConfigSettings{},
	}
	return instance
}

func ConvertInstanceToNativeConfig(instance *HyperdriveEthereumConfigSettings) *sharedconfig.NativeHyperdriveEthereumSettings {
	native := &sharedconfig.NativeHyperdriveEthereumSettings{
		ApiPort:            instance.ApiPort,
		AutoTxMaxFee:       instance.AutoTxMaxFee,
		MaxPriorityFee:     instance.MaxPriorityFee,
		AutoTxGasThreshold: instance.AutoTxGasThreshold,
		Network:            instance.Network,
		// ClientMode:               instance.ClientMode,
		ContainerTag: instance.ContainerTag,
	}
	return native
}

// GetChangedServices returns a list of services that would be affected by the new settings
func (s *HyperdriveEthereumConfigSettings) GetChangedServices(oldSettings *HyperdriveEthereumConfigSettings) ([]string, error) {
	cfg := NewHyperdriveEthereumConfig()
	newModSettings := hdconfig.CreateModuleSettings(cfg)
	err := newModSettings.CopySettingsFromKnownType(s)
	if err != nil {
		return nil, fmt.Errorf("error copying new settings: %w", err)
	}

	oldModSettings := hdconfig.CreateModuleSettings(cfg)
	err = oldModSettings.CopySettingsFromKnownType(oldSettings)
	if err != nil {
		return nil, fmt.Errorf("error copying old settings: %w", err)
	}

	// Compare the settings - if there are no differences, return nil
	diff := hdconfig.CompareSettings(cfg, oldModSettings, newModSettings)
	if len(diff.ParameterDifferences) == 0 && len(diff.SectionDifferences) == 0 {
		return nil, nil
	}

	// Any parameter changes will affect the service, so just return it
	changedServices := []string{
		shared.ServiceContainerName,
	}
	return changedServices, nil
}

// func (c *HyperdriveEthereumConfigSettings) GetDockerArtifactName(moduleName string) string {
// 	switch moduleName {
// 	case "beacon-node":
// 		return c.DockerConfig.BeaconNodeContainer
// 	case "execution-client":
// 		return c.DockerConfig.ExecutionClientContainer
// 	default:
// 		return ""
// 	}
// }

func (c *HyperdriveEthereumConfigSettings) GetAllModuleConfigs() []any {
	return []any{
		c.LocalBeaconClient,
		c.LocalExecutionClient,
	}
}
