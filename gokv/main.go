package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jnnkrdb/gokv/conf"
)

func main() {

	// activate logging
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	fmt.Printf("Hello world!")

	for {

		log.Printf("New Line, HomeDir: %s, BinDir: %s", conf.GOKV_HOME, conf.GOKV_BINARY_PATH)

		log.Printf("Nodes:\n")

		for i, node := range conf.NC.HA.Nodes {
			log.Printf("Node %d: ws://%s:%d", i, node.URI, node.GossipPortTCP)
		}

		time.Sleep(30 * time.Second)
	}
}
