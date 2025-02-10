package hdmodule

import (
	"encoding/json"
	"fmt"

	"github.com/nodeset-org/hyperdrive-ethereum/shared"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
	"github.com/urfave/cli/v2"
)

// Request format for `get-log-file`
type GetLogFileRequest struct {
	utils.KeyedRequest

	// The log file source to retrieve
	Source string `json:"source"`
}

// Response format for `get-log-file`
type getLogFileResponse struct {
	// The path to the log file
	Path string `json:"path"`
}

// Handle the `get-log-file` command
func getLogFile(c *cli.Context, handler utils.KeyedRequestHandler[*GetLogFileRequest]) error {
	// Get the request
	request, err := handler.HandleKeyedRequest(c)
	if err != nil {
		return fmt.Errorf("error reading set-settings request: %w", err)
	}

	// Get the path
	path := ""
	switch request.Source {
	case "adapter":
		path = utils.AdapterLogFile
	case shared.ServiceContainerName:
		path = shared.ServiceLogFile
	}

	// Create the response
	response := getLogFileResponse{
		Path: path,
	}

	// Marshal it
	bytes, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("error marshalling get-log-file response: %w", err)
	}

	// Print it
	fmt.Println(string(bytes))
	return nil
}
