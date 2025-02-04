package hdmodule

import (
	"context"
	"fmt"

	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
)

const (
	GetConfigMetadataCommandString string = HyperdriveModuleCommand + " get-config-metadata"
)

// Get the module config metadata from the adapter
func (c *AdapterClient) GetConfigMetadata(ctx context.Context) (hdconfig.IModuleConfiguration, error) {
	request := &utils.KeyedRequest{
		Key: c.key,
	}
	configMap := map[string]any{}

	// Get the config from the adapter
	err := runCommand(c, ctx, GetConfigMetadataCommandString, request, &configMap)
	if err != nil {
		return nil, fmt.Errorf("error getting configuration metadata: %w", err)
	}

	// Unmarshal the config from the map
	response, err := hdconfig.UnmarshalConfigurationFromMap(configMap)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling configuration metadata: %w", err)
	}
	return response, nil
}
