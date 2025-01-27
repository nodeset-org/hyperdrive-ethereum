package config

import (
	sharedconfig "github.com/nodeset-org/hyperdrive-ethereum/shared/config"
	hdconfig "github.com/nodeset-org/hyperdrive/modules/config"
)

func NewHyperdriveEthereumConfig() *sharedconfig.HyperdriveEthereumConfig {
	cfg := &sharedconfig.HyperdriveEthereumConfig{}

	// TODO: HN
	// // ExampleBool
	// cfg.ExampleBool.ID = hdconfig.Identifier(ids.ExampleBoolID)
	// cfg.ExampleBool.Name = "Example Boolean"
	// cfg.ExampleBool.Description.Default = "This is an example of a boolean parameter. It doesn't directly affect the service, but it does control the behavior of some other config parameters."
	// cfg.ExampleBool.AffectedContainers = []string{shared.ServiceContainerName}
	// cfg.ExampleBool.Value = cfg.ExampleBool.Default

	// // ExampleInt
	// cfg.ExampleInt.ID = hdconfig.Identifier(ids.ExampleIntID)
	// cfg.ExampleInt.Name = "Example Integer"
	// cfg.ExampleInt.Description.Default = "This is an example of an integer parameter."
	// cfg.ExampleInt.AffectedContainers = []string{shared.ServiceContainerName}
	// cfg.ExampleInt.Value = cfg.ExampleInt.Default

	// // ExampleUint
	// cfg.ExampleUint.ID = hdconfig.Identifier(ids.ExampleUintID)
	// cfg.ExampleUint.Name = "Example Unsigned Integer"
	// cfg.ExampleUint.Description.Default = "This is an example of an unsigned integer parameter."
	// cfg.ExampleUint.AffectedContainers = []string{shared.ServiceContainerName}
	// cfg.ExampleUint.Value = cfg.ExampleUint.Default

	// // ExampleFloat
	// cfg.ExampleFloat.ID = hdconfig.Identifier(ids.ExampleFloatID)
	// cfg.ExampleFloat.Name = "Example Float"
	// cfg.ExampleFloat.Description.Default = "This is an example of a float parameter with a minimum and maximum set."
	// cfg.ExampleFloat.Default = 50
	// cfg.ExampleFloat.MinValue = 0.0
	// cfg.ExampleFloat.MaxValue = 100.0
	// cfg.ExampleFloat.Value = cfg.ExampleFloat.Default
	// cfg.ExampleFloat.AffectedContainers = []string{shared.ServiceContainerName}

	// // ExampleString
	// cfg.ExampleString.ID = hdconfig.Identifier(ids.ExampleStringID)
	// cfg.ExampleString.Name = "Example String"
	// cfg.ExampleString.Description.Default = "This is an example of a string parameter. It has a max length and regex pattern set."
	// cfg.ExampleString.MaxLength = 10
	// cfg.ExampleString.Regex = "^[a-zA-Z]*$"
	// cfg.ExampleString.Value = cfg.ExampleString.Default
	// cfg.ExampleString.AffectedContainers = []string{shared.ServiceContainerName}

	// // Options for ExampleChoice
	// options := make([]hdconfig.ParameterMetadataOption[nativecfg.ExampleOption], 3)
	// options[0].Name = "One"
	// options[0].Description.Default = "This is the first option."
	// options[0].Value = nativecfg.ExampleOption_One

	// thresholdString := strconv.FormatFloat(FloatThreshold, 'f', -1, 64)
	// options[1].Name = "Two"
	// options[1].Description.Default = "This is the second option. It is hidden when ExampleFloat is less than " + thresholdString + "."
	// options[1].Description.Template = fmt.Sprintf("{{if lt .GetValue %s %s}}This option is hidden because the float is less than %s.{{else}}This option is visible because the float is greater than or equal to %s.{{end}}", ids.ExampleFloatID, thresholdString, thresholdString, thresholdString)
	// options[1].Value = nativecfg.ExampleOption_Two
	// options[1].Disabled.Default = true
	// options[1].Disabled.Template = "{{if eq .GetValue " + ids.ExampleBoolID + " true}}false{{else}}{{.UseDefault}}{{end}}"

	// options[2].Name = "Three"
	// options[2].Description.Default = "This is the third option."
	// options[2].Value = nativecfg.ExampleOption_Three

	// // ExampleChoice
	// cfg.ExampleChoice.ID = hdconfig.Identifier(ids.ExampleChoiceID)
	// cfg.ExampleChoice.Name = "Example Choice"
	// cfg.ExampleChoice.Description.Default = "This is an example of a choice parameter between multiple options."
	// cfg.ExampleChoice.Options = options
	// cfg.ExampleChoice.Default = options[0].Value
	// cfg.ExampleChoice.Value = cfg.ExampleChoice.Default
	// cfg.ExampleChoice.AffectedContainers = []string{}

	// // Subconfigs
	// cfg.SubConfig = NewSubConfig()
	// cfg.ServerConfig = NewServerConfig()

	return cfg
}

func (cfg *sharedconfig.HyperdriveEthereumConfig) GetParameters() []hdconfig.IParameter {
	return []hdconfig.IParameter{
		// &cfg.ExampleBool,
		// &cfg.ExampleInt,
		// &cfg.ExampleFloat,
		// &cfg.ExampleString,
		// &cfg.ExampleChoice,
	}
}

func (cfg *sharedconfig.HyperdriveEthereumConfig) GetSections() []hdconfig.ISection {
	return []hdconfig.ISection{
		// cfg.SubConfig,
		// cfg.ServerConfig,
	}
}

func CreateInstanceFromNativeConfig(native *sharedconfig.NativeHyperdriveEthereumConfig) *ExampleConfigSettings {
	instance := &sharedconfig.HyperdriveEthereumConfig{
		// ExampleBool:   native.ExampleBool,
		// ExampleInt:    native.ExampleInt,
		// ExampleFloat:  native.ExampleFloat,
		// ExampleString: native.ExampleString,
		// ExampleChoice: native.ExampleChoice,
		// SubConfig: &SubConfigSettings{
		// 	SubExampleBool:   native.SubConfig.SubExampleBool,
		// 	SubExampleChoice: native.SubConfig.SubExampleChoice,
		// },
		// ServerConfig: &ServerConfigSettings{},
	}
	return instance
}

func ConvertInstanceToNativeConfig(instance *ExampleConfigSettings) *sharedconfig.NativeHyperdriveEthereumConfig {
	native := &sharedconfig.NativeHyperdriveEthereumConfig{
		// ExampleBool:   instance.ExampleBool,
		// ExampleInt:    instance.ExampleInt,
		// ExampleFloat:  instance.ExampleFloat,
		// ExampleString: instance.ExampleString,
	}
	return native
}
