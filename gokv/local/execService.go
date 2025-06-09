package local

import (
	"log"
	"time"

	"github.com/jnnkrdb/gokv/conf"
	"github.com/jnnkrdb/gokv/local/gossip"
	"github.com/jnnkrdb/gokv/pkg/server/tcpSocket"
)

func RunService() {

	// start tcp socket
	go tcpSocket.RunTCPSocket(conf.GOSSIP_PORT, gossip.ReceiveGossip)

	// start gosspi spreading
	go gossip.SpreadGossip()

	for {
		log.Printf("Nodes:\n")

		for i, node := range conf.NC.HA.Nodes {
			log.Printf("Node %d: ws://%s", i, node.String())
		}

		time.Sleep(30 * time.Second)
	}
}
