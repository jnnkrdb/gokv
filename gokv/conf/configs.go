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
	SELF_NAME                   string
	SELF_NAMESPACE              string
	SELF_WEBSOCKET_SERVICE_NAME string
	SELF_UID                    string
	CLUSTER_INTERNAL_DOMAIN     string
	NC                          *NodeConfig = &NodeConfig{}
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

	type envloader struct {
		EnvVar      string
		Dest        *string
		DefaultFunc func() string
	}

	loadEnvs := func(el envloader) {
		// ------------------------------------------------------ run the functions to load the envs
		if s, ok := os.LookupEnv(el.EnvVar); ok {
			*el.Dest = s
		} else {
			*el.Dest = el.DefaultFunc()
		}
	}

	for _, el := range []envloader{
		{
			EnvVar: "INSTANCE_NAME", Dest: &SELF_NAME, DefaultFunc: func() string {
				if s, err := os.Hostname(); err != nil {
					log.Fatalf("[ERR] where is my hostname?: %s\n", err.Error())
					return ""
				} else {
					return s
				}
			},
		}, // get the own instance name
		{EnvVar: "INSTANCE_NAMESPACE", Dest: &SELF_NAMESPACE, DefaultFunc: func() string { return "default" }},                       // get the own instance namespace
		{EnvVar: "SELF_WEBSOCKET_SERVICE_NAME", Dest: &SELF_WEBSOCKET_SERVICE_NAME, DefaultFunc: func() string { return "gokv-ws" }}, // get the own instances websocket service name
		{EnvVar: "CLUSTER_INTERNAL_DOMAIN", Dest: &CLUSTER_INTERNAL_DOMAIN, DefaultFunc: func() string { return "cluster.local" }},   // get the cluster interal url
		{EnvVar: "INSTANCE_UID", Dest: &SELF_UID, DefaultFunc: func() string { return uuid.New().String() }},                         // get the uid, which is used to identify the websocket connections
	} {
		loadEnvs(el)
	}
}

// this struct is parsed from an config yaml
type NodeConfig struct {

	// this key says if the debug function should be enabled
	Debug bool `yaml:"debug"`

	// this section handles the stoarge configs
	Storage struct {
		Type string `yaml:"type"`
	} `yaml:"storage"`
}
