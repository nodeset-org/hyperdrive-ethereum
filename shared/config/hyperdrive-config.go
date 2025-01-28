package config

import (
	"fmt"
	"os"
	"path/filepath"

	"al.essio.dev/pkg/shellescape"

	sharedconfig "github.com/nodeset-org/hyperdrive/shared/config"
	"gopkg.in/yaml.v2"
)

// Load the Hyperdrive configuration from a file; the Hyperdrive user directory will be set to the directory containing the config file
func LoadFromFile(configFilePath string, systemDir string) (*HyperdriveEthereumConfig, error) {
	// Return nil if the file doesn't exist
	_, err := os.Stat(configFilePath)
	if os.IsNotExist(err) {
		return nil, nil
	}

	// Read the file
	configBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not read Hyperdrive settings file at %s: %w", shellescape.Quote(configFilePath), err)
	}

	// Attempt to parse it out into a settings map
	var settings map[string]any
	if err := yaml.Unmarshal(configBytes, &settings); err != nil {
		return nil, fmt.Errorf("could not parse config file: %w", err)
	}

	// Deserialize it into a config object
	cfg := NewHyperdriveEthereumConfig(filepath.Dir(configFilePath), systemDir)
	err = cfg.deserialize(settings)
	if err != nil {
		return nil, fmt.Errorf("error deserializing Hyperdrive config: %w", err)
	}
	return cfg, nil
}

// Deserialize the Hyperdrive configuration from an instance
func (cfg *HyperdriveEthereumConfig) deserialize(instance map[string]any) error {
	err := config.UnmarshalConfigurationInstance(instance, cfg)
	if err != nil {
		return fmt.Errorf("error deserializing Hyperdrive config: %w", err)
	}

	// Deserialize the version
	version, ok := instance[versionName].(string)
	if !ok {
		return fmt.Errorf("hyperdrive version is not a string, it's a %T", instance[versionName])
	}
	cfg.Version = version

	// Deserialize the module enable status
	moduleEnabledMap, ok := instance[moduleEnabledMapName].(map[string]any)
	if !ok {
		return fmt.Errorf("module enable status is not a map, it's a %T", instance[moduleEnabledMapName])
	}
	for module, enabled := range moduleEnabledMap {
		enabledBool, ok := enabled.(bool)
		if !ok {
			return fmt.Errorf("module enable status for %s is not a bool, it's a %T", module, enabled)
		}
		cfg.moduleEnableStatus[module] = enabledBool
	}
	return nil
}

// Serialize the Hyperdrive configuration to an instance
// func (cfg *HyperdriveEthereumConfig) serialize() map[string]any {
// 	instance := config.CreateInstance(cfg)
// 	instance[versionName] = cfg.Version

// 	// Serialize the module enable status
// 	moduleEnabledMap := map[string]any{}
// 	for module, enabled := range cfg.moduleEnableStatus {
// 		moduleEnabledMap[module] = enabled
// 	}
// 	instance[moduleEnabledMapName] = moduleEnabledMap
// 	return instance
// }

// Load the configurations for each module - any errors for individual modules is reported in the module config container instead of returned here, so make sure to check those when interacting with them.
func (cfg *HyperdriveEthereumConfig) LoadModuleConfigs() error {
	moduleLoader, err := sharedconfig.NewModuleConfigLoader(cfg.ProjectName.Value, cfg.GetModuleSystemDir(), cfg.GetAdapterKeyPath())
	if err != nil {
		return fmt.Errorf("error creating module config loader: %w", err)
	}
	modules, err := moduleLoader.LoadModuleConfigs()
	if err != nil {
		return fmt.Errorf("error loading module configs: %w", err)
	}

	// Add any missing modules to the config's enable status
	for _, module := range modules {
		if _, exists := cfg.moduleEnableStatus[string(module.Descriptor.Name)]; !exists {
			cfg.moduleEnableStatus[string(module.Descriptor.Name)] = false
		}
	}

	// Set the module enable status
	for _, module := range modules {
		module.Enabled = cfg.moduleEnableStatus[string(module.Descriptor.Name)]
	}
	cfg.ModuleConfigs = modules
	return nil
}

func (cfg *HyperdriveEthereumConfig) GetModuleSystemDir() string {
	return filepath.Join(cfg.systemPath, "modules")
}

func (cfg *HyperdriveEthereumConfig) GetAdapterKeyPath() string {
	return filepath.Join(cfg.hyperdriveUserDirectory, "secrets", "adapter.key")
}

func (cfg *HyperdriveEthereumConfig) ConvertToNative() *NativeHyperdriveEthereumConfig {
	native := &NativeHyperdriveEthereumConfig{}

	// native.ExampleBool = cfg.ExampleBool.Value
	// native.ExampleInt = cfg.ExampleInt.Value
	// native.ExampleUint = cfg.ExampleUint.Value
	// native.ExampleFloat = cfg.ExampleFloat.Value
	// native.ExampleString = cfg.ExampleString.Value
	// native.ExampleChoice = cfg.ExampleChoice.Value
	// native.SubConfig.SubExampleBool = cfg.SubConfig.SubExampleBool.Value
	// native.SubConfig.SubExampleChoice = cfg.SubConfig.SubExampleChoice.Value
	return native
}
