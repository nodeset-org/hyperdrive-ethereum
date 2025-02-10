package app

import (
	"fmt"
	"os"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config"
	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"

	"github.com/nodeset-org/hyperdrive-ethereum/shared"
	"github.com/urfave/cli/v2"
)

func CreateApp() *cli.App {
	// Initialize application
	app := cli.NewApp()

	// Set application info
	app.Name = "Hyperdrive Ethereum Adapter"
	app.Usage = "Adapter for the Hyperdrive Ethereum module"
	app.Version = shared.HyperdriveEthereumVersion
	app.Authors = []*cli.Author{
		{
			Name:  "Nodeset",
			Email: "info@nodeset.io",
		},
	}
	app.Copyright = "(c) 2025 NodeSet LLC"

	// Enable Bash Completion
	app.EnableBashCompletion = true

	// Set application flags
	app.Flags = []cli.Flag{
		utils.ConfigDirFlag,
		utils.LogDirFlag,
		utils.KeyFileFlag,
	}

	// Register commands
	config.RegisterCommands(app)

	app.Before = func(c *cli.Context) error {
		if utils.Mode == utils.AdapterMode_Project {
			// Make the authenticator
			auth, err := utils.NewAuthenticator(c)
			if err != nil {
				return err
			}
			c.App.Metadata[utils.AuthenticatorMetadataKey] = auth
		}

		return nil
	}
	app.BashComplete = func(c *cli.Context) {
		// Load the context and flags prior to autocomplete
		err := app.Before(c)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}

		// Run the default autocomplete
		cli.DefaultAppComplete(c)
	}

	return app
}
