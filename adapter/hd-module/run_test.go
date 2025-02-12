package hdmodule

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

type MockApp struct {
	ReceivedArgs []string
}

func (m *MockApp) Run(args []string) error {
	m.ReceivedArgs = args
	return nil // Simulate successful execution
}

func TestRun_Success(t *testing.T) {
	// Create a mock app to capture the args
	mockApp := &cli.App{
		Action: func(c *cli.Context) error {
			// Simulate successful execution
			return nil
		},
	}

	// Setup CLI context
	appInstance := cli.NewApp()
	set := flag.NewFlagSet("test", 0)

	ctx := cli.NewContext(appInstance, set, nil)

	// Run the function with the mock app instance
	err := run(ctx, mockApp)

	// Assertions
	assert.NoError(t, err, "Expected no error during execution")
}
