package hdmodule

import (
	"bytes"
	"os"
	"testing"

	"github.com/goccy/go-json"
	"github.com/nodeset-org/hyperdrive-ethereum/adapter/app"
	hdtemplate "github.com/nodeset-org/hyperdrive/shared/templates"
	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

const ExampleSettingsJson = `{
	"modules": {
		"NodeSet/hyperdrive-ethereum": {
			"enabled": true,
			"version": "0.1.0",
			"settings": {
				"maxInboundPeers": 30,
				"maxOutboundPeers": 50
			}
		}
	}
}`

func TestCallConfigFunction(t *testing.T) {
	// Build the request
	request := CallConfigFunctionRequest{
		FuncName: "GetMaxPeers",
	}
	err := json.Unmarshal([]byte(ExampleSettingsJson), &request.Settings)
	require.NoError(t, err)

	// Marshal the request to JSON
	inputBytes, err := json.Marshal(request)
	require.NoError(t, err)

	// Add newline to simulate ENTER key
	inputBytes = append(inputBytes, '\n')

	// Redirect stdout to a new pipe
	oldStdout := os.Stdout
	rOut, wOut, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = wOut
	defer func() {
		os.Stdout = oldStdout
	}()

	// Run the function
	app := app.CreateApp()
	app.Reader = bytes.NewReader(inputBytes)

	ctx := cli.NewContext(app, nil, nil)
	err = callConfigFunction(ctx)
	require.NoError(t, err)

	// Read from stdout
	wOut.Close()
	outputBuf := new(bytes.Buffer)
	_, err = outputBuf.ReadFrom(rOut)
	require.NoError(t, err)

	var result hdtemplate.CallConfigFunctionResponse
	err = json.Unmarshal(outputBuf.Bytes(), &result)
	require.NoError(t, err)

	// MaxInboundPeers + MaxOutboundPeers = 30 + 50 = 80
	require.Equal(t, "80", result.Result)
	t.Log("call-config-function ran successfully")
}
