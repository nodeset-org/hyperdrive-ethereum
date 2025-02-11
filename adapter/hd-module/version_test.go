package hdmodule

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/nodeset-org/hyperdrive-ethereum/shared"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	// Create a pipe to capture stdout
	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	defer func() {
		os.Stdout = oldStdout // Restore original stdout
		r.Close()
	}()

	// Redirect stdout to pipe
	os.Stdout = w

	// Run the function
	err := version()
	assert.NoError(t, err, "version() returned an error:: %v", err)

	// Close writer and read output
	w.Close()
	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)

	assert.NoError(t, err, "Failed to read from pipe: %v", err)

	// Parse the captured output
	var response versionResponse
	err = json.Unmarshal(buf.Bytes(), &response)
	assert.NoError(t, err, "Failed to parse JSON output: %v", err)

	// Check the version matches
	expectedVersion := shared.HyperdriveEthereumVersion
	assert.Equal(t, expectedVersion, response.Version, "Expected version %q, got %q", expectedVersion, response.Version)
}
