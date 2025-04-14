package hdmodule

import (
	"encoding/json"
	"fmt"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config"
	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
	sharedIds "github.com/nodeset-org/hyperdrive-ethereum/shared/ids"

	sharedconfig "github.com/nodeset-org/hyperdrive-ethereum/shared/config"
	hdconfig "github.com/nodeset-org/hyperdrive/config"
	modconfig "github.com/nodeset-org/hyperdrive/modules/config"
	"github.com/urfave/cli/v2"
)

type ProcessSettingsRequest struct {
	// The current config settings
	CurrentSettings *hdconfig.HyperdriveSettings `json:"currentSettings"`

	// The new (proposed) config settings
	NewSettings *hdconfig.HyperdriveSettings `json:"newSettings"`
}

type ProcessSettingsResponse struct {
	Errors []string `json:"errors"`

	// A list of ports that will be exposed
	Ports map[string]uint16 `json:"ports"`

	// A list of services that need to be restarted as a result of the new settings
	ServicesToRestart []string `json:"servicesToRestart"`
}

func processSettings(c *cli.Context) error {
	request, err := utils.HandleRequest[*ProcessSettingsRequest](c)
	if err != nil {
		return fmt.Errorf("error reading process-settings request: %w", err)
	}

	// Process the settings
	response, err := processSettingsImpl(request.CurrentSettings, request.NewSettings)
	if err != nil {
		return err
	}

	// Marshal it
	bytes, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("error marshalling process-settings response: %w", err)
	}

	// Print it
	fmt.Println(string(bytes))
	return nil
}

// Process the settings
func processSettingsImpl(oldHdSettings *hdconfig.HyperdriveSettings, newHdSettings *hdconfig.HyperdriveSettings) (*ProcessSettingsResponse, error) {
	// Construct the old (current) module settings from the Hyperdrive config
	var oldSettings = config.NewHyperdriveEthereumConfigSettings()
	oldModInstance, exists := oldHdSettings.Modules[utils.FullyQualifiedModuleName]
	if !exists {
		// Create an instance with the default settings
		cfg := config.NewHyperdriveEthereumConfig()
		oldModSettings := modconfig.CreateModuleSettings(cfg)
		err := oldModSettings.ConvertToKnownType(&oldSettings)
		if err != nil {
			return nil, fmt.Errorf("error creating default settings: %w", err)
		}
	} else {
		err := oldModInstance.DeserializeSettingsIntoKnownType(&oldSettings)
		if err != nil {
			return nil, fmt.Errorf("error loading old settings: %w", err)
		}
	}

	// Construct the new (proposed) module settings from the Hyperdrive config
	var newSettings = config.NewHyperdriveEthereumConfigSettings()

	newModInstance, exists := newHdSettings.Modules[utils.FullyQualifiedModuleName]
	if !exists {
		return nil, fmt.Errorf("could not find new settings for %s", utils.FullyQualifiedModuleName)
	}
	err := newModInstance.DeserializeSettingsIntoKnownType(&newSettings)
	if err != nil {
		return nil, fmt.Errorf("error loading new settings: %w", err)
	}

	// This is where any examples of validation will go when added
	errors := []string{}

	// Get the open ports
	ports := map[string]uint16{}
	// TODO: Verify with Joe
	if newSettings.ServerConfig != nil && newSettings.ServerConfig.PortMode != sharedconfig.PortMode_Closed {
		ports[sharedIds.LocalServerConfigID+"/"+sharedIds.PortID] = uint16(newSettings.ServerConfig.Port)
	}

	// Get the list of services that need to be restarted
	servicesToRestart, err := newSettings.GetChangedServices(oldSettings)
	if err != nil {
		return nil, fmt.Errorf("error getting changed services: %w", err)
	}
	if servicesToRestart == nil {
		servicesToRestart = []string{}
	}

	// Create the response
	response := &ProcessSettingsResponse{
		Errors:            errors,
		Ports:             ports,
		ServicesToRestart: servicesToRestart,
	}
	return response, nil
}
