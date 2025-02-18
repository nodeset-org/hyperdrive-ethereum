package hdmodule

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config"

	"github.com/goccy/go-json"
	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config/ids"
	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
	modconfig "github.com/nodeset-org/hyperdrive/modules/config"
	hdconfig "github.com/nodeset-org/hyperdrive/shared/config"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestProcessSettings_Success(t *testing.T) {
	mockSettings := &hdconfig.HyperdriveSettings{
		Modules: map[string]*modconfig.ModuleInstance{
			utils.FullyQualifiedModuleName: {
				Enabled: true,
				Version: "0.1.0",
				Settings: map[string]any{
					"server": map[string]any{
						"portMode": config.PortMode_External,
						"port":     8080,
					},
				},
			},
		},
	}

	mockHandler := MockKeyedRequestHandler[*ProcessSettingsRequest]{
		ReturnRequest: &ProcessSettingsRequest{Settings: mockSettings},
		ReturnError:   nil,
	}
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(app, set, nil)

	err := processSettings(ctx, mockHandler)
	assert.NoError(t, err, "processSettings should not return an error for a valid config")

	w.Close()
	os.Stdout = oldStdout
	_, err = buf.ReadFrom(r)
	assert.NoError(t, err, "Failed to read from pipe")

	var response ProcessSettingsResponse
	err = json.Unmarshal(buf.Bytes(), &response)
	assert.NoError(t, err, "Failed to parse JSON output")

	assert.Empty(t, response.Errors, "Expected no errors in the response")

	expectedPortKey := ids.ServerConfigID + "/" + ids.PortModeID
	expectedPorts := map[string]uint16{
		expectedPortKey: 8080,
	}
	assert.Equal(t, expectedPorts, response.Ports, "Expected correct port mapping")
}

func TestProcessSettings_HandleKeyedRequestError(t *testing.T) {
	mockHandler := MockKeyedRequestHandler[*ProcessSettingsRequest]{
		ReturnRequest: nil,
		ReturnError:   fmt.Errorf("mock error"),
	}

	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(app, set, nil)

	err := processSettings(ctx, mockHandler)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error reading set-settings request: mock error")
}

func TestProcessSettings_MissingModuleConfig(t *testing.T) {
	mockHandler := MockKeyedRequestHandler[*ProcessSettingsRequest]{
		ReturnRequest: &ProcessSettingsRequest{Settings: &hdconfig.HyperdriveSettings{
			// Missing config to simulate failure
			Modules: map[string]*modconfig.ModuleInstance{},
		}},
		ReturnError: nil,
	}

	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(app, set, nil)

	err := processSettings(ctx, mockHandler)
	expectedError := fmt.Sprintf("could not find settings for %s", utils.FullyQualifiedModuleName)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedError)
}

func TestProcessSettings_DeserializeError(t *testing.T) {
	mockHandler := MockKeyedRequestHandler[*ProcessSettingsRequest]{
		ReturnRequest: &ProcessSettingsRequest{Settings: &hdconfig.HyperdriveSettings{
			Modules: map[string]*modconfig.ModuleInstance{
				utils.FullyQualifiedModuleName: {
					Enabled: true,
					Version: "0.1.0",
					Settings: map[string]any{
						"server": map[string]any{
							"portMode": "invalid_value", // bad data
							"port":     "not_a_number",  // bad data
						},
					},
				},
			},
		}},
		ReturnError: nil,
	}

	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(app, set, nil)

	err := processSettings(ctx, mockHandler)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error loading settings")
}

type BrokenMarshalStruct struct{}

func (b BrokenMarshalStruct) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("mock JSON marshalling error")
}

func TestProcessSettings_JSONMarshallingError(t *testing.T) {
	mockHandler := MockKeyedRequestHandler[*ProcessSettingsRequest]{
		ReturnRequest: &ProcessSettingsRequest{Settings: &hdconfig.HyperdriveSettings{
			Modules: map[string]*modconfig.ModuleInstance{
				utils.FullyQualifiedModuleName: {
					Enabled: true,
					Version: "0.1.0",
					Settings: map[string]any{
						"server": BrokenMarshalStruct{},
					},
				},
			},
		}},
		ReturnError: nil,
	}

	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(app, set, nil)

	err := processSettings(ctx, mockHandler)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error loading settings")
	assert.Contains(t, err.Error(), "error serializing module settings to JSON")
	assert.Contains(t, err.Error(), "mock JSON marshalling error")
}
