package hdmodule

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
	hdconfig "github.com/nodeset-org/hyperdrive/config"
	"github.com/urfave/cli/v2"
)

// Request format for `start`
type startRequest struct {
	utils.KeyedRequest

	// The config settings being used
	Settings *hdconfig.HyperdriveSettings `json:"settings"`
}

// Handle the `start` command
func startServices(c *cli.Context) error {
	// Ensure environment variables are set
	if utils.ComposeDir == "" {
		return fmt.Errorf("%s not set", utils.ComposeDirEnvVarName)
	}
	if utils.ComposeProject == "" {
		return fmt.Errorf("%s not set", utils.ComposeProjectEnvVarName)
	}

	// Handle request
	_, err := utils.HandleKeyedRequest[*startRequest](c)
	if err != nil {
		return err
	}

	// Correctly reference the actual compose files
	beaconNodeFile := filepath.Join(utils.ComposeDir, "beacon-node.yml")
	executionClientFile := filepath.Join(utils.ComposeDir, "execution-client.yml")

	// Verify that the files exist before running the command
	for _, file := range []string{beaconNodeFile, executionClientFile} {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return fmt.Errorf("required compose file missing: %s", file)
		}
	}

	// Start the services with both YAML files
	args := []string{
		"compose",
		"-p",
		utils.ComposeProject,
		"-f", beaconNodeFile,
		"-f", executionClientFile,
		"up",
		"-d",
		"--quiet-pull",
	}

	cmd := exec.Command("docker", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error starting services: %w", err)
	}

	return nil
}
