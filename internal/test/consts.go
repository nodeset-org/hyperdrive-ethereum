package internal_test

const (
	AdapterTag                  string = "nodeset/hyperdrive-ethereum-adapter:v0.1.0"
	ServiceTag                  string = "nodeset/hyperdrive-ethereum-service:v0.1.0"
	ProjectName                 string = "hde-test"
	GlobalAdapterContainerName  string = "hd_em_adapter"
	ProjectAdapterContainerName string = "hd_" + ProjectName + "_em_adapter"
	ServiceContainerName        string = ProjectName + "_ethereum"
	LogDir                      string = "/tmp/hde-adapter-test/log"
	SystemDir                   string = "/tmp/hde-adapter-test/sys"
	CfgDir                      string = "/tmp/hde-adapter-test/cfg"
	UserDir                     string = "/tmp/hde-adapter-test/user"
	KeyPath                     string = UserDir + "/secrets/adapter.key"
	UserDataPath                string = "/tmp/hde-adapter-test/data"
	TestKey                     string = "test-key"
)
