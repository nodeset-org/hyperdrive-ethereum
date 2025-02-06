package config

import (
	"github.com/urfave/cli/v2"
)

// TODO (HN)
// Destroy and resync the Beacon Node from scratch
func resyncBeaconNode(c *cli.Context) error {
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

	// TODO (HN): WIP
	// Create the configuration manager
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

	// // Restart Hyperdrive
	// fmt.Printf("Rebuilding %s and restarting Hyperdrive...\n", beaconContainerName)
	// err = startService(c, true)
	// if err != nil {
	// 	return fmt.Errorf("Error starting Hyperdrive: %s", err)
	// }

	// fmt.Printf("\nDone! Your Beacon Node is now resyncing. You can follow its progress with `hyperdrive service logs bn`.\n")
	return nil
}
