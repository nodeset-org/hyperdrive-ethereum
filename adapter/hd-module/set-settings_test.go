package hdmodule

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/goccy/go-json"
	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config"
	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
	modconfig "github.com/nodeset-org/hyperdrive/modules/config"
	hdconfig "github.com/nodeset-org/hyperdrive/shared/config"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

type MockAdapterConfigManager struct {
	*config.AdapterConfigManager
}

func (m *MockAdapterConfigManager) SaveConfigToDisk() error {
	return nil
}

type MockAdapterConfigManagerSaveToDiskError struct {
	*config.AdapterConfigManager
}

func (m *MockAdapterConfigManagerSaveToDiskError) SaveConfigToDisk() error {
	return fmt.Errorf("mock save error")
}

func TestSetSettings_Success(t *testing.T) {
	configDir := "/tmp/hyperdrive-test"

	err := os.MkdirAll(configDir, 0755)
	assert.NoError(t, err, "Failed to create config directory: %v", err)

	defer os.RemoveAll(configDir)

	var settingsMap map[string]any
	jsonData := `{"some_setting": "value"}`

	err = json.Unmarshal([]byte(jsonData), &settingsMap)
	assert.NoError(t, err, "Failed to parse JSON: %v", err)

	mockHandler := MockKeyedRequestHandler[*setSettingsRequest]{
		ReturnRequest: &setSettingsRequest{Settings: &hdconfig.HyperdriveSettings{
			Modules: map[string]*modconfig.ModuleInstance{
				utils.FullyQualifiedModuleName: {
					Enabled:  true,
					Version:  "0.1.0",
					Settings: settingsMap,
				},
			},
		}},
		ReturnError: nil,
	}

	mockConfigManagerFactory := func(c *cli.Context) (config.AdapterConfigManagerInterface, error) {
		return &MockAdapterConfigManager{
			AdapterConfigManager: &config.AdapterConfigManager{
				AdapterConfig: &config.HyperdriveEthereumConfigSettings{},
			},
		}, nil
	}

	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)

	ctx := cli.NewContext(app, set, nil)

	err = setSettings(ctx, mockHandler, mockConfigManagerFactory)
	assert.NoError(t, err, "Expected no error, but got: %v", err)
}

func TestSetSettings_HandleKeyedRequestError(t *testing.T) {
	mockHandler := MockKeyedRequestHandler[*setSettingsRequest]{
		ReturnRequest: nil,
		ReturnError:   fmt.Errorf("mock error"),
	}

	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(app, set, nil)

	err := setSettings(ctx, mockHandler, func(*cli.Context) (config.AdapterConfigManagerInterface, error) {
		return &MockAdapterConfigManager{}, nil
	})

	expectedError := "error reading set-settings request: mock error"
	assert.EqualError(t, err, expectedError, "Expected a specific error message for HandleKeyedRequest failure")

}

func TestSetSettings_MissingModuleConfig(t *testing.T) {
	mockHandler := MockKeyedRequestHandler[*setSettingsRequest]{
		ReturnRequest: &setSettingsRequest{Settings: &hdconfig.HyperdriveSettings{
			Modules: map[string]*modconfig.ModuleInstance{},
		}},
		ReturnError: nil,
	}

	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(app, set, nil)

	err := setSettings(ctx, mockHandler, func(*cli.Context) (config.AdapterConfigManagerInterface, error) {
		return &MockAdapterConfigManager{}, nil
	})

	expectedError := fmt.Sprintf("could not find config for %s", utils.FullyQualifiedModuleName)
	assert.EqualError(t, err, expectedError, "Expected error for missing module config")

}

func TestSetSettings_DeserializeError(t *testing.T) {
	mockHandler := MockKeyedRequestHandler[*setSettingsRequest]{
		ReturnRequest: &setSettingsRequest{Settings: &hdconfig.HyperdriveSettings{
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

	err := setSettings(ctx, mockHandler, func(*cli.Context) (config.AdapterConfigManagerInterface, error) {
		return &MockAdapterConfigManager{
			AdapterConfigManager: &config.AdapterConfigManager{
				AdapterConfig: &config.HyperdriveEthereumConfigSettings{},
			},
		}, nil
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error loading settings")
	assert.Contains(t, err.Error(), "error serializing module settings to JSON")
	assert.Contains(t, err.Error(), "mock JSON marshalling error") // Ensure root cause is present
}

func TestSetSettings_SaveConfigError(t *testing.T) {
	configDir := "/tmp/hyperdrive-test"
	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create config directory: %v", err)
	}
	defer os.RemoveAll(configDir)

	var settingsMap map[string]any
	jsonData := `{"some_setting": "value"}`

	err = json.Unmarshal([]byte(jsonData), &settingsMap)
	assert.NoError(t, err, "Failed to parse JSON: %v", err)

	mockHandler := MockKeyedRequestHandler[*setSettingsRequest]{
		ReturnRequest: &setSettingsRequest{Settings: &hdconfig.HyperdriveSettings{
			Modules: map[string]*modconfig.ModuleInstance{
				utils.FullyQualifiedModuleName: {
					Enabled:  true,
					Version:  "0.1.0",
					Settings: settingsMap,
				},
			},
		}},
		ReturnError: nil,
	}

	mockConfigManagerFactory := func(c *cli.Context) (config.AdapterConfigManagerInterface, error) {
		return &MockAdapterConfigManagerSaveToDiskError{
			AdapterConfigManager: &config.AdapterConfigManager{
				AdapterConfig: &config.HyperdriveEthereumConfigSettings{},
			},
		}, nil
	}

	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)

	ctx := cli.NewContext(app, set, nil)

	err = setSettings(ctx, mockHandler, mockConfigManagerFactory)

	expectedError := "error saving config: mock save error"
	assert.EqualError(t, err, expectedError, "Expected error for failed config save")

}
