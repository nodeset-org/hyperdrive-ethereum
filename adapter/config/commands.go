package config

import (
	"github.com/urfave/cli/v2"
	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config/service"
)

func RegisterCommands(app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:    "services",
		Aliases: []string{"s"},
		Usage:   "Commands for manager Node manager services",
		Subcommands: []*cli.Command{
			{
				Name:    "resync-ec",
				Aliases: []string{"resync-eth1"},
				Usage:   fmt.Sprintf("%sDeletes the main Execution client's chain data and resyncs it from scratch. Only use this as a last resort!%s", terminal.ColorRed, terminal.ColorReset),
				Action: func(c *cli.Context) error {
					// Validate args
					utils.ValidateArgCount(c, 0)

					// Run command
					return service.ResyncExecutionClient(c)
				},
			},
			{
				Name:    "resync-bn",
				Aliases: []string{"resync-eth2"},
				Usage:   fmt.Sprintf("%sDeletes the Beacon Node's chain data and resyncs it from scratch. Only use this as a last resort!%s", terminal.ColorRed, terminal.ColorReset),
				Action: func(c *cli.Context) error {
					// Validate args
					utils.ValidateArgCount(c, 0)

					// Run command
					return service.ResyncBeaconNode(c)
				},
			},
		},
	}
}