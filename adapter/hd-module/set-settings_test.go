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
	"github.com/urfave/cli/v2"
)

// type MockAdapterConfigManager struct {
// 	SaveConfigErr error
// }

// func (m *MockAdapterConfigManager) SaveConfigToDisk() error {
// 	return m.SaveConfigErr
// }

var MockNewAdapterConfigManager = func(c *cli.Context) (*config.AdapterConfigManager, error) {
	return &config.AdapterConfigManager{}, nil
}

type MockAdapterConfigManager struct {
	*config.AdapterConfigManager
}

func (m *MockAdapterConfigManager) SaveConfigToDisk() error {
	return fmt.Errorf("mock save error")
}

// func TestSetSettings_Success(t *testing.T) {
// 	// Setup mock handler with valid settings
// 	mockSettings := &hdconfig.HyperdriveSettings{
// 		Modules: map[string]*modconfig.ModuleInstance{
// 			utils.FullyQualifiedModuleName: {
// 				Settings: `{"some_setting": "value"}`,
// 			},
// 		},
// 	}

// 	mockHandler := MockKeyedRequestHandler[*setSettingsRequest]{
// 		ReturnRequest: &setSettingsRequest{Settings: mockSettings},
// 		ReturnError:   nil,
// 	}

// 	// Set up CLI context
// 	app := cli.NewApp()
// 	set := flag.NewFlagSet("test", 0)
// 	ctx := cli.NewContext(app, set, nil)

// 	// Run function
// 	err := setSettings(ctx, mockHandler)
// 	if err != nil {
// 		t.Fatalf("Expected no error, but got: %v", err)
// 	}
// }

// func TestSetSettings_HandleKeyedRequestError(t *testing.T) {
// 	mockHandler := MockKeyedRequestHandler[*setSettingsRequest]{
// 		ReturnRequest: nil,
// 		ReturnError:   fmt.Errorf("mock error"),
// 	}

// 	app := cli.NewApp()
// 	set := flag.NewFlagSet("test", 0)
// 	ctx := cli.NewContext(app, set, nil)

// 	err := setSettings(ctx, mockHandler)
// 	expectedError := "error reading set-settings request: mock error"
// 	if err == nil || err.Error() != expectedError {
// 		t.Errorf("Expected error %q, but got %v", expectedError, err)
// 	}
// }

// func TestSetSettings_MissingModuleConfig(t *testing.T) {
// 	mockHandler := MockKeyedRequestHandler[*setSettingsRequest]{
// 		ReturnRequest: &setSettingsRequest{Settings: &hdconfig.HyperdriveSettings{
// 			Modules: map[string]*modconfig.ModuleInstance{},
// 		}},
// 		ReturnError: nil,
// 	}

// 	app := cli.NewApp()
// 	set := flag.NewFlagSet("test", 0)
// 	ctx := cli.NewContext(app, set, nil)

// 	err := setSettings(ctx, mockHandler)
// 	expectedError := fmt.Sprintf("could not find config for %s", utils.FullyQualifiedModuleName)
// 	if err == nil || err.Error() != expectedError {
// 		t.Errorf("Expected error %q, but got %v", expectedError, err)
// 	}
// }

// func TestSetSettings_DeserializeError(t *testing.T) {
// 	mockHandler := MockKeyedRequestHandler[*setSettingsRequest]{
// 		ReturnRequest: &setSettingsRequest{Settings: &hdconfig.HyperdriveSettings{
// 			Modules: map[string]*modconfig.ModuleInstance{
// 				utils.FullyQualifiedModuleName: {
// 					Settings: `{"some_setting": value_without_quotes}`,
// 				},
// 			},
// 		}},
// 		ReturnError: nil,
// 	}

// 	app := cli.NewApp()
// 	set := flag.NewFlagSet("test", 0)
// 	ctx := cli.NewContext(app, set, nil)

// 	err := setSettings(ctx, mockHandler)
// 	if err == nil || !errors.Is(err, fmt.Errorf("error loading settings:")) {
// 		t.Errorf("Expected JSON deserialization error, but got %v", err)
// 	}
// }

// func TestSetSettings_ConfigManagerError(t *testing.T) {
// 	// Override constructor to simulate failure
// 	originalFunc := config.NewAdapterConfigManager
// 	config.NewAdapterConfigManager = func(c *cli.Context) (*config.AdapterConfigManager, error) {
// 		return nil, fmt.Errorf("mock config manager error")
// 	}
// 	defer func() { config.NewAdapterConfigManager = originalFunc }()

// 	mockHandler := MockKeyedRequestHandler[*setSettingsRequest]{
// 		ReturnRequest: &setSettingsRequest{Settings: &hdconfig.HyperdriveSettings{
// 			Modules: map[string]*modconfig.ModuleInstance{
// 				utils.FullyQualifiedModuleName: {
// 					Settings: `{"some_setting": "value"}`,
// 				},
// 			},
// 		}},
// 		ReturnError: nil,
// 	}

// 	app := cli.NewApp()
// 	set := flag.NewFlagSet("test", 0)
// 	ctx := cli.NewContext(app, set, nil)

// 	err := setSettings(ctx, mockHandler)
// 	expectedError := "error creating config manager: mock config manager error"
// 	if err == nil || err.Error() != expectedError {
// 		t.Errorf("Expected error %q, but got %v", expectedError, err)
// 	}
// }

func TestSetSettings_SaveConfigError(t *testing.T) {
	configDir := "/tmp/hyperdrive-test"

	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create config directory: %v", err)
	}

	defer func() {
		err := os.RemoveAll(configDir)
		if err != nil {
			t.Errorf("Failed to remove test config directory: %v", err)
		}
	}()

	var settingsMap map[string]any
	jsonData := `{"some_setting": "value"}`

	err = json.Unmarshal([]byte(jsonData), &settingsMap)
	if err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

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

	// Define a mock config manager factory that returns an error
	mockConfigManagerFactory := func(c *cli.Context) (config.AdapterConfigManagerInterface, error) {
		mockConfig := &MockAdapterConfigManager{
			AdapterConfigManager: &config.AdapterConfigManager{
				AdapterConfig: &config.HyperdriveEthereumConfigSettings{},
			},
		}
		return mockConfig, nil
	}

	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	set.String(utils.ConfigDirFlag.Name, configDir, "doc")
	_ = set.Parse([]string{"--" + utils.ConfigDirFlag.Name, configDir})

	ctx := cli.NewContext(app, set, nil)

	err = setSettings(ctx, mockHandler, mockConfigManagerFactory)

	expectedError := "error saving config: mock save error"
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error %q, but got %v", expectedError, err)
	}
}
