package conf

import (
	"log"
	"net"
	"strings"
	"sync"
)

var (

	// list of open connections of the pool
	GOSSIP_CONNS       = make(map[string]*net.TCPConn)
	GOSSIP_CONNS_MUTEX = sync.Mutex{}
)

// check all possible connections and if they are open or not
func CheckConnections() {

	// prohibit access to connection pool for as long
	// as the connections get checked or created
	GOSSIP_CONNS_MUTEX.Lock()

	// calculate the nodes where a connection is needed
	for _, rn := range NC.HA.Nodes {
		if strings.Contains(rn.URI, SELF_NAME) {
			continue
		}

		// create connection if no exist
		if _, ok := GOSSIP_CONNS[rn.String()]; !ok {
			if tcpServer, err := net.ResolveTCPAddr("tcp", rn.String()); err != nil {
				log.Printf("[ERR] error resolving tcpServer [%s]: %v\n", rn, err)
				continue
			} else {
				if conn, err := net.DialTCP("tcp", nil, tcpServer); err != nil {
					log.Printf("[ERR] error dialing tcpServer [%s]: %v\n", rn, err)
					continue
				} else {
					GOSSIP_CONNS[rn.String()] = conn
				}
			}
		}
	}

	// finish lock
	GOSSIP_CONNS_MUTEX.Unlock()
}
