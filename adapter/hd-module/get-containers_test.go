package hdmodule

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/nodeset-org/hyperdrive-ethereum/shared"
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
	if err != nil {
		t.Fatalf("getContainers() returned an error: %v", err)
	}

	// Read captured output
	w.Close()
	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}

	// Parse the JSON response
	var response getContainersResponse
	if err := json.Unmarshal(buf.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse JSON output: %v", err)
	}

	// Verify the container list contains the expected values
	expectedContainers := []string{shared.ServiceContainerName}
	if len(response.Containers) != len(expectedContainers) {
		t.Errorf("Expected %d containers, got %d", len(expectedContainers), len(response.Containers))
	}

	for i, container := range expectedContainers {
		if response.Containers[i] != container {
			t.Errorf("Expected container %q, got %q", container, response.Containers[i])
		}
	}
}
