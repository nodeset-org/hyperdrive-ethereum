package config

import (
	"path/filepath"

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
	EnableIPv6               hdconfig.BoolParameter
	ProjectName              hdconfig.StringParameter
	ApiPort                  hdconfig.UintParameter
	UserDataPath             hdconfig.StringParameter
	AutoTxMaxFee             hdconfig.FloatParameter
	MaxPriorityFee           hdconfig.FloatParameter
	AutoTxGasThreshold       hdconfig.FloatParameter
	AdditionalDockerNetworks hdconfig.StringParameter
	ClientTimeout            hdconfig.UintParameter

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

	// Logging
	Logging *sharedconfig.LoggingConfig

	// Modules
	// ModuleConfigs []*hdconfig.ModuleConfig

	// Internal fields
	Version                 string
	hyperdriveUserDirectory string
	systemPath              string
	moduleEnableStatus      map[string]bool
}

// TODO (HN) Has to match up top
type HyperdriveEthereumConfigSettings struct {
	ExampleBool   bool                    `json:"exampleBool"`
	ExampleInt    int64                   `json:"exampleInt"`
	ExampleFloat  float64                 `json:"exampleFloat"`
	ExampleString string                  `json:"exampleString"`
	ExampleChoice nativecfg.ExampleOption `json:"exampleChoice"`

	SubConfig    *SubConfigSettings    `json:"subConfig"`
	ServerConfig *ServerConfigSettings `json:"server" yaml:"server"`
}

func NewHyperdriveEthereumConfig(hdDir string, systemPath string) *HyperdriveEthereumConfig {
	cfg := &HyperdriveEthereumConfig{
		hyperdriveUserDirectory: hdDir,
		systemPath:              systemPath,
		moduleEnableStatus:      make(map[string]bool),
	}

	// Project Name
	cfg.ProjectName.ID = hdconfig.Identifier(ids.ProjectNameID)
	cfg.ProjectName.Name = "Project Name"
	cfg.ProjectName.Description.Default = "This is the prefix that will be attached to all of the Docker containers managed by Hyperdrive."
	cfg.ProjectName.Default = DefaultProjectName
	cfg.ProjectName.AffectedContainers = []string{string(sharedconfig.ContainerID_All)}

	// API Port
	cfg.ApiPort.ID = hdconfig.Identifier(ids.ApiPortID)
	cfg.ApiPort.Name = "Service API Port"
	cfg.ApiPort.Description.Default = "The port that Hyperdrive's API server should run on within the internal Docker network. Note this is bound to the local machine only; it cannot be accessed by other machines."
	cfg.ApiPort.Default = uint64(DefaultApiPort)
	cfg.ApiPort.AffectedContainers = []string{string(ContainerID_Daemon)}

	// Enable IPv6
	cfg.EnableIPv6.ID = hdconfig.Identifier(ids.EnableIPv6ID)
	cfg.EnableIPv6.Name = "Enable IPv6"
	cfg.EnableIPv6.Description.Default = "Enable IPv6 networking for Hyperdrive services. This is useful if you have an IPv6 network and want to use it for Hyperdrive.\n\nIf this isn't the first time you're starting Hyperdrive, you'll have to recreate the network after changing this box with `hyperdrive service down` and `hyperdrive service start` for it to take effect.\n\n[orange]NOTE: For IPv6 support to work, you must manually set up your Docker daemon to support it. Please follow the instructions at https://docs.docker.com/config/daemon/ipv6/#dynamic-ipv6-subnet-allocation before checking this box."
	cfg.EnableIPv6.Default = DefaultEnableIPv6
	cfg.EnableIPv6.AffectedContainers = []string{string(sharedconfig.ContainerID_All)}

	// User Data Path
	cfg.UserDataPath.ID = hdconfig.Identifier(ids.UserDataPathID)
	cfg.UserDataPath.Name = "User Data Path"
	cfg.UserDataPath.Description.Default = "The absolute path of your personal `data` folder that contains secrets such as your node wallet's encrypted file, the password for your node wallet, and all of the validator keys for any Hyperdrive modules."
	cfg.UserDataPath.Default = filepath.Join(hdDir, "data")
	cfg.UserDataPath.AffectedContainers = []string{string(ContainerID_Daemon)}

	// Additional Docker Networks
	cfg.AdditionalDockerNetworks.ID = hdconfig.Identifier(ids.AdditionalDockerNetworksID)
	cfg.AdditionalDockerNetworks.Name = "Additional Docker Networks"
	cfg.AdditionalDockerNetworks.Description.Default = "List any other externally-managed Docker networks running on this machine that you'd like to give the Hyperdrive services access to here. Use a comma-separated list of network names.\n\nTo get a list of local Docker networks, run `docker network ls`."
	cfg.AdditionalDockerNetworks.AffectedContainers = []string{string(sharedconfig.ContainerID_All)}

	// Client Timeout
	cfg.ClientTimeout.ID = hdconfig.Identifier(ids.ClientTimeoutID)
	cfg.ClientTimeout.Name = "Client Timeout"
	cfg.ClientTimeout.Description.Default = "The maximum time (in seconds) that Hyperdrive will wait for a response during HTTP requests (such as Execution Client, Beacon Node, or nodeset.io requests) before timing out."
	cfg.ClientTimeout.Default = uint64(DefaultClientTimeout)
	cfg.ClientTimeout.AffectedContainers = []string{string(ContainerID_Daemon)}

	// Container Tag
	cfg.ContainerTag.ID = hdconfig.Identifier(ids.ContainerTagID)
	cfg.ContainerTag.Name = "Service Container Tag"
	cfg.ContainerTag.Description.Default = "The tag name of the Hyperdrive Daemon image to use."
	cfg.ContainerTag.AffectedContainers = []string{string(ContainerID_Daemon)}
	cfg.ContainerTag.OverwriteOnUpgrade = true
	cfg.ContainerTag.Default = hyperdriveTag

	// AutoTxMaxFee
	cfg.AutoTxMaxFee.ID = hdconfig.Identifier(ids.AutoTxMaxFeeID)
	cfg.AutoTxMaxFee.Name = "Auto TX Max Fee"
	cfg.AutoTxMaxFee.Description.Default = "Set this if you want all of Hyperdrive's automatic transactions to use this specific max fee value (in gwei), which is the most you'd be willing to pay (*including the priority fee*).\n\nA value of 0 will use the suggested max fee based on the current network conditions.\n\nAny other value will ignore the network suggestion and use this value instead."
	cfg.AutoTxMaxFee.Default = DefaultAutoTxMaxFee
	cfg.AutoTxMaxFee.AffectedContainers = []string{string(ContainerID_Daemon)}

	// MaxPriorityFee
	cfg.MaxPriorityFee.ID = hdconfig.Identifier(ids.MaxPriorityFeeID)
	cfg.MaxPriorityFee.Name = "Max Priority Fee"
	cfg.MaxPriorityFee.Description.Default = "The default value for the priority fee (in gwei) for all of your transactions, including automatic ones. This describes how much you're willing to pay *above the network's current base fee* - the higher this is, the more ETH you give to the validators for including your transaction, which generally means it will be included in a block faster (as long as your max fee is sufficiently high to cover the current network conditions).\n\nMust be larger than 0."
	cfg.MaxPriorityFee.Default = DefaultMaxPriorityFee
	cfg.MaxPriorityFee.AffectedContainers = []string{string(ContainerID_Daemon)}

	// AutoTxGasThreshold
	cfg.AutoTxGasThreshold.ID = hdconfig.Identifier(ids.AutoTxGasThresholdID)
	cfg.AutoTxGasThreshold.Name = "Auto TX Gas Threshold"
	cfg.AutoTxGasThreshold.Description.Default = "The threshold (in gwei) that the recommended network gas price must be under in order for automated transactions to be submitted when due. A value of 0 will disable non-essential automatic transactions.\n\nNOTE: If Auto TX Max Fee is set, this setting will be ignored."
	cfg.AutoTxGasThreshold.Default = DefaultAutoTxGasThreshold
	cfg.AutoTxGasThreshold.AffectedContainers = []string{string(ContainerID_Daemon)}

	// Create the subconfigs
	cfg.Logging = sharedconfig.NewLoggingConfig()
	cfg.LocalBeaconClient = sharedconfig.NewLocalBeaconConfig()
	cfg.LocalExecutionClient = sharedconfig.NewLocalExecutionConfig()
	cfg.ExternalBeaconClient = sharedconfig.NewExternalBeaconConfig()
	cfg.ExternalExecutionClient = sharedconfig.NewExternalExecutionConfig()
	cfg.Fallback = sharedconfig.NewFallbackConfig()
	// cfg.Metrics = NewMetricsConfig()
	// cfg.MevBoost = NewMevBoostConfig(cfg)

	return cfg
}

