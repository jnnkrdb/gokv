package main

import (
	"flag"
	"log"

	"github.com/jnnkrdb/gokv/conf"
	"github.com/jnnkrdb/gokv/local"
)

func main() {

	flag.IntVar(&conf.GOSSIP_PORT, "gossip-port", 3453, "The port used by the service, to spread gossip between the nodes.")
	flag.IntVar(&conf.HTTP_PORT, "http-port", 80, "The port used by the service, to provide the http server.")

	// parse the flags
	flag.Parse()

	// activate logging
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	// print configs
	log.Printf("config:\n%v\n\n", *conf.NC)

	// run node
	local.RunService()
}
