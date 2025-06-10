package synchro

import "fmt"

type Host struct {
	URI           string `json:"uri" yaml:"uri"`
	GossipPortTCP int    `json:"gossip-port-tcp" yaml:"gossipPortTCP"`
}

// create string from Host -> format: "<uri>:<gossip-port-tcp>""
func (h Host) String() string {
	return fmt.Sprintf("%s:%d", h.URI, h.GossipPortTCP)
}
