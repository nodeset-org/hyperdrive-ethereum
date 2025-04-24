package hdmodule

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestGetConfigMetadata(t *testing.T) {
	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	defer func() {
		os.Stdout = oldStdout
		r.Close()
	}()
	os.Stdout = w

	app := cli.NewApp()
	ctx := cli.NewContext(app, nil, nil)

	err := getConfigMetadata(ctx)
	assert.NoError(t, err, "getConfigMetadata() returned an error: %v", err)

	w.Close()
	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	assert.NoError(t, err, "Failed to read from pipe: %v", err)

	var configMap map[string]interface{}
	if err := json.Unmarshal(buf.Bytes(), &configMap); err != nil {
		t.Fatalf("Failed to parse JSON output: %v", err)
	}

	expectedConfig := hdconfig.MarshalConfigurationToMap(config.NewHyperdriveEthereumConfig())
	for key := range expectedConfig {
		if _, exists := configMap[key]; !exists {
			t.Errorf("Expected key %q in config, but not found", key)
		}
	}
}
