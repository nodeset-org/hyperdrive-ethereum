package testing

import (
	"fmt"
	"os"
	"path/filepath"

	hdservices "github.com/nodeset-org/hyperdrive-daemon/module-utils/services"
	hdconfig "github.com/nodeset-org/hyperdrive-daemon/shared/config"
	hdtesting "github.com/nodeset-org/hyperdrive-daemon/testing"
	hecommon "github.com/nodeset-org/hyperdrive-ethereum/common"
	heconfig "github.com/nodeset-org/hyperdrive-ethereum/shared/config"
	"github.com/rocket-pool/node-manager-core/log"
)

const (
	deploymentName string = "localtest"
)

// HyperdriveEthereumTestManager for managing testing resources and services
type HyperdriveEthereumTestManager struct {
	*hdtesting.HyperdriveTestManager

	// The complete Hyperdrive Ethereum node
	node *HyperdriveEthereumNode

	// The StakeWise operator mock
	// operatorMock *OperatorMock

	// The snapshot ID of the baseline snapshot
	baselineSnapshotID string
}

// Creates a new TestManager instance
func NewHyperdriveEthereumTestManager() (*HyperdriveEthereumTestManager, error) {
	tm, err := hdtesting.NewHyperdriveTestManagerWithDefaults(provisionNetworkSettings)
	if err != nil {
		return nil, fmt.Errorf("error creating test manager: %w", err)
	}

	// Get the HD artifacts
	hdNode := tm.GetNode()
	hdSp := hdNode.GetServiceProvider()
	hdCfg := hdSp.GetConfig()
	hdClient := hdNode.GetApiClient()

	// Make HyperdriveEthereum resources
	resources := getTestResources(hdSp.GetResources(), deploymentName)
	swCfg, err := heconfig.NewHyperdriveEthereumConfig(hdCfg, []*heconfig.HyperdriveEthereumSettings{})
	if err != nil {
		closeTestManager(tm)
		return nil, fmt.Errorf("error creating HyperdriveEthereum config: %v", err)
	}

	// Make the module directory
	moduleDir := filepath.Join(hdCfg.UserDataPath.Value, hdconfig.ModulesName, heconfig.ModuleName)
	err = os.MkdirAll(moduleDir, 0755)
	if err != nil {
		closeTestManager(tm)
		return nil, fmt.Errorf("error creating module directory [%s]: %v", moduleDir, err)
	}

	// Make a new service provider
	moduleSp, err := hdservices.NewModuleServiceProviderFromArtifacts(hdClient, hdCfg, swCfg, hdSp.GetResources(), moduleDir, heconfig.ModuleName, heconfig.ClientLogName, hdSp.GetEthClient(), hdSp.GetBeaconClient())
	if err != nil {
		closeTestManager(tm)
		return nil, fmt.Errorf("error creating service provider: %v", err)
	}
	heSP, err := hecommon.NewHyperdriveEthereumerviceProviderFromCustomServices(moduleSp, swCfg, resources)
	if err != nil {
		closeTestManager(tm)
		return nil, fmt.Errorf("error creating HyperdriveEthereum service provider: %v", err)
	}

	// Create the HyperdriveEthereum node
	nodeIP := "localhost"
	node, err := newHyperdriveEthereumNode(heSP, nodeIP, tm.GetLogger(), hdNode)
	if err != nil {
		closeTestManager(tm)
		return nil, fmt.Errorf("error creating HyperdriveEthereum node: %v", err)
	}

	// Disable automining
	err = tm.ToggleAutoMine(false)
	if err != nil {
		closeTestManager(tm)
		return nil, fmt.Errorf("error disabling automining: %v", err)
	}

	// Return
	module := &HyperdriveEthereumTestManager{
		HyperdriveTestManager: tm,
		node:                  node,
	}
	tm.RegisterModule(module)

	baselineSnapshot, err := tm.CreateSnapshot()
	if err != nil {
		return nil, fmt.Errorf("error creating baseline snapshot: %w", err)
	}
	module.baselineSnapshotID = baselineSnapshot
	return module, nil
}

// ===============
// === Getters ===
// ===============

func (m *HyperdriveEthereumTestManager) GetModuleName() string {
	return "hyperdrive-ethereum"
}

// Get the node handle
func (m *HyperdriveEthereumTestManager) GetNode() *HyperdriveEthereumNode {
	return m.node
}

// ====================
// === Snapshotting ===
// ====================

// Reverts the service states to the baseline snapshot
func (m *HyperdriveEthereumTestManager) DependsOnHyperdriveEthereumBaseline() error {
	err := m.RevertSnapshot(m.baselineSnapshotID)
	if err != nil {
		return fmt.Errorf("error reverting to baseline snapshot: %w", err)
	}
	return nil
}

// Takes a snapshot of the service states
func (m *HyperdriveEthereumTestManager) TakeModuleSnapshot() (any, error) {
	snapshotName, err := m.HyperdriveTestManager.TakeModuleSnapshot()
	if err != nil {
		return nil, fmt.Errorf("error taking snapshot: %w", err)
	}
	// The OperatorMock doesn't have any state so it doesn't need snapshotting
	return snapshotName, nil
}

func (m *HyperdriveEthereumTestManager) RevertModuleToSnapshot(moduleState any) error {
	err := m.HyperdriveTestManager.RevertModuleToSnapshot(moduleState)
	if err != nil {
		return fmt.Errorf("error reverting to snapshot: %w", err)
	}

	// Reload the SW wallet to undo any changes made during the test
	wallet := m.node.sp.GetWallet()
	err = wallet.Reload()
	if err != nil {
		return fmt.Errorf("error reloading HyperdriveEthereum wallet: %v", err)
	}

	// Reload the available key manager
	err = m.node.sp.GetAvailableKeyManager().Reload()
	if err != nil {
		return fmt.Errorf("error reloading available key manager: %v", err)
	}
	return nil
}

// Closes the test manager, shutting down the nodeset mock server and all other resources
func (m *HyperdriveEthereumTestManager) CloseModule() error {
	err := m.node.Close()
	if err != nil {
		return fmt.Errorf("error closing HyperdriveEthereum node: %w", err)
	}
	if m.HyperdriveTestManager != nil {
		m.HyperdriveTestManager = nil
	}
	return nil
}

// ==========================
// === Internal Functions ===
// ==========================

// Closes the Hyperdrive test manager, logging any errors
func closeTestManager(tm *hdtesting.HyperdriveTestManager) {
	err := tm.Close()
	if err != nil {
		tm.GetLogger().Error("Error closing test manager", log.Err(err))
	}
}
