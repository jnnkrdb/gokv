package gossip

import (
	"encoding/json"
	"log"
	"net"

	"github.com/jnnkrdb/gokv/local/gossip/functions"
	"github.com/jnnkrdb/gokv/pkg/server/tcpSocket"
)

// this function handles the incomming gossip tcp connections and
// handles the responses
func ReceiveGossip(c net.Conn) {
	defer c.Close()

	var decoder = json.NewDecoder(c)
	var req = &tcpSocket.TCPRequest{}
	var resp = tcpSocket.TCPResponse{}

	if err := decoder.Decode(&req); err != nil {
		log.Printf("[WRN] error decoding to json: %v\n", err)
		log.Printf("[WRN] dropping request\n")
		return
	}

	resp.Initiator = req.Initiator
	resp.RequestCmd = req.RequestCmd

	functions.Handle(*req, &resp)
	tcpSocket.SendJSON(c, resp)
	return
}
