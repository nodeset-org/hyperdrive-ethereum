package config

const clientDataVolumeName string = "/ethclient"

// A Docker container name
type ContainerID string

// Enum to describe the names / IDs of various containers controlled by NMC
const (
	// Unknown
	ContainerID_Unknown ContainerID = ""

	// The daemon
	ContainerID_Daemon ContainerID = "daemon"

	// The Execution client
	ContainerID_ExecutionClient ContainerID = "ec"

	// The Beacon node (Beacon Node)
	ContainerID_BeaconNode ContainerID = "bn"

	// The Validator client
	ContainerID_ValidatorClient ContainerID = "vc"

	// MEV-Boost
	ContainerID_MevBoost ContainerID = "mev-boost"

	// The Node Exporter
	ContainerID_Exporter ContainerID = "exporter"

	// Prometheus
	ContainerID_Prometheus ContainerID = "prometheus"

	// Grafana
	ContainerID_Grafana ContainerID = "grafana"
)

// A client ownership mode
type ClientMode string

// Enum to describe client modes
const (
	// Unknown
	ClientMode_Unknown ClientMode = ""

	// Locally-owned clients (managed by the NMC service)
	ClientMode_Local ClientMode = "local"

	// Externally-managed clients (managed by the user)
	ClientMode_External ClientMode = "external"
)
