package hdmodule

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
	"github.com/urfave/cli/v2"
)

// Mocked HandleKeyedRequest to simulate different scenarios
func MockHandleKeyedRequest(c *cli.Context) (*GetLogFileRequest, error) {
	return &GetLogFileRequest{Source: "adapter"}, fmt.Errorf("mock error")
}

// TestGetLogFile_AdapterSource tests the getLogFile function with "adapter" source
func TestGetLogFile_AdapterSource(t *testing.T) {
	// Backup the original function
	originalFunc := utils.HandleKeyedRequest[*GetLogFileRequest]

	// Restore after test
	defer func() { utils.HandleKeyedRequest = originalFunc }()

	// Assign our mock function to a variable and use it explicitly
	utils.HandleKeyedRequest = MockHandleKeyedRequest

	// Capture stdout
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Set up CLI context with "adapter" source
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	set.String("source", "adapter", "doc")
	ctx := cli.NewContext(app, set, nil)

	// Call the function
	err := getLogFile(ctx)
	if err != nil {
		t.Fatalf("getLogFile() returned an error: %v", err)
	}

	// Restore stdout and read captured output
	w.Close()
	os.Stdout = oldStdout
	_, err = buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}

	// Parse the captured output
	var response getLogFileResponse
	if err := json.Unmarshal(buf.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse JSON output: %v", err)
	}

	// Verify the response
	expectedPath := utils.AdapterLogFile
	if response.Path != expectedPath {
		t.Errorf("Expected path %q, got %q", expectedPath, response.Path)
	}
}
