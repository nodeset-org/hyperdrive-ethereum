package config

import (
	"fmt"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config/utils/terminal"
	"github.com/nodeset-org/hyperdrive-ethereum/adapter/utils"

	"github.com/urfave/cli/v2"
)

var (
	ignoreSlashTimerFlag *cli.BoolFlag = &cli.BoolFlag{
		Name:  "ignore-slash-timer",
		Usage: fmt.Sprintf("Bypass the safety timer that forces a delay when switching to a new Beacon Node.\n%sUsing this flag to bypass the slashing timer could result in a *major* loss of ETH! Only use this is if you absolutely understand the risks!%s", terminal.ColorRed, terminal.ColorReset),
	}
	// tailFlag *cli.StringFlag = &cli.StringFlag{
	// 	Name:    "tail",
	// 	Aliases: []string{"t"},
	// 	Usage:   "The number of lines to show from the end of the logs (number or \"all\")",
	// 	Value:   "100",
	// }
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
					return resyncExecutionClient(c)
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
					return resyncBeaconNode(c)
				},
			},
			// {
			// 	Name:    "start",
			// 	Aliases: []string{"s"},
			// 	Usage:   "Start the Hyperdrive service",
			// 	Flags: []cli.Flag{
			// 		ignoreSlashTimerFlag,
			// 		// nodeset.RegisterEmailFlag,
			// 		// wallet.PasswordFlag,
			// 		// wallet.SavePasswordFlag,
			// 		// utils.YesFlag,
			// 	},
			// 	Action: func(c *cli.Context) error {
			// 		// Validate args
			// 		utils.ValidateArgCount(c, 0)

			// 		// Run command
			// 		return startService(c, false)
			// 	},
			// },
		},
	})
}
