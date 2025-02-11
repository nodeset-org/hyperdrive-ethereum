package hdmodule

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/nodeset-org/hyperdrive-ethereum/shared"
)

// Response format for `version`
type versionResponse struct {
	// Version of the module
	Version string `json:"version"`
}

func version() error {
	version := versionResponse{
		Version: shared.HyperdriveEthereumVersion,
	}

	bytes, err := json.Marshal(version)
	if err != nil {
		return fmt.Errorf("error marshalling version response: %w", err)
	}

	fmt.Println(string(bytes))
	return nil
}
