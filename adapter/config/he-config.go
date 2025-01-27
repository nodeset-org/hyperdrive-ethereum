package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
	"gopkg.in/yaml.v2"

	sharedconfig "github.com/nodeset-org/hyperdrive-ethereum/shared/config"
	"github.com/urfave/cli/v2"
)

type PortMode string

const (
	ConfigFileMode os.FileMode = 0644

	PortMode_Closed    PortMode = "closed"
	PortMode_Localhost PortMode = "localhost"
	PortMode_External  PortMode = "external"
)

type ServerConfig struct {
	hdconfig.SectionMetadataHeader

	Port hdconfig.UintParameterMetadata

	PortMode hdconfig.ChoiceParameterMetadata[PortMode]
}

// Configuration manager
type AdapterConfigManager struct {
	// The adapter configuration
	AdapterConfig *sharedconfig.HyperdriveEthereumConfig

	// The native configuration manager
	nativeConfigManager *sharedconfig.ConfigManager

	// The path to the adapter configuration file
	adapterConfigPath string
}

// Create a new configuration manager for the adapter
func NewAdapterConfigManager(c *cli.Context) (*AdapterConfigManager, error) {
	configDir := c.String(utils.ConfigDirFlag.Name)
	if configDir == "" {
		return nil, fmt.Errorf("config directory is required")
	}
	return &AdapterConfigManager{
		nativeConfigManager: sharedconfig.NewConfigManager(filepath.Join(configDir, utils.ServiceConfigFile)),
		adapterConfigPath:   filepath.Join(configDir, utils.AdapterConfigFile),
	}, nil
}

func NewHyperdriveEthereumConfig() *sharedconfig.HyperdriveEthereumConfig {
	cfg := &sharedconfig.HyperdriveEthereumConfig{}

	// TODO: HN
	// // ExampleBool
	// cfg.ExampleBool.ID = hdconfig.Identifier(ids.ExampleBoolID)
	// cfg.ExampleBool.Name = "Example Boolean"
	// cfg.ExampleBool.Description.Default = "This is an example of a boolean parameter. It doesn't directly affect the service, but it does control the behavior of some other config parameters."
	// cfg.ExampleBool.AffectedContainers = []string{shared.ServiceContainerName}
	// cfg.ExampleBool.Value = cfg.ExampleBool.Default

	// // ExampleInt
	// cfg.ExampleInt.ID = hdconfig.Identifier(ids.ExampleIntID)
	// cfg.ExampleInt.Name = "Example Integer"
	// cfg.ExampleInt.Description.Default = "This is an example of an integer parameter."
	// cfg.ExampleInt.AffectedContainers = []string{shared.ServiceContainerName}
	// cfg.ExampleInt.Value = cfg.ExampleInt.Default

	// // ExampleUint
	// cfg.ExampleUint.ID = hdconfig.Identifier(ids.ExampleUintID)
	// cfg.ExampleUint.Name = "Example Unsigned Integer"
	// cfg.ExampleUint.Description.Default = "This is an example of an unsigned integer parameter."
	// cfg.ExampleUint.AffectedContainers = []string{shared.ServiceContainerName}
	// cfg.ExampleUint.Value = cfg.ExampleUint.Default

	// // ExampleFloat
	// cfg.ExampleFloat.ID = hdconfig.Identifier(ids.ExampleFloatID)
	// cfg.ExampleFloat.Name = "Example Float"
	// cfg.ExampleFloat.Description.Default = "This is an example of a float parameter with a minimum and maximum set."
	// cfg.ExampleFloat.Default = 50
	// cfg.ExampleFloat.MinValue = 0.0
	// cfg.ExampleFloat.MaxValue = 100.0
	// cfg.ExampleFloat.Value = cfg.ExampleFloat.Default
	// cfg.ExampleFloat.AffectedContainers = []string{shared.ServiceContainerName}

	// // ExampleString
	// cfg.ExampleString.ID = hdconfig.Identifier(ids.ExampleStringID)
	// cfg.ExampleString.Name = "Example String"
	// cfg.ExampleString.Description.Default = "This is an example of a string parameter. It has a max length and regex pattern set."
	// cfg.ExampleString.MaxLength = 10
	// cfg.ExampleString.Regex = "^[a-zA-Z]*$"
	// cfg.ExampleString.Value = cfg.ExampleString.Default
	// cfg.ExampleString.AffectedContainers = []string{shared.ServiceContainerName}

	// // Options for ExampleChoice
	// options := make([]hdconfig.ParameterMetadataOption[nativecfg.ExampleOption], 3)
	// options[0].Name = "One"
	// options[0].Description.Default = "This is the first option."
	// options[0].Value = nativecfg.ExampleOption_One

	// thresholdString := strconv.FormatFloat(FloatThreshold, 'f', -1, 64)
	// options[1].Name = "Two"
	// options[1].Description.Default = "This is the second option. It is hidden when ExampleFloat is less than " + thresholdString + "."
	// options[1].Description.Template = fmt.Sprintf("{{if lt .GetValue %s %s}}This option is hidden because the float is less than %s.{{else}}This option is visible because the float is greater than or equal to %s.{{end}}", ids.ExampleFloatID, thresholdString, thresholdString, thresholdString)
	// options[1].Value = nativecfg.ExampleOption_Two
	// options[1].Disabled.Default = true
	// options[1].Disabled.Template = "{{if eq .GetValue " + ids.ExampleBoolID + " true}}false{{else}}{{.UseDefault}}{{end}}"

	// options[2].Name = "Three"
	// options[2].Description.Default = "This is the third option."
	// options[2].Value = nativecfg.ExampleOption_Three

	// // ExampleChoice
	// cfg.ExampleChoice.ID = hdconfig.Identifier(ids.ExampleChoiceID)
	// cfg.ExampleChoice.Name = "Example Choice"
	// cfg.ExampleChoice.Description.Default = "This is an example of a choice parameter between multiple options."
	// cfg.ExampleChoice.Options = options
	// cfg.ExampleChoice.Default = options[0].Value
	// cfg.ExampleChoice.Value = cfg.ExampleChoice.Default
	// cfg.ExampleChoice.AffectedContainers = []string{}

	// // Subconfigs
	// cfg.SubConfig = NewSubConfig()
	// cfg.ServerConfig = NewServerConfig()

	return cfg
}

