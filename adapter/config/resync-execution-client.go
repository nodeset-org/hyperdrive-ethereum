package config

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
	"github.com/urfave/cli/v2"
)

// Destroy and resync the Execution client from scratch
func resyncExecutionClient(c *cli.Context) error {
	cfgMgr, err := NewAdapterConfigManager(c)
	if err != nil {
		return fmt.Errorf("error creating config manager: %w", err)
	}
	cfg, err := cfgMgr.LoadConfigFromDisk()
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}

	// Check the client mode
	if !cfg.IsLocalMode() {
		fmt.Println("You use an externally-managed Execution Client. Hyperdrive cannot resync it for you.")
		return nil
	}

	if !confirmPrompt(c, "Are you SURE you want to delete and resync your execution client from scratch? This cannot be undone!") {
		fmt.Println("Cancelled.")
		return nil
	}

	containerName := fmt.Sprintf("%s_ec", utils.ComposeProject)
	volumeName := fmt.Sprintf("%sdata", containerName)

	fmt.Printf("Preparing to resync Execution Client: %s\n", containerName)
	fmt.Printf("This will DELETE volume %s and force a full resync.\n\n", volumeName)

	// 1. Stop the Execution Client container
	fmt.Printf("Stopping container %s...\n", containerName)
	if err := runDockerCommand("stop", containerName); err != nil {
		fmt.Printf("Warning: stop failed or container not running (%v)\n", err)
	}

	// 2. Remove the container
	fmt.Printf("Removing container %s...\n", containerName)
	if err := runDockerCommand("rm", containerName); err != nil {
		fmt.Printf("Warning: container may not exist (%v)\n", err)
	}

	// 3. Remove the volume
	fmt.Printf("Deleting volume %s...\n", volumeName)
	if err := runDockerCommand("volume", "rm", volumeName); err != nil {
		fmt.Printf("Warning: volume may not exist (%v)\n", err)
	}

	// 4. Restart services
	fmt.Println("Restarting services...")

	if utils.ComposeDir == "" {
		return fmt.Errorf("%s not set", utils.ComposeDirEnvVarName)
	}
	if utils.ComposeProject == "" {
		return fmt.Errorf("%s not set", utils.ComposeProjectEnvVarName)
	}

	executionClientFile := filepath.Join(utils.ComposeDir, "ec.yml")

	for _, file := range []string{executionClientFile} {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return fmt.Errorf("required compose file missing: %s", file)
		}
	}

	args := []string{
		"compose",
		"-p", utils.ComposeProject,
		"-f", executionClientFile,
		"up",
		"-d",
		"--quiet-pull",
	}

	cmd := exec.Command("docker", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error starting services: %w", err)
	}

	fmt.Printf("\nDone! Your Execution Client is now resyncing. You can follow its progress with `hyperdrive service logs el`.\n")
	return nil
}
