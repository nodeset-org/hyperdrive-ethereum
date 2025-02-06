package adapter_test

import (
	"context"
	"testing"

	"github.com/nodeset-org/hyperdrive-ethereum/shared"
	"github.com/stretchr/testify/require"
)

func TestGetVersion(t *testing.T) {
	// version, err := gac.GetVersion(context.Background())
	// if err != nil {
	// 	t.Errorf("error getting version: %v", err)
	// }
	// t.Logf("Adapter version: %s", version)
	// require.Equal(t, "0.2.0", version)
	// Run the command inside the module container
	ctx := context.Background()

	// cmd := []string{"hd-module", "version"}
	var response struct {
		Version string `json:"version"`
	}

	err := pac.RunCommand(ctx, "hd-module version", nil, &response)

	// Ensure the command executed successfully
	require.NoError(t, err, "error running hd-module version")

	// Validate that the version is correct
	require.Equal(t, shared.HyperdriveEthereumVersion, response.Version, "incorrect version returned")
}
