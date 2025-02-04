package internal_test

const (
	AdapterTag                  string = "nodeset/hyperdrive-ethereum-adapter:v0.2.0"
	ServiceTag                  string = "nodeset/hyperdrive-ethereum-service:v0.2.0"
	ProjectName                 string = "he-test"
	GlobalAdapterContainerName  string = "hd_em_adapter"
	ProjectAdapterContainerName string = "hd_" + ProjectName + "_em_adapter"
	ServiceContainerName        string = ProjectName + "_ethereum"
	LogDir                      string = "/tmp/he-adapter-test/log"
	SystemDir                   string = "/tmp/he-adapter-test/sys"
	CfgDir                      string = "/tmp/he-adapter-test/cfg"
	UserDir                     string = "/tmp/he-adapter-test/user"
	KeyPath                     string = UserDir + "/secrets/adapter.key"
	UserDataPath                string = "/tmp/he-adapter-test/data"
	TestKey                     string = "test-key"
)
