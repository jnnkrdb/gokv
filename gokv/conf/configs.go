package conf

import (
	"log"
	"os"

	"github.com/google/uuid"
	"gopkg.in/yaml.v3"
)

// these ports are used for communication for
// client to server and server to server
const (
	HTTP_PORT   int = 80
	GOSSIP_PORT int = 5334

	CONFIG_YAML string = "/opt/gokv/cfg/gokv.yaml"
)

var (
	SELF_NAME string
	SELF_UID  string
	NC        *NodeConfig = &NodeConfig{}
)

// initialize the service from config file at $GOKV_HOME/gokv.yaml
func init() {

	// load the config from the config.yaml
	if yamlF, err := os.ReadFile(CONFIG_YAML); err != nil {
		log.Fatalf("[ERR] couldn't read config file: %s\n", err.Error())
	} else {
		if err = yaml.Unmarshal(yamlF, NC); err != nil {
			log.Fatalf("[ERR] couldn't parse config file: %s\n", err.Error())
		}
	}

	// get the own instance name
	if s, ok := os.LookupEnv("INSTANCE_NAME"); ok {
		SELF_NAME = s
	} else {
		if s, err := os.Hostname(); err != nil {
			log.Fatalf("[ERR] where is my hostname?: %s\n", err.Error())
		} else {
			SELF_NAME = s
		}
	}

	// get the uid, which is used to identify the websocket connections
	if s, ok := os.LookupEnv("INSTANCE_UID"); ok {
		SELF_UID = s
	} else {
		SELF_UID = uuid.New().String()
	}
}

// this struct is parsed from an config yaml
type NodeConfig struct {

	// this anonymous struct contains information about the ha cluster
	HA struct {
		Nodes          []string `yaml:"nodes"`
		SyncTimeoutSec int      `yaml:"syncTimeoutSec"`
	} `yaml:"ha"`

	// this key says if the debug function should be enabled
	Debug bool `yaml:"debug"`

	// this section handles the stoarge configs
	Storage struct {
		Type string `yaml:"type"`
	} `yaml:"storage"`
}
