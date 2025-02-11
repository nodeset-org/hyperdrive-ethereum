package hdmodule

import (
	"fmt"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config"
	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
	hdconfig "github.com/nodeset-org/hyperdrive/shared/config"

	"github.com/urfave/cli/v2"
)

// Request format for `set-config`
type setSettingsRequest struct {
	utils.KeyedRequest

	// The config instance to process
	Settings *hdconfig.HyperdriveSettings `json:"settings"`
}

// Handle the `set-config` command
func setSettings(
	c *cli.Context,
	handler utils.KeyedRequestHandler[*setSettingsRequest],
	configManagerFactory func(*cli.Context) (config.AdapterConfigManagerInterface, error),
) error {
	// Get the request
	request, err := handler.HandleKeyedRequest(c)
	if err != nil {
		return fmt.Errorf("error reading set-settings request: %w", err)
	}

	// Construct the module settings from the Hyperdrive config
	modInstance, exists := request.Settings.Modules[utils.FullyQualifiedModuleName]
	if !exists {
		return fmt.Errorf("could not find config for %s", utils.FullyQualifiedModuleName)
	}

	var settings config.HyperdriveEthereumConfigSettings
	err = modInstance.DeserializeSettingsIntoKnownType(&settings)
	if err != nil {
		return fmt.Errorf("error loading settings: %w", err)
	}

	// Use injected config manager factory
	cfgMgr, err := configManagerFactory(c)
	if err != nil {
		return fmt.Errorf("error creating config manager: %w", err)
	}
	err = cfgMgr.SetAdapterConfig(&settings)
	if err != nil {
		return fmt.Errorf("error setting config: %w", err)
	}

	// Save it
	err = cfgMgr.SaveConfigToDisk()
	if err != nil {
		return fmt.Errorf("error saving config: %w", err)
	}

	return nil
}
