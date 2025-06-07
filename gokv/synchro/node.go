package synchro

var KnownHosts []Host

type Host struct {
	URI           string `json:"uri" yaml:"uri"`
	GossipPortTCP int    `json:"gossip-port-tcp" yaml:"gossipPortTCP"`
}

// this function reads the known hosts from a config file and
// parses them into a list
func init() {

}
