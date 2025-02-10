package hdmodule

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
	"github.com/nodeset-org/hyperdrive-ethereum/shared"
	"github.com/urfave/cli/v2"
)

// MockKeyedRequestHandler is a generic mock handler for tests
type MockKeyedRequestHandler[RequestType utils.IKeyedRequest] struct {
	ReturnRequest RequestType
	ReturnError   error
}

// Implements the interface by returning the pre-defined request/error
func (m MockKeyedRequestHandler[RequestType]) HandleKeyedRequest(c *cli.Context) (RequestType, error) {
	return m.ReturnRequest, m.ReturnError
}

func TestGetLogFile_AdapterSource(t *testing.T) {
	mockHandler := MockKeyedRequestHandler[*GetLogFileRequest]{
		ReturnRequest: &GetLogFileRequest{Source: "adapter"},
		ReturnError:   nil,
	}

	verifyGetLogFileOutput(t, mockHandler, utils.AdapterLogFile)
}

func TestGetLogFile_ServiceContainerSource(t *testing.T) {
	mockHandler := MockKeyedRequestHandler[*GetLogFileRequest]{
		ReturnRequest: &GetLogFileRequest{Source: shared.ServiceContainerName},
		ReturnError:   nil,
	}

	verifyGetLogFileOutput(t, mockHandler, shared.ServiceLogFile)
}

func TestGetLogFile_HandleKeyedRequestError(t *testing.T) {
	mockHandler := MockKeyedRequestHandler[*GetLogFileRequest]{
		ReturnRequest: &GetLogFileRequest{Source: "adapter"},
		ReturnError:   fmt.Errorf("mock error"),
	}

	// Set up CLI context
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(app, set, nil)

	// Call the function
	err := getLogFile(ctx, mockHandler)

	// Validate error handling
	if err == nil {
		t.Fatalf("Expected an error, but getLogFile() returned nil")
	}

	expectedErrorMsg := "error reading set-settings request: mock error"
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected error message %q, but got %q", expectedErrorMsg, err.Error())
	}
}

func TestGetLogFile_UnknownSource(t *testing.T) {
	mockHandler := MockKeyedRequestHandler[*GetLogFileRequest]{
		ReturnRequest: &GetLogFileRequest{Source: "unknown-source"},
		ReturnError:   nil,
	}

	verifyGetLogFileOutput(t, mockHandler, "")
}

func verifyGetLogFileOutput(t *testing.T, handler MockKeyedRequestHandler[*GetLogFileRequest], expectedPath string) {
	// Capture stdout
	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	os.Stdout = w

	// Set up CLI context
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(app, set, nil)

	// Call the function
	err := getLogFile(ctx, handler)
	if err != nil {
		t.Fatalf("getLogFile() returned an error: %v", err)
	}

	// Restore stdout and read captured output
	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}

	// Parse the captured output
	var response getLogFileResponse
	if err := json.Unmarshal(buf.Bytes(), &response); err != nil {
		t.Fatalf("Failed to parse JSON output: %v", err)
	}

	// Verify the response path
	if response.Path != expectedPath {
		t.Errorf("Expected path %q, got %q", expectedPath, response.Path)
	}
}
