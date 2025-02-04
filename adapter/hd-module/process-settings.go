package hdmodule

import (
	"context"
	"fmt"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
)

const (
	ProcessSettingsCommandString string = HyperdriveModuleCommand + " process-settings"
)

// Request format for `process-settings`
type ProcessSettingsRequest struct {
	utils.KeyedRequest

	// TODO (HN): Ask Joe Should this be HyperdriveSettings?
	// The config instance to process
	Settings map[string]any `json:"settings"`
}

// Response format for `process-config`
type ProcessSettingsResponse struct {
	// A list of errors that occurred during processing, if any
	Errors []string `json:"errors"`

	// A list of ports that will be exposed
	Ports map[string]uint16 `json:"ports"`
}

// Handle the `process-settings` command
func (c *AdapterClient) ProcessSettings(ctx context.Context, settings map[string]any) (ProcessSettingsResponse, error) {
	request := &ProcessSettingsRequest{
		KeyedRequest: utils.KeyedRequest{
			Key: c.key,
		},
		Settings: settings,
	}
	response := ProcessSettingsResponse{}
	err := runCommand(c, ctx, ProcessSettingsCommandString, request, &response)
	if err != nil {
		return response, fmt.Errorf("error processing module settings: %w", err)
	}
	return response, nil
}
