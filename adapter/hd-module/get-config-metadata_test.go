package hdmodule

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
	"github.com/urfave/cli/v2"
)

func TestGetConfigMetadata(t *testing.T) {
	// Create a pipe to capture stdout
	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	defer func() {
		os.Stdout = oldStdout
		r.Close()
	}()
	os.Stdout = w

	// Create a fake CLI context
	app := cli.NewApp()
	ctx := cli.NewContext(app, nil, nil)

	// Call the function
	err := getConfigMetadata(ctx)
	if err != nil {
		t.Fatalf("getConfigMetadata() returned an error: %v", err)
	}

	// Close writer and read output
	w.Close()
	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}

	// Parse the captured output
	var configMap map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &configMap); err != nil {
		t.Fatalf("Failed to parse JSON output: %v", err)
	}

	// Verify it contains expected configuration keys
	expectedConfig := hdconfig.MarshalConfigurationToMap(config.NewHyperdriveEthereumConfig())
	for key := range expectedConfig {
		if _, exists := configMap[key]; !exists {
			t.Errorf("Expected key %q in config, but not found", key)
		}
	}
}
