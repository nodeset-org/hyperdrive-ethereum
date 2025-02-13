package hdmodule

import (
	"encoding/json"
	"fmt"

	hdconfig "github.com/nodeset-org/hyperdrive/shared/config"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config"
	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config/ids"
	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
	"github.com/urfave/cli/v2"
)

type ProcessSettingsRequest struct {
	utils.KeyedRequest

	Settings *hdconfig.HyperdriveSettings `json:"settings"`
}

type processConfigResponse struct {
	Errors []string `json:"errors"`

	Ports map[string]uint16 `json:"ports"`
}

func processSettings(
	c *cli.Context,
	handler utils.KeyedRequestHandler[*ProcessSettingsRequest]) error {
	request, err := handler.HandleKeyedRequest(c)
	if err != nil {
		return fmt.Errorf("error reading set-settings request: %w", err)
	}
	modInstance, exists := request.Settings.Modules[utils.FullyQualifiedModuleName]
	if !exists {
		return fmt.Errorf("could not find settings for %s", utils.FullyQualifiedModuleName)
	}
	var settings config.HyperdriveEthereumConfigSettings
	err = modInstance.DeserializeSettingsIntoKnownType(&settings)
	if err != nil {
		return fmt.Errorf("error loading settings: %w", err)
	}

	errors := []string{}

	ports := map[string]uint16{}

	if settings.ServerConfig.PortMode != config.PortMode_Closed {
		ports[ids.ServerConfigID+"/"+ids.PortModeID] = uint16(settings.ServerConfig.Port)
	}

	response := processConfigResponse{
		Errors: errors,
		Ports:  ports,
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("error marshalling process-config response: %w", err)
	}

	fmt.Println(string(bytes))
	return nil
}
