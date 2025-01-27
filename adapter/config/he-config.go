package config

import (
	"fmt"
	"path/filepath"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"

	sharedconfig "github.com/nodeset-org/hyperdrive-ethereum/shared/config"
	"github.com/urfave/cli/v2"
)

type PortMode string

const (
	PortMode_Closed    PortMode = "closed"
	PortMode_Localhost PortMode = "localhost"
	PortMode_External  PortMode = "external"
)

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
