package hdmodule

import (
	"bytes"
	"encoding/json"
	"flag"
	"os"
	"testing"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
	"github.com/urfave/cli/v2"
)

type MockKeyedRequestHandler[RequestType utils.IKeyedRequest] struct {
	ReturnRequest RequestType
	ReturnError   error
}

func (m MockKeyedRequestHandler[RequestType]) HandleKeyedRequest(c *cli.Context) (RequestType, error) {
	return m.ReturnRequest, m.ReturnError
}

// TestGetLogFile_AdapterSource tests the getLogFile function with "adapter" source
func TestGetLogFile_AdapterSource(t *testing.T) {
	mockHandler := MockKeyedRequestHandler[*GetLogFileRequest]{
		ReturnRequest: &GetLogFileRequest{Source: "adapter"},
		ReturnError:   nil,
	}

	// Capture stdout
	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	os.Stdout = w

	// Set up CLI context with "adapter" source
	app := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(app, set, nil)

	// Call the function
	err := getLogFile(ctx, mockHandler)
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

	// Verify the response
	expectedPath := utils.AdapterLogFile
	if response.Path != expectedPath {
		t.Errorf("Expected path %q, got %q", expectedPath, response.Path)
	}
}