func (cfg HyperdriveEthereumConfig) GetParameters() []hdconfig.IParameter {
	return []hdconfig.IParameter{
		&cfg.ProjectName,
		&cfg.ApiPort,
		&cfg.EnableIPv6,
		&cfg.UserDataPath,
		&cfg.AdditionalDockerNetworks,
		&cfg.ClientTimeout,
		&cfg.ContainerTag,
		&cfg.AutoTxMaxFee,
		&cfg.MaxPriorityFee,
		&cfg.AutoTxGasThreshold,
	}
}

func (cfg HyperdriveEthereumConfig) GetSections() []hdconfig.ISection {
	return []hdconfig.ISection{
		cfg.Logging,
	}
}

// func CreateInstanceFromNativeConfig(native *sharedconfig.NativeHyperdriveEthereumConfig) *ExampleConfigSettings {
// 	instance := &HyperdriveEthereumConfig{
// 		// ExampleBool:   native.ExampleBool,
// 		// ExampleInt:    native.ExampleInt,
// 		// ExampleFloat:  native.ExampleFloat,
// 		// ExampleString: native.ExampleString,
// 		// ExampleChoice: native.ExampleChoice,
// 		// SubConfig: &SubConfigSettings{
// 		// 	SubExampleBool:   native.SubConfig.SubExampleBool,
// 		// 	SubExampleChoice: native.SubConfig.SubExampleChoice,
// 		// },
// 		// ServerConfig: &ServerConfigSettings{},
// 	}
// 	return instance
// }

// func ConvertInstanceToNativeConfig(instance *ExampleConfigSettings) *sharedconfig.NativeHyperdriveEthereumConfig {
// 	native := &sharedconfig.NativeHyperdriveEthereumConfig{
// 		// ExampleBool:   instance.ExampleBool,
// 		// ExampleInt:    instance.ExampleInt,
// 		// ExampleFloat:  instance.ExampleFloat,
// 		// ExampleString: instance.ExampleString,
// 	}
// 	return native
// }
