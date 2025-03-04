package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	ConfigFileMode os.FileMode = 0644
)

type NativeHyperdriveEthereumSettings struct {
	EnableIPv6               bool    `json:"enableIPv6" yaml:"enableIPv6"`
	ProjectName              string  `json:"projectName" yaml:"projectName"`
	ApiPort                  uint    `json:"apiPort" yaml:"apiPort"`
	UserDataPath             string  `json:"userDataPath" yaml:"userDataPath"`
	AutoTxMaxFee             float64 `json:"autoTxMaxFee" yaml:"autoTxMaxFee"`
	MaxPriorityFee           float64 `json:"maxPriorityFee" yaml:"maxPriorityFee"`
	AutoTxGasThreshold       float64 `json:"autoTxGasThreshold" yaml:"autoTxGasThreshold"`
	AdditionalDockerNetworks string  `json:"additionalDockerNetworks" yaml:"additionalDockerNetworks"`
	ClientTimeout            uint    `json:"clientTimeout" yaml:"clientTimeout"`

	Network    Network    `json:"network" yaml:"network"`
	ClientMode ClientMode `json:"clientMode" yaml:"clientMode"`

	ContainerTag string `json:"containerTag" yaml:"containerTag"`
}

// Configuration manager
type ConfigManager struct {
	// The configuration
	Config *NativeHyperdriveEthereumSettings

	// The path to the configuration file
	ConfigPath string
}

// Create a new configuration manager
func NewConfigManager(path string) *ConfigManager {
	return &ConfigManager{
		ConfigPath: path,
	}
}

// Load the configuration from a file
func (m *ConfigManager) LoadConfigFromFile() (*NativeHyperdriveEthereumSettings, error) {
	// Check if the file exists
	_, err := os.Stat(m.ConfigPath)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}

	// Load it
	bytes, err := os.ReadFile(m.ConfigPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file [%s]: %w", m.ConfigPath, err)
	}

	// Deserialize it
	cfg := NativeHyperdriveEthereumSettings{}
	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		return nil, fmt.Errorf("error deserializing config file [%s]: %w", m.ConfigPath, err)
	}
	m.Config = &cfg
	return &cfg, nil
}

// Save the configuration to a file. If the config hasn't been loaded yet, this doesn't do anything.
func (m *ConfigManager) SaveConfigToFile() error {
	if m.Config == nil {
		return nil
	}

	// Serialize it
	bytes, err := yaml.Marshal(m.Config)
	if err != nil {
		return fmt.Errorf("error serializing config: %w", err)
	}

	// Write it
	err = os.WriteFile(m.ConfigPath, bytes, ConfigFileMode)
	if err != nil {
		return fmt.Errorf("error writing config file [%s]: %w", m.ConfigPath, err)
	}
	return nil
}
