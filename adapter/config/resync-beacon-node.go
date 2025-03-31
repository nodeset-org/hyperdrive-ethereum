package config

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	hdconfig "github.com/nodeset-org/hyperdrive/config"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
	"github.com/urfave/cli/v2"
)

type ResyncBeaconNodeRequest struct {
	utils.KeyedRequest

	Settings *hdconfig.HyperdriveSettings `json:"settings"`
}

// TODO (HN)
// Destroy and resync the Beacon Node from scratch
func resyncBeaconNode(
	c *cli.Context,
	handler utils.KeyedRequestHandler[*ResyncBeaconNodeRequest],

) error {
	request, err := handler.HandleKeyedRequest(c)
	if err != nil {
		return fmt.Errorf("error reading set-settings request: %w", err)
	}

	modInstance, exists := request.Settings.Modules[utils.FullyQualifiedModuleName]
	if !exists {
		return fmt.Errorf("could not find config for %s", utils.FullyQualifiedModuleName)
	}

	var settings HyperdriveEthereumConfigSettings
	err = modInstance.DeserializeSettingsIntoKnownType(&settings)
	if err != nil {
		return fmt.Errorf("error loading settings: %w", err)
	}
	containerName := fmt.Sprintf("%s_bn", utils.ComposeProject)
	volumeName := "bndata"

	fmt.Printf("Preparing to resync Beacon Node: %s\n", containerName)
	fmt.Printf("This will DELETE volume %s and force a full resync.\n\n", volumeName)

	// 1. Stop the Beacon Node container
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

	beaconNodeFile := filepath.Join(utils.ComposeDir, "bn.yml")
	executionClientFile := filepath.Join(utils.ComposeDir, "ec.yml")

	for _, file := range []string{beaconNodeFile, executionClientFile} {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return fmt.Errorf("required compose file missing: %s", file)
		}
	}

	args := []string{
		"compose",
		"-p", utils.ComposeProject,
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
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error starting services: %w", err)
	}

	fmt.Printf("\nDone! Your Beacon Node is now resyncing. You can follow its progress with `hyperdrive service logs bn`.\n")
	return nil
}

func runDockerCommand(args ...string) error {
	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// func confirmPrompt(message string) bool {
// 	fmt.Printf("%s [y/N]: ", message)
// 	reader := bufio.NewReader(os.Stdin)
// 	text, _ := reader.ReadString('\n')
// 	text = strings.ToLower(strings.TrimSpace(text))
// 	return text == "y" || text == "yes"
// }

// // Get the merged config
// cfg, isNew, err := hd.LoadConfig()
// if err != nil {
// 	return err
// }
// if isNew {
// 	return fmt.Errorf("Settings file not found. Please run `hyperdrive service config` to set up Hyperdrive.")
// }

// // Stop the BN
// beaconContainerName := cfg.Hyperdrive.GetDockerArtifactName(string(ContainerID_BeaconNode))
// fmt.Printf("Stopping %s...\n", beaconContainerName)
// err = hd.StopContainer(beaconContainerName)
// if err != nil {
// 	fmt.Printf("%sWARNING: Stopping Beacon Node container failed: %s%s\n", terminal.ColorYellow, err.Error(), terminal.ColorReset)
// }

// // Get the BN volume name
// volume, err := hd.GetClientVolumeName(beaconContainerName, clientDataVolumeName)
// if err != nil {
// 	return fmt.Errorf("Error getting Beacon Node volume name: %w", err)
// }

// // Remove the BN
// fmt.Printf("Deleting %s...\n", beaconContainerName)
// err = hd.RemoveContainer(beaconContainerName)
// if err != nil {
// 	return fmt.Errorf("Error deleting Beacon Node container: %w", err)
// }

// // Delete the BN volume
// fmt.Printf("Deleting volume %s...\n", volume)
// err = hd.DeleteVolume(volume)
// if err != nil {
// 	return fmt.Errorf("Error deleting volume: %w", err)
// }

// // TODO (HN): WIP
// // Create the configuration manager
// cfgMgr, err := NewAdapterConfigManager(c)
// if err != nil {
// 	return fmt.Errorf("error creating config manager: %w", err)
// }
// cfg, err := cfgMgr.LoadConfigFromDisk()
// if err != nil {
// 	return fmt.Errorf("error loading config: %w", err)
// }
// if cfg == nil {
// 	return fmt.Errorf("config has not been created yet")
// }

// // Get the current checkpoint sync URL
// checkpointSyncUrl := string(cfg.LocalBeaconClient.CheckpointSyncProvider)
// if checkpointSyncUrl == "" {
// 	fmt.Printf("%sYou do not have a checkpoint sync provider configured.\nIf you have active validators, they %swill be considered offline and will lose ETH%s%s until your Beacon Node finishes syncing.\nWe strongly recommend you configure a checkpoint sync provider with `hyperdrive service config` so it syncs instantly before running this.%s\n\n", terminal.ColorRed, terminal.ColorBold, terminal.ColorReset, terminal.ColorRed, terminal.ColorReset)
// } else {
// 	fmt.Printf("You have a checkpoint sync provider configured (%s).\nYour Beacon Node will use it to sync to the head of the Beacon Chain instantly after being rebuilt.\n\n", checkpointSyncUrl)
// }

// // Prompt for confirmation
// if !(c.Bool(adapterutils.YesFlag.Name) || utils.Confirm(fmt.Sprintf("%sAre you SURE you want to delete and resync your main Beacon Node from scratch? This cannot be undone!%s", terminal.ColorRed, terminal.ColorReset))) {
// 	fmt.Println("Cancelled.")
// 	return nil
// }

// // Restart containers
// fmt.Printf("Rebuilding %s and restarting Hyperdrive...\n", beaconContainerName)
// err = startService(c, true)
// if err != nil {
// 	return fmt.Errorf("Error starting Hyperdrive: %s", err)
// }

// fmt.Printf("\nDone! Your Beacon Node is now resyncing. You can follow its progress with `hyperdrive service logs bn`.\n")
// return nil
