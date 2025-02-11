package hdmodule

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/nodeset-org/hyperdrive-ethereum/shared"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestGetContainers_ValidResponse(t *testing.T) {
	// Capture stdout
	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	os.Stdout = w

	// Create a fake CLI context
	app := cli.NewApp()
	ctx := cli.NewContext(app, nil, nil)

	// Call the function
	err := getContainers(ctx)
	assert.NoError(t, err, "getContainers() returned an error: %v", err)

	// Read captured output
	w.Close()
	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	assert.NoError(t, err, "Failed to read from pipe: %v", err)

	// Parse the JSON response
	var response getContainersResponse
	if err := json.Unmarshal(buf.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse JSON output: %v", err)
	}

	// Verify the container list contains the expected values
	expectedContainers := []string{shared.ServiceContainerName}
	assert.Equal(t, expectedContainers, response.Containers, "Expected and actual containers should match")

	for i, container := range expectedContainers {
		assert.Equal(t, container, response.Containers[i], "Expected container %q, got %q", container, response.Containers[i])
	}
}
