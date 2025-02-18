package config

type DockerSettings struct {
	BeaconNodeContainer      string `json:"beaconNodeContainer"`
	ExecutionClientContainer string `json:"executionClientContainer"`
	DockerNetwork            string `json:"dockerNetwork"`
	BeaconDataVolume         string `json:"beaconDataVolume"`
	ExecutionDataVolume      string `json:"executionDataVolume"`
	ContainerTag             string `json:"containerTag"`
}
