package config

import (
	hdconfig "github.com/nodeset-org/hyperdrive-daemon/shared/config"
)

// A merged set of general resources and StakeWise-specific resources for the selected network
type MergedResources struct {
	// General resources
	*hdconfig.MergedResources
}
