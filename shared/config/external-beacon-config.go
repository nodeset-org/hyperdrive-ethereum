package config

import (
	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
)

// Configuration for external Beacon Nodes
type ExternalBeaconConfig struct {
	hdconfig.SectionHeader

	// The selected BN
	BeaconNode hdconfig.ChoiceParameter[BeaconNode] //Parameter[BeaconNode]

	// The URL of the HTTP endpoint
	HttpUrl hdconfig.StringParameter

	// The URL of the Prysm gRPC endpoint (only needed if using Prysm VCs)
	PrysmRpcUrl hdconfig.StringParameter
}

type ExternalBeaconConfigSettings struct {
	BeaconNode  BeaconNode `json:"beaconNode" yaml:"beaconNode"`
	HttpUrl     string     `json:"httpUrl" yaml:"httpUrl"`
	PrysmRpcUrl string     `json:"prysmRpcUrl" yaml:"prysmRpcUrl"`
}

// Generates a new ExternalBeaconConfig configuration
func NewExternalBeaconConfig() *ExternalBeaconConfig {
	cfg := &ExternalBeaconConfig{}
	cfg.ID = hdconfig.Identifier(ids.ExternalBnId)
	cfg.Name = "External BN"
	cfg.Description.Default = "External Beacon Node"

	cfg.HttpUrl.ID = hdconfig.Identifier(ids.HttpUrlID)
	cfg.HttpUrl.Name = "HTTP URL"
	cfg.HttpUrl.Description.Default = "The URL of the HTTP Beacon API endpoint for your external client.\nNOTE: If you are running it on the same machine as this node, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead, for example 'http://192.168.1.100:5052'."
	cfg.HttpUrl.Default = ""

	cfg.PrysmRpcUrl.ID = hdconfig.Identifier(ids.PrysmRpcUrlID)
	cfg.PrysmRpcUrl.Name = "Prysm RPC URL"
	cfg.PrysmRpcUrl.Description.Default = "The URL of Prysm's gRPC API endpoint for your external Beacon Node. Prysm's Validator Client will need this in order to connect to it.\nNOTE: If you are running it on the same machine as this node, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead, for example 'http://192.168.1.100:5053'."
	cfg.PrysmRpcUrl.Default = ""

	// Options for BeaconNode
	options := make([]hdconfig.ParameterOption[BeaconNode], 5)
	options[0].Name = string(BeaconNode_Lighthouse)
	options[0].Description.Default = "Select if your external client is Lighthouse."
	options[0].Value = BeaconNode_Lighthouse

	options[1].Name = string(BeaconNode_Lodestar)
	options[1].Description.Default = "Select if your external client is Lodestar."
	options[1].Value = BeaconNode_Lodestar

	options[2].Name = string(BeaconNode_Nimbus)
	options[2].Description.Default = "Select if your external client is Nimbus."
	options[2].Value = BeaconNode_Nimbus

	options[3].Name = string(BeaconNode_Prysm)
	options[3].Description.Default = "Select if your external client is Prysm."
	options[3].Value = BeaconNode_Prysm

	options[4].Name = string(BeaconNode_Teku)
	options[4].Description.Default = "Select if your external client is Teku."
	options[4].Value = BeaconNode_Teku

	cfg.BeaconNode.ID = hdconfig.Identifier(ids.BnID)
	cfg.BeaconNode.Name = "Beacon Node"
	cfg.BeaconNode.Description.Default = "Select which Beacon Node your external client is."
	cfg.BeaconNode.Options = options
	cfg.BeaconNode.Default = BeaconNode_Nimbus

	return cfg
}

// Get the parameters for this config
func (cfg *ExternalBeaconConfig) GetParameters() []hdconfig.IParameter {
	return []hdconfig.IParameter{
		&cfg.BeaconNode,
		&cfg.HttpUrl,
		&cfg.PrysmRpcUrl,
	}
}

func (cfg ExternalBeaconConfig) GetSections() []hdconfig.ISection {
	return []hdconfig.ISection{}
}

func (cfg *ExternalBeaconConfig) GetDescription() hdconfig.DynamicProperty[string] {
	return hdconfig.DynamicProperty[string]{}
}

func (cfg *ExternalBeaconConfig) GetDisabled() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}

func (cfg *ExternalBeaconConfig) GetHidden() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}
