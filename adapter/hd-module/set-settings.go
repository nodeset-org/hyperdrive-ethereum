package hdmodule

import (
	"context"
	"fmt"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
)

const (
	SetSettingsCommandString string = HyperdriveModuleCommand + " set-settings"
)

// Request format for `set-config`
type SetSettingsRequest struct {
	utils.KeyedRequest

	// The config instance to process
	Settings map[string]any `json:"settings"`
	// Settings *hdconfig.HyperdriveSettings `json:"settings"`
}

// Handle the `set-config` command
func (c *AdapterClient) SetSettings(ctx context.Context, settings map[string]any) error {
	request := &SetSettingsRequest{
		KeyedRequest: utils.KeyedRequest{
			Key: c.key,
		},
		Settings: settings,
	}

	err := runCommand[SetSettingsRequest, struct{}](c, ctx, SetSettingsCommandString, request, nil)
	if err != nil {
		return fmt.Errorf("error setting module settings: %w", err)
	}
	return nil
}
