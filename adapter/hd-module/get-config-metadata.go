package hdmodule

import (
	"encoding/json"
	"fmt"

	"github.com/nodeset-org/hyperdrive-ethereum/adapter/config"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
	"github.com/urfave/cli/v2"
)

func getConfigMetadata(c *cli.Context) error {
	// Get the config
	cfg := config.NewHyperdriveEthereumConfig()

	// Create the response
	cfgMap := hdconfig.MarshalConfigurationToMap(cfg)
	bytes, err := json.Marshal(cfgMap)
	if err != nil {
		return fmt.Errorf("error marshalling config: %w", err)
	}

	// Print it
	fmt.Println(string(bytes))
	return nil
}
