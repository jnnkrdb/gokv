package gossip

import (
	"log"
	"time"

	"github.com/jnnkrdb/gokv/conf"
)

// this function connects other tcp sockets from the list and send them
// requests to synchronize their knowledge
func SpreadGossip() {

	// set a defautl value if none is given
	if conf.NC.HA.SyncTimeoutSec == 0 {
		conf.NC.HA.SyncTimeoutSec = 10
	}

	for {
		log.Printf("[INF] another gossip round\n")

		time.Sleep(time.Second * time.Duration(conf.NC.HA.SyncTimeoutSec))

	}
}
