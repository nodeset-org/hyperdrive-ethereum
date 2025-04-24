package hdmodule

import (
	"fmt"

	adapterApp "github.com/nodeset-org/hyperdrive-ethereum/adapter/app"
	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config"
	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"
	"github.com/urfave/cli/v2"
)

// Handles `hd-module` commands
func RegisterCommands(app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:    "hd-module",
		Aliases: []string{"hd"},
		Usage:   "Handle Hyperdrive module commands",
		Subcommands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Flags:   []cli.Flag{},
				Usage:   "Print the module version.",
				Action: func(c *cli.Context) error {
					// Validate args
					utils.ValidateArgCount(c, 0)

					// Run
					return version()
				},
			},
			{
				Name:    "get-config-metadata",
				Aliases: []string{"c"},
				Flags:   []cli.Flag{},
				Usage:   "Get the metadata for the module's configuration, representing how to present the parameters to the user.",
				Action: func(c *cli.Context) error {
					// Validate args
					utils.ValidateArgCount(c, 0)

					// Run
					return getConfigMetadata(c)
				},
			},
			{
				Name:    "process-settings",
				Aliases: []string{"p"},
				Flags:   []cli.Flag{},
				Usage:   "Process the module's configuration, validating it without saving.",
				Action: func(c *cli.Context) error {
					// Validate args
					utils.ValidateArgCount(c, 0)

					// Run
					return processSettings(c)
				},
			},
			{
				Name:    "set-settings",
				Aliases: []string{"ss"},
				Flags:   []cli.Flag{},
				Usage:   "Sets the module's configuration, saving it to disk.",
				Action: func(c *cli.Context) error {
					// Validate args
					utils.ValidateArgCount(c, 0)
					handler := utils.DefaultKeyedRequestHandler[*SetSettingsRequest]{}
					cfgMgr, err := config.NewAdapterConfigManager(c)

					if err != nil {
						return fmt.Errorf("error creating config manager: %w", err)
					}

					// Run
					return setSettings(c, handler, cfgMgr)
				},
			},
			{
				Name:    "start",
				Aliases: []string{"s"},
				Flags:   []cli.Flag{},
				Usage:   "Start the module's services (stopping and restarting any that changed due to the new configuration).",
				Action: func(c *cli.Context) error {
					// Validate args
					utils.ValidateArgCount(c, 0)

					// Run
					return startServices(c)
				},
			},
			{
				Name:    "run",
				Aliases: []string{"r"},
				Flags:   []cli.Flag{},
				Usage:   "Run a command.",
				Action: func(c *cli.Context) error {
					// Validate args
					utils.ValidateArgCount(c, 0)
					handler := utils.DefaultKeyedRequestHandler[*RunRequest]{}
					// Run
					app := adapterApp.CreateApp()

					return run(c, app, handler)
				},
			},
			{
				Name:    "call-config-function",
				Aliases: []string{"ccf"},
				Flags:   []cli.Flag{},
				Usage:   "Call a function in the module's configuration.",
				Action: func(c *cli.Context) error {
					// Validate args
					utils.ValidateArgCount(c, 0)

					// Run
					return callConfigFunction(c)
				},
			},
		},
	})
}
