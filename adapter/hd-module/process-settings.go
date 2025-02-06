package hdmodule

import (
	"encoding/json"
	"fmt"

	hdconfig "github.com/nodeset-org/hyperdrive/shared/config"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config"
	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
	"github.com/urfave/cli/v2"
)

// Request format for `process-settings`
type processSettingsRequest struct {
	utils.KeyedRequest

	// The config instance to process
	Settings *hdconfig.HyperdriveSettings `json:"settings"`
}

// Response format for `process-config`
type processConfigResponse struct {
	// A list of errors that occurred during processing, if any
	Errors []string `json:"errors"`

	// A list of ports that will be exposed
	Ports map[string]uint16 `json:"ports"`
}

// Handle the `process-settings` command
func processSettings(c *cli.Context) error {
	// Get the request
	request, err := utils.HandleKeyedRequest[*processSettingsRequest](c)
	if err != nil {
		return fmt.Errorf("error reading set-settings request: %w", err)
	}
	// TODO: (HN)
	// Construct the module settings from the Hyperdrive config
	modInstance, exists := request.Settings.Modules[utils.FullyQualifiedModuleName]
	if !exists {
		return fmt.Errorf("could not find settings for %s", utils.FullyQualifiedModuleName)
	}
	var settings config.HyperdriveEthereumConfigSettings
	err = modInstance.DeserializeSettingsIntoKnownType(&settings)
	if err != nil {
		return fmt.Errorf("error loading settings: %w", err)
	}

	// This is where any examples of validation will go when added
	errors := []string{}

	// Get the open ports
	ports := map[string]uint16{}

	// if cfg.ServerConfig.PortMode.Value != config.PortMode_Closed {
	// 	ports[ids.ServerConfigID+"/"+ids.PortModeID] = uint16(cfg.ServerConfig.Port.Value)
	// }

	// Create the response
	response := processConfigResponse{
		Errors: errors,
		Ports:  ports,
	}

	// Marshal it
	bytes, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("error marshalling process-config response: %w", err)
	}

	// Print it
	fmt.Println(string(bytes))
	return nil
}
