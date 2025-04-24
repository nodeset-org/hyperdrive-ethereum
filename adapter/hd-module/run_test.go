package hdmodule

import (
	"flag"
	"fmt"
	"testing"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

type BrokenMarshalStruct struct{}

func (b BrokenMarshalStruct) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("mock JSON marshalling error")
}

type MockKeyedRequestHandler[RequestType utils.IKeyedRequest] struct {
	ReturnRequest RequestType
	ReturnError   error
}

func (m MockKeyedRequestHandler[RequestType]) HandleKeyedRequest(c *cli.Context) (RequestType, error) {
	return m.ReturnRequest, m.ReturnError
}

type MockApp struct {
	ReceivedArgs []string
}

func (m *MockApp) Run(args []string) error {
	m.ReceivedArgs = args
	return nil
}

func TestRun_Success(t *testing.T) {
	mockHandler := MockKeyedRequestHandler[*RunRequest]{
		ReturnRequest: &RunRequest{Command: "echo test"},
		ReturnError:   nil,
	}

	mockApp := &cli.App{
		Action: func(c *cli.Context) error {

			return nil
		},
	}

	appInstance := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(appInstance, set, nil)

	err := run(ctx, mockApp, mockHandler)

	assert.NoError(t, err, "Expected no error during execution")
}

func TestRun_CommandParsingError(t *testing.T) {
	mockHandler := MockKeyedRequestHandler[*RunRequest]{
		ReturnRequest: &RunRequest{Command: `"unterminated quote`},
		ReturnError:   nil,
	}

	mockApp := &cli.App{}

	appInstance := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(appInstance, set, nil)

	err := run(ctx, mockApp, mockHandler)

	assert.Error(t, err, "Expected an error due to command parsing failure")
	assert.Contains(t, err.Error(), "error parsing command", "Error should mention parsing issue")
}

func TestRun_RecursiveCallError(t *testing.T) {
	mockHandler := MockKeyedRequestHandler[*RunRequest]{
		ReturnRequest: &RunRequest{Command: "hd-module process-settings"},
		ReturnError:   nil,
	}

	mockApp := &cli.App{}

	appInstance := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(appInstance, set, nil)

	err := run(ctx, mockApp, mockHandler)

	assert.Error(t, err, "Expected an error due to recursive call")
	assert.Contains(t, err.Error(), "recursive calls to `run` are not allowed", "Error should indicate recursion prevention")
}

func TestRun_RequestHandlingError(t *testing.T) {
	mockHandler := MockKeyedRequestHandler[*RunRequest]{
		ReturnRequest: nil,
		ReturnError:   fmt.Errorf("mock request error"),
	}

	mockApp := &cli.App{}

	appInstance := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(appInstance, set, nil)

	err := run(ctx, mockApp, mockHandler)

	assert.Error(t, err, "Expected an error when request handling fails")
	assert.Contains(t, err.Error(), "mock request error", "Error should match the mock request error")
}

func TestRun_EmptyCommandError(t *testing.T) {
	mockHandler := MockKeyedRequestHandler[*RunRequest]{
		ReturnRequest: &RunRequest{Command: ""},
		ReturnError:   nil,
	}

	mockApp := &cli.App{}

	appInstance := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(appInstance, set, nil)

	err := run(ctx, mockApp, mockHandler)

	assert.Error(t, err, "Expected an error due to empty command")
	assert.Contains(t, err.Error(), "error parsing command", "Error should mention missing command issue")
}

func TestRun_CommandExecutionWithArgs(t *testing.T) {
	mockHandler := MockKeyedRequestHandler[*RunRequest]{
		ReturnRequest: &RunRequest{Command: "ls -la"},
		ReturnError:   nil,
	}

	mockApp := &cli.App{
		Action: func(c *cli.Context) error {
			expectedArgs := []string{"ls", "-la"}
			actualArgs := c.Args().Slice()
			assert.Equal(t, expectedArgs, actualArgs, "Unexpected command arguments")
			return nil
		},
	}

	appInstance := cli.NewApp()
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(appInstance, set, nil)

	err := run(ctx, mockApp, mockHandler)

	assert.NoError(t, err, "Expected successful execution with arguments")
}
