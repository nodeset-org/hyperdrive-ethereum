package testing

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"

	hdservices "github.com/nodeset-org/hyperdrive-daemon/module-utils/services"
	hdconfig "github.com/nodeset-org/hyperdrive-daemon/shared/config"
	hdtesting "github.com/nodeset-org/hyperdrive-daemon/testing"
	heclient "github.com/nodeset-org/hyperdrive-ethereum/client"
	hecommon "github.com/nodeset-org/hyperdrive-ethereum/common"
	heconfig "github.com/nodeset-org/hyperdrive-ethereum/shared/config"
)

const (
	apiAuthKey string = "sw-test-key"
)

// A complete Hyperdrive Ethereum node instance
type HyperdriveEthereumNode struct {
	// The daemon's service provider
	sp hecommon.IHyperdriveEthereumServiceProvider

	// An HTTP API client for the daemon
	client *heclient.ApiClient

	// The client logger
	logger *slog.Logger

	// The Hyperdrive node parent
	hdNode *hdtesting.HyperdriveNode

	// Wait groups for graceful shutdown
	apiWg   *sync.WaitGroup
	relayWg *sync.WaitGroup
}

// Create a new HyperdriveEthereum node, including its folder structure, service provider, server manager, and API client.
func newHyperdriveEthereumNode(sp hecommon.IHyperdriveEthereumServiceProvider, address string, clientLogger *slog.Logger, hyperdriveNode *hdtesting.HyperdriveNode) (*HyperdriveEthereumNode, error) {
	// // Create the server
	// apiWg := &sync.WaitGroup{}
	// cfg := sp.GetConfig()
	// serverAuthMgr := auth.NewAuthorizationManager("", "he-server", auth.DefaultRequestLifespan)
	// serverAuthMgr.SetKey([]byte(apiAuthKey))
	// serverMgr, err := swserver.NewServerManager(sp, address, cfg.ApiPort.Value, apiWg, serverAuthMgr)
	// if err != nil {
	// 	return nil, fmt.Errorf("error creating constellation server: %v", err)
	// }

	// // Create the relay
	// relayWg := &sync.WaitGroup{}
	// relayServer, err := relay.NewRelayServer(sp, address, cfg.RelayPort.Value)
	// if err != nil {
	// 	return nil, fmt.Errorf("error creating relay server: %v", err)
	// }
	// err = relayServer.Start(relayWg)
	// if err != nil {
	// 	return nil, fmt.Errorf("error starting relay server: %v", err)
	// }

	// // Create the client
	// urlString := fmt.Sprintf("http://%s:%d/%s", address, serverMgr.GetPort(), heconfig.ApiClientRoute)
	// url, err := url.Parse(urlString)
	// if err != nil {
	// 	return nil, fmt.Errorf("error parsing client URL [%s]: %v", urlString, err)
	// }
	// clientAuthMgr := auth.NewAuthorizationManager("", "sw-client", auth.DefaultRequestLifespan)
	// clientAuthMgr.SetKey([]byte(apiAuthKey))
	// apiClient := heclient.NewApiClient(url, clientLogger, nil, clientAuthMgr)

	return &HyperdriveEthereumNode{
		// sp:          sp,
		// serverMgr:   serverMgr,
		// relayServer: relayServer,
		// client:      apiClient,
		// logger:      clientLogger,
		// hdNode:      hyperdriveNode,
		// apiWg:       apiWg,
		// relayWg:     relayWg,
	}, nil
}

// Closes the HyperdriveEthereum node. The caller is responsible for stopping the Hyperdrive daemon owning this module.
func (n *HyperdriveEthereumNode) Close() error {
	// if n.serverMgr != nil {
	// 	n.serverMgr.Stop()
	// 	n.apiWg.Wait()
	// 	n.serverMgr = nil
	// 	n.logger.Info("Stopped HyperdriveEthereum daemon API server")
	// }
	// if n.relayServer != nil {
	// 	err := n.relayServer.Stop()
	// 	if err != nil {
	// 		n.logger.Warn("relay server didn't shutdown cleanly", "error", err.Error())
	// 	}
	// 	n.relayWg.Wait()
	// 	n.relayServer = nil
	// 	n.logger.Info("Stopped HyperdriveEthereum relay server")
	// }
	return n.hdNode.Close()
}

// Get the daemon's service provider
func (n *HyperdriveEthereumNode) GetServiceProvider() hecommon.IHyperdriveEthereumServiceProvider {
	return n.sp
}

// // Get the HTTP API server for the node's daemon
// func (n *HyperdriveEthereumNode) GetServerManager() *swserver.ServerManager {
// 	return n.serverMgr
// }

// // Get the HTTP Relay server for the node's daemon
// func (n *HyperdriveEthereumNode) GetRelayServer() *relay.RelayServer {
// 	return n.relayServer
// }

// Get the HTTP API client for interacting with the node's daemon server
func (n *HyperdriveEthereumNode) GetApiClient() *heclient.ApiClient {
	return n.client
}

// Get the Hyperdrive node for this HyperdriveEthereum module
func (n *HyperdriveEthereumNode) GetHyperdriveNode() *hdtesting.HyperdriveNode {
	return n.hdNode
}

// Create a new HyperdriveEthereum node based on this one's configuration, but with a custom folder, address, and port.
func (n *HyperdriveEthereumNode) CreateSubNode(hdNode *hdtesting.HyperdriveNode, address string, port uint16) (*HyperdriveEthereumNode, error) {
	// Get the HD artifacts
	hdSp := hdNode.GetServiceProvider()
	hdCfg := hdSp.GetConfig()
	hdClient := hdNode.GetApiClient()

	// Make Constellation resources
	resources := getTestResources(hdSp.GetResources(), deploymentName)
	csCfg, err := heconfig.NewHyperdriveEthereumConfig(hdCfg, []*heconfig.HyperdriveEthereumSettings{})
	if err != nil {
		return nil, fmt.Errorf("error creating Constellation config: %v", err)
	}
	csCfg.ApiPort.Value = port

	// Make sure the module directory exists
	moduleDir := filepath.Join(hdCfg.UserDataPath.Value, hdconfig.ModulesName, heconfig.ModuleName)
	err = os.MkdirAll(moduleDir, 0755)
	if err != nil {
		return nil, fmt.Errorf("error creating data and modules directories [%s]: %v", moduleDir, err)
	}

	// Make a new service provider
	moduleSp, err := hdservices.NewModuleServiceProviderFromArtifacts(
		hdClient,
		hdCfg,
		csCfg,
		hdSp.GetResources(),
		moduleDir,
		heconfig.ModuleName,
		heconfig.ClientLogName,
		hdSp.GetEthClient(),
		hdSp.GetBeaconClient(),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating service provider: %v", err)
	}
	csSp, err := hecommon.NewHyperdriveEthereumServiceProviderFromCustomServices(moduleSp, csCfg, resources)
	if err != nil {
		return nil, fmt.Errorf("error creating hyperdrive-ethereum service provider: %v", err)
	}
	return newHyperdriveEthereumNode(csSp, address, n.logger, hdNode)
}
