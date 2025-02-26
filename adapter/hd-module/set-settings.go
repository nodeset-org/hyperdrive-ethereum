package hdmodule

import (
	"fmt"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config"
	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
	hdconfig "github.com/nodeset-org/hyperdrive/shared/config"

	"github.com/urfave/cli/v2"
)

type SetSettingsRequest struct {
	utils.KeyedRequest

	Settings *hdconfig.HyperdriveSettings `json:"settings"`
}

func setSettings(
	c *cli.Context,
	handler utils.KeyedRequestHandler[*SetSettingsRequest],
	cfgMgr config.AdapterConfigManagerInterface,
) error {
	request, err := handler.HandleKeyedRequest(c)
	if err != nil {
		return fmt.Errorf("error reading set-settings request: %w", err)
	}

	modInstance, exists := request.Settings.Modules[utils.FullyQualifiedModuleName]
	if !exists {
		return fmt.Errorf("could not find config for %s", utils.FullyQualifiedModuleName)
	}

	var settings config.HyperdriveEthereumConfigSettings
	err = modInstance.DeserializeSettingsIntoKnownType(&settings)
	if err != nil {
		return fmt.Errorf("error loading settings: %w", err)
	}

	err = cfgMgr.SetAdapterConfig(&settings)
	if err != nil {
		return fmt.Errorf("error setting config: %w", err)
	}

	err = cfgMgr.SaveConfigToDisk()
	if err != nil {
		return fmt.Errorf("error saving config: %w", err)
	}

	return nil
}
