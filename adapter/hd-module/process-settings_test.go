package hdmodule

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/nodeset-org/hyperdrive-ethereum/shared"
	hdconfig "github.com/nodeset-org/hyperdrive/config"
	"github.com/stretchr/testify/require"
)

const (
	oldSettingsJson string = `{
        "modules": {
            "NodeSet/hyperdrive-ethereum": {
                "enabled": true,
                "version": "0.2.0",
                "settings": {
                    "projectName": "hyperdrive-ethereum",
                    "apiPort": 8085,
                    "autoTxMaxFee": 100.5,
                    "maxPriorityFee": 2,
                    "autoTxGasThreshold": 1.5,
                    "network": "mainnet",
                    "clientMode": "local",
                    "localExecutionClient": {
                        "executionClient": "geth",
                        "httpPort": 8545,
                        "wsPort": 8546,
                        "enginePort": 8551,
                        "openApiPorts": "open",
                        "p2pPort": 30303
                    },
                    "externalExecutionClient": {
                        "endpoint": "https://external-execution-node.io",
                        "enginePort": 8552
                    },
                    "localBeaconClient": {
                        "checkpointSyncProvider": "https://beacon-node.io"
                    },
                    "externalBeaconClient": {
                        "endpoint": "https://external-beacon-node.io"
                    },
                    "fallback": {
                        "enabled": true,
                        "retryDelay": 5
                    },
                    "server": {
                        "port": 8085,
                        "portMode": "open"
                    },
                    "containerTag": "v0.2.0",
                    "version": "0.2.0",
                    "isNew": false
                }
            }
        }
    }`

	newSettingsJson string = `{
        "modules": {
            "NodeSet/hyperdrive-ethereum": {
                "enabled": true,
                "version": "0.3.0",
                "settings": {
                    "projectName": "hyperdrive-ethereum",
                    "apiPort": 8085,
                    "autoTxMaxFee": 100.5,
                    "maxPriorityFee": 2,
                    "autoTxGasThreshold": 1.5,
                    "network": "mainnet",
                    "clientMode": "local",
                    "localExecutionClient": {
                        "executionClient": "geth",
                        "httpPort": 8545,
                        "wsPort": 8546,
                        "enginePort": 8551,
                        "openApiPorts": "open",
                        "p2pPort": 30303
                    },
                    "externalExecutionClient": {
                        "endpoint": "https://external-execution-node.io",
                        "enginePort": 8552
                    },
                    "localBeaconClient": {
                        "checkpointSyncProvider": "https://beacon-node.io"
                    },
                    "externalBeaconClient": {
                        "endpoint": "https://external-beacon-node.io"
                    },
                    "fallback": {
                        "enabled": true,
                        "retryDelay": 5
                    },
                    "server": {
                        "port": 8085,
                        "portMode": "open"
                    },
                    "containerTag": "v0.3.0",
                    "version": "0.3.0",
                    "isNew": true
                }
            }
        }
    }`
)

func TestProcessSettings(t *testing.T) {
	oldSettings := new(hdconfig.HyperdriveSettings)
	err := json.Unmarshal([]byte(oldSettingsJson), oldSettings)
	require.NoError(t, err)

	newSettings := new(hdconfig.HyperdriveSettings)
	err = json.Unmarshal([]byte(newSettingsJson), newSettings)
	require.NoError(t, err)

	// Process the settings
	response, err := processSettingsImpl(oldSettings, newSettings)
	require.NoError(t, err)

	// Check the response
	require.Len(t, response.Errors, 0)
	require.Len(t, response.Ports, 1)
	require.Equal(t, uint16(8085), response.Ports["server/port"])
	require.Len(t, response.ServicesToRestart, 1)
	require.Equal(t, shared.ServiceContainerName, response.ServicesToRestart[0])
	t.Log("Settings processed properly")
}
