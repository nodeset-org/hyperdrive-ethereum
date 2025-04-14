package hdmodule

import (
	"fmt"

	"github.com/goccy/go-json"
	sharedconfig "github.com/nodeset-org/hyperdrive-ethereum/shared/config"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
	hdconfig "github.com/nodeset-org/hyperdrive/config"
	hdtemplate "github.com/nodeset-org/hyperdrive/shared/templates"

	"github.com/urfave/cli/v2"
)

type CallConfigFunctionRequest struct {
	FuncName string                       `json:"funcName"`
	Settings *hdconfig.HyperdriveSettings `json:"settings"`
}

func callConfigFunction(c *cli.Context) error {
	request, err := utils.HandleRequest[*CallConfigFunctionRequest](c)
	if err != nil {
		return err
	}

	modInstance, exists := request.Settings.Modules[utils.FullyQualifiedModuleName]
	if !exists {
		return fmt.Errorf("could not find settings for module %s", utils.FullyQualifiedModuleName)
	}

	var settings sharedconfig.RethConfigSettings
	err = modInstance.DeserializeSettingsIntoKnownType(&settings)
	if err != nil {
		return fmt.Errorf("error loading settings: %w", err)
	}

	switch request.FuncName {
	case "GetMaxPeers":
		response := hdtemplate.CallConfigFunctionResponse{
			Result: fmt.Sprintf("%d", settings.GetMaxPeers()),
		}
		bytes, err := json.Marshal(response)
		if err != nil {
			return fmt.Errorf("error marshalling derived value: %w", err)
		}
		fmt.Println(string(bytes))
		return nil
	default:
		return fmt.Errorf("unknown function: %s", request.FuncName)
	}
}
