package hdmodule

import (
	"bytes"
	"flag"
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

	var response processConfigResponse
	err = json.Unmarshal(buf.Bytes(), &response)
	assert.NoError(t, err, "Failed to parse JSON output")

	assert.Empty(t, response.Errors, "Expected no errors in the response")

	expectedPortKey := ids.ServerConfigID + "/" + ids.PortModeID
	expectedPorts := map[string]uint16{
		expectedPortKey: 8080,
	}
	assert.Equal(t, expectedPorts, response.Ports, "Expected correct port mapping")
}
