package hdmodule

import (
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
					handler := utils.DefaultKeyedRequestHandler[*ProcessSettingsRequest]{}

					// Run
					return processSettings(c, handler)
				},
			},
			{
				Name:    "set-settings",
				Aliases: []string{"s"},
				Flags:   []cli.Flag{},
				Usage:   "Sets the module's configuration, saving it to disk.",
				Action: func(c *cli.Context) error {
					// Validate args
					utils.ValidateArgCount(c, 0)
					handler := utils.DefaultKeyedRequestHandler[*setSettingsRequest]{}

					// Run
					return setSettings(c, handler, config.NewAdapterConfigManager)
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
		},
	})
}
