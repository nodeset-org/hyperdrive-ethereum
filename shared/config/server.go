package config

import (
	"github.com/nodeset-org/hyperdrive-ethereum/shared"
	"github.com/nodeset-org/hyperdrive-ethereum/shared/ids"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
)

type PortMode string

const (
	PortMode_Closed    PortMode = "closed"
	PortMode_Localhost PortMode = "localhost"
	PortMode_External  PortMode = "external"
)

const (
	containerTag string = "nodeset/hyperdrive-ethereum-service:v" + shared.HyperdriveEthereumVersion
)

const (
	// Name of the service container
	ServiceContainerName string = "ethereum"

	// Name of the log file for the service
	ServiceLogFile string = "service.log"

	// Default port for the server API
	DefaultServerApiPort uint16 = 8080
)

type ServerConfig struct {
	hdconfig.SectionHeader

	ContainerTag hdconfig.StringParameter

	Port hdconfig.UintParameter

	PortMode hdconfig.ChoiceParameter[PortMode]
}

type ServerConfigSettings struct {
	ContainerTag string `json:"containerTag" yaml:"containerTag"`

	Port     uint64   `json:"port" yaml:"port"`
	PortMode PortMode `json:"portMode" yaml:"portMode"`
}

func NewServerConfig() *ServerConfig {
	cfg := &ServerConfig{}
	cfg.ID = hdconfig.Identifier(ids.LocalServerConfigID)
	cfg.Name = "Service Config"
	cfg.Description.Default = "This is the configuration for the module's service. This isn't used by the service directly, but it is used by Hyperdrive itself in the service's Docker Compose file template to configure the service during its starting process."

	// Container Tag
	cfg.ContainerTag.ID = hdconfig.Identifier(ids.ContainerTagID)
	cfg.ContainerTag.Name = "Container Tag"
	cfg.ContainerTag.Description.Default = "This is the tag used for the service's Docker container image."
	cfg.ContainerTag.Default = containerTag

	// Port
	cfg.Port.ID = hdconfig.Identifier(ids.PortID)
	cfg.Port.Name = "API Port"
	cfg.Port.Description.Default = "This is the API port the server should run on."
	cfg.Port.Default = uint64(DefaultServerApiPort)
	cfg.Port.MinValue = 0
	cfg.Port.MaxValue = 65535

	// Options for PortMode
	options := make([]hdconfig.ParameterOption[PortMode], 3)
	options[0].Name = "Closed"
	options[0].Description.Default = "The API is only accessible to internal Docker container traffic."
	options[0].Value = PortMode_Closed

	options[1].Name = "Localhost Only"
	options[1].Description.Default = "The API is accessible from internal Docker containers and your own local machine, but no other external machines."
	options[1].Value = PortMode_Localhost

	options[2].Name = "All External Traffic"
	options[2].Description.Default = "The port is accessible to everything, including external machines.\n\n[orange]Use with caution!"
	options[2].Value = PortMode_External

	// PortMode
	cfg.PortMode.ID = hdconfig.Identifier(ids.PortModeID)
	cfg.PortMode.Name = "Expose API Port"
	cfg.PortMode.Description.Default = "Determine how the server's HTTP API restricts its access from various sources."
	cfg.PortMode.Options = options
	cfg.PortMode.Default = PortMode_Closed

	return cfg
}

func (cfg ServerConfig) GetParameters() []hdconfig.IParameter {
	return []hdconfig.IParameter{
		&cfg.ContainerTag,
		&cfg.Port,
		&cfg.PortMode,
	}
}

func (cfg ServerConfig) GetSections() []hdconfig.ISection {
	return []hdconfig.ISection{}
}

func (cfg *ServerConfig) GetDescription() hdconfig.DynamicProperty[string] {
	return hdconfig.DynamicProperty[string]{}
}

func (cfg *ServerConfig) GetDisabled() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}

func (cfg *ServerConfig) GetHidden() hdconfig.DynamicProperty[bool] {
	return hdconfig.DynamicProperty[bool]{}
}
