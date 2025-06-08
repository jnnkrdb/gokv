package synchro

import "fmt"

var KnownHosts []Host

type Host struct {
	URI           string `json:"uri" yaml:"uri"`
	GossipPortTCP int    `json:"gossip-port-tcp" yaml:"gossipPortTCP"`
}

// create string from Host
func (h Host) String() string {
	return fmt.Sprintf("%s:%d", h.URI, h.GossipPortTCP)
}

// this function reads the known hosts from a config file and
// parses them into a list
func init() {

}
