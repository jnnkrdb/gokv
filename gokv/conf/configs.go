package conf

import (
	"log"
	"os"

	"github.com/jnnkrdb/gokv/synchro"
	"gopkg.in/yaml.v3"
)

var (
	GOKV_HOME        string = ReadEnv("GOKV_HOME", "NONE")
	GOKV_BINARY_PATH string = ReadEnv("GOKV_BINARY_PATH", "NONE")

	NC *NodeConfig = &NodeConfig{}
)

// initialize the service from config file at $GOKV_HOME/gokv.yaml
func init() {
	yamlF, err := os.ReadFile(GOKV_HOME + "/gokv.yaml")
	if err != nil {
		log.Fatalf("couldn't read config file: %s", err.Error())
	}

	err = yaml.Unmarshal(yamlF, NC)
	if err != nil {
		log.Fatalf("couldn't parse config file: %s", err.Error())
	}

}

// this struct is parsed from an config yaml
type NodeConfig struct {

	// this anonymous struct contains information about the ha cluster
	HA struct {
		Nodes []synchro.Host `yaml:"nodes"`
	} `yaml:"ha"`
}
