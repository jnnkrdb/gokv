package gossip

import (
	"encoding/json"
	"log"
	"net"
	"sync"
	"time"

	"github.com/jnnkrdb/gokv/conf"
	"github.com/jnnkrdb/gokv/pkg/server/tcpSocket"
)

// this function connects other tcp sockets from the list and send them
// requests to synchronize their knowledge
func SpreadGossipTCP(req tcpSocket.TCPRequest) {

	conf.GOSSIP_CONNS_MUTEX.Lock()

	var wg = sync.WaitGroup{}
	wg.Add(len(conf.GOSSIP_CONNS))
	for i := range conf.GOSSIP_CONNS {

		// run the spread func as a go routine
		go func(c *net.TCPConn) {
			defer wg.Done()

			var start = time.Now()

			if err := json.NewEncoder(c).Encode(req); err != nil {
				log.Printf("[WRN][%s] couldn't send json via tcp: %v\n", c.RemoteAddr().String(), err)
				return
			}

			var resp = tcpSocket.TCPResponse{}
			if err := json.NewDecoder(c).Decode(&resp); err != nil {
				log.Printf("[WRN][%s] couldn't receive json via tcp: %v\n", c.RemoteAddr().String(), err)
				return
			}

			log.Printf("[INF][%s] received response State[%s]: %v\n", c.RemoteAddr().String(), resp.RequestState, resp.Load)
			log.Printf("[INF][%s] needed %dms\n", c.RemoteAddr().String(), time.Since(start).Milliseconds())

		}(conf.GOSSIP_CONNS[i])
	}

	wg.Wait()

	conf.GOSSIP_CONNS_MUTEX.Lock()
}
