package tcpSocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

// this function starts the tcp socket, which is used to "spread the gossip between nodes"
func RunTCPSocket(port int, handlerFuncAsync func(net.Conn)) {
	// starting tcp socket
	log.Printf("[INF] starting tcp-socket on localhost:%d\n", port)
	if lstn, err := net.Listen("tcp", fmt.Sprintf(":%d", port)); err != nil {
		log.Panic("[ERR] could not start tcp-socket: %s\n", err.Error())
	} else {
		defer lstn.Close()
		for {
			if _c, err := lstn.Accept(); err != nil {
				log.Printf("[WRN] error accept: %v\n", err)
			} else {
				go handlerFuncAsync(_c)
			}
		}
	}
}

// send an json struct to an open tcp connection
func SendJSON(c net.Conn, response interface{}) {
	if byt, err := json.Marshal(response); err != nil {
		log.Printf("[ERR] error marshalling json: %v\n", err)
		c.Write([]byte(`{"msg":"error decoding json payload to []byte"}`))
	} else {
		c.Write(byt)
	}
}
