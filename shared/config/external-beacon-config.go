package config

import (
	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
)

// Configuration for external Beacon Nodes
type ExternalBeaconConfig struct {
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

	cfg.HttpUrl.ID = hdconfig.Identifier(ids.HttpUrlID)
	cfg.HttpUrl.Name = "HTTP URL"
	cfg.HttpUrl.Description.Default = "The URL of the HTTP Beacon API endpoint for your external client.\nNOTE: If you are running it on the same machine as this node, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead, for example 'http://192.168.1.100:5052'."
	// cfg.HttpUrl.AffectedContainers = []string{string(ContainerID_Daemon), string(ContainerID_ValidatorClient)}

	cfg.PrysmRpcUrl.ID = hdconfig.Identifier(ids.PrysmRpcUrlID)
	cfg.PrysmRpcUrl.Name = "Prysm RPC URL"
	cfg.PrysmRpcUrl.Description.Default = "The URL of Prysm's gRPC API endpoint for your external Beacon Node. Prysm's Validator Client will need this in order to connect to it.\nNOTE: If you are running it on the same machine as this node, addresses like `localhost` and `127.0.0.1` will not work due to Docker limitations. Enter your machine's LAN IP address instead, for example 'http://192.168.1.100:5053'."
	// cfg.PrysmRpcUrl.AffectedContainers = []string{string(ContainerID_ValidatorClient)}

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
	// cfg.BeaconNode.AffectedContainers = []string{string(ContainerID_ValidatorClient)}
	cfg.BeaconNode.Options = options

	return cfg
}

// The title for the config
func (cfg *ExternalBeaconConfig) GetTitle() string {
	return "External Beacon Node"
}

// Get the parameters for this config
func (cfg *ExternalBeaconConfig) GetParameters() []hdconfig.IParameter {
	return []hdconfig.IParameter{
		&cfg.BeaconNode,
		&cfg.HttpUrl,
		&cfg.PrysmRpcUrl,
	}
}

// Get the sections underneath this one
func (cfg *ExternalBeaconConfig) GetSections() map[string]hdconfig.ISection {
	return map[string]hdconfig.ISection{}
}
