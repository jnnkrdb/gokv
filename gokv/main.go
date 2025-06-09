package main

import (
	"log"

	"github.com/jnnkrdb/gokv/conf"
	"github.com/jnnkrdb/gokv/local"
)

func main() {

	// activate logging
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	// print configs
	log.Printf("config:\n%v\n\n", *conf.NC)

	// run node
	local.RunService()
}
