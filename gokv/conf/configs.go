package conf

import (
	"log"
	"os"

	"github.com/jnnkrdb/gokv/pkg/synchro"
	"gopkg.in/yaml.v3"
)

// these ports are used for communication for
// client to server and server to server
const (
	HTTP_PORT   int = 80
	GOSSIP_PORT int = 5334
)

var (
	NC *NodeConfig = &NodeConfig{}
)

// initialize the service from config file at $GOKV_HOME/gokv.yaml
func init() {
	yamlF, err := os.ReadFile("/opt/gokv/cfg/gokv.yaml")
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
		Nodes          []synchro.Host `yaml:"nodes"`
		SyncTimeoutSec int            `yaml:"syncTimeoutSec"`
	} `yaml:"ha"`
}
