package testing

import (
	hdconfig "github.com/nodeset-org/hyperdrive-daemon/shared/config"
	heconfig "github.com/nodeset-org/hyperdrive-ethereum/shared/config"
	"github.com/rocket-pool/node-manager-core/config"
)

// Returns a new StakewiseResources instance with test network values
func getTestResourcesgetTestResources(hdResources *hdconfig.MergedResources, deploymentName string) *heconfig.MergedResources {
	return &heconfig.MergedResources{
		MergedResources: hdResources,
	}
}

// Provisions a NetworkSettings instance with updated addresses
func provisionNetworkSettings(networkSettings *config.NetworkSettings) *config.NetworkSettings {
	return networkSettings
}
