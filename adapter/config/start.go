package config

import (
	"fmt"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config/utils/terminal"
	"github.com/urfave/cli/v2"
)

// Start the Hyperdrive service
func startService(c *cli.Context, ignoreConfigSuggestion bool) error {
	// Get Hyperdrive client
	hd, err := client.NewHyperdriveClientFromCtx(c)
	if err != nil {
		return err
	}

	// Load configuration
	cfg, isNew, err := hd.LoadConfig()
	if err != nil {
		return fmt.Errorf("Error loading user settings: %w", err)
	}
	if isNew {
		return fmt.Errorf("No configuration detected. Please run `hyperdrive service config` to set up Hyperdrive before running it.")
	}

	// Validate config
	errors := cfg.Validate()
	if len(errors) > 0 {
		fmt.Printf("%sYour configuration encountered errors. You must correct the following by changing the settings with `hyperdrive service config` in order to start Hyperdrive:\n\n", terminal.ColorRed)
		for _, err := range errors {
			fmt.Printf("%s\n\n", err)
		}
		fmt.Println(terminal.ColorReset)
		return nil
	}

	// Perform anti-slashing safety check
	if !c.Bool(ignoreSlashTimerFlag.Name) {
		firstRun, err := checkForValidatorChange(hd, cfg)
		if err != nil {
			fmt.Printf("%sWARNING: couldn't verify validator client restart safety:\n\t%s\n", terminal.ColorYellow, err.Error())
			fmt.Println("If changing clients, wait 15 minutes before starting.")
			if !cliutils.Confirm(fmt.Sprintf("Press y to acknowledge and start Hyperdrive:%s", terminal.ColorReset)) {
				fmt.Println("Cancelled.")
				return nil
			}
		} else if firstRun {
			if !cliutils.Confirm(fmt.Sprintf("Press y to acknowledge and start Hyperdrive:%s", terminal.ColorReset)) {
				fmt.Println("Cancelled.")
				return nil
			}
		}
	}

	// Start only the Beacon Node and Execution Client
	fmt.Println("Starting Beacon Node and Execution Client...")
	err = hd.StartService([]string{
		"beacon-node",
		"execution-client",
	})
	if err != nil {
		return fmt.Errorf("error starting beacon and execution client: %w", err)
	}

	fmt.Println("Beacon Node and Execution Client started successfully.")
	return nil
}
