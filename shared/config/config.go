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

// Example of a configuration for a service
type NativeHyperdriveEthereumConfig struct {
	// TODO
	ExampleBool bool `json:"exampleBool" yaml:"exampleBool"`

	ExampleInt int64 `json:"exampleInt" yaml:"exampleInt"`

	ExampleUint uint64 `json:"exampleUint" yaml:"exampleUint"`

	ExampleFloat float64 `json:"exampleFloat" yaml:"exampleFloat"`

	ExampleString string `json:"exampleString" yaml:"exampleString"`
}

// Configuration manager
type ConfigManager struct {
	// The configuration
	Config *NativeHyperdriveEthereumConfig

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
func (m *ConfigManager) LoadConfigFromFile() (*NativeHyperdriveEthereumConfig, error) {
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
	cfg := NativeHyperdriveEthereumConfig{}
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