func ConvertToMetadata(native *sharedconfig.NativeHyperdriveEthereumConfig) *sharedconfig.HyperdriveEthereumConfig {
	cfg := NewHyperdriveEthereumConfig()

	// cfg.ExampleBool.Value = native.ExampleBool
	// cfg.ExampleInt.Value = native.ExampleInt
	// cfg.ExampleUint.Value = native.ExampleUint
	// cfg.ExampleFloat.Value = native.ExampleFloat
	// cfg.ExampleString.Value = native.ExampleString
	// cfg.ExampleChoice.Value = native.ExampleChoice

	// cfg.SubConfig.SubExampleBool.Value = native.SubConfig.SubExampleBool
	// cfg.SubConfig.SubExampleChoice.Value = native.SubConfig.SubExampleChoice

	return cfg
}

// Load the configuration from disk
func (m *AdapterConfigManager) LoadConfigFromDisk() (*sharedconfig.HyperdriveEthereumConfig, error) {
	// Load the native config
	nativeCfg, err := m.nativeConfigManager.LoadConfigFromFile()
	if err != nil {
		return nil, fmt.Errorf("error loading service config: %w", err)
	}

	// Check if the adapter config exists
	_, err = os.Stat(m.adapterConfigPath)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}

	// Load it
	bytes, err := os.ReadFile(m.adapterConfigPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file [%s]: %w", m.adapterConfigPath, err)
	}

	// Deserialize it
	cfgInstance := map[string]any{}
	err = yaml.Unmarshal(bytes, &cfgInstance)
	if err != nil {
		return nil, fmt.Errorf("error deserializing adapter config file [%s]: %w", m.adapterConfigPath, err)
	}
	// serverCfg := NewServerConfig()
	// portInt := cfgInstance[ids.PortID].(int)
	// serverCfg.Port.Value = uint64(portInt)
	// serverCfg.PortMode.Value = PortMode(cfgInstance[ids.PortModeID].(string))

	// Merge the configs
	modCfg := ConvertToMetadata(nativeCfg)
	// modCfg.ServerConfig = serverCfg
	m.AdapterConfig = modCfg
	return modCfg, nil
}

// Save the configuration to a file. If the config hasn't been loaded yet, this doesn't do anything.
func (m *AdapterConfigManager) SaveConfigToDisk() error {
	if m.AdapterConfig == nil {
		return nil
	}

	// Save the native config
	nativeCfg := m.AdapterConfig.ConvertToNative()
	m.nativeConfigManager.Config = nativeCfg
	err := m.nativeConfigManager.SaveConfigToFile()
	if err != nil {
		return fmt.Errorf("error saving service config: %w", err)
	}

	// Serialize the adapter config
	modCfg := hdconfig.CreateInstanceFromMetadata(m.AdapterConfig.ServerConfig)
	bytes, err := yaml.Marshal(modCfg)
	if err != nil {
		return fmt.Errorf("error serializing adapter config: %w", err)
	}

	// Write it
	err = os.WriteFile(m.adapterConfigPath, bytes, ConfigFileMode)
	if err != nil {
		return fmt.Errorf("error writing adapter config file [%s]: %w", m.adapterConfigPath, err)
	}
	return nil
}
