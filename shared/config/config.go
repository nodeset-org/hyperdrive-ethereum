package config

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
