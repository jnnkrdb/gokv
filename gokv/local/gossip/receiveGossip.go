package gossip

import (
	"encoding/json"
	"log"
	"net"

	"github.com/jnnkrdb/gokv/pkg/messaging"
	"github.com/jnnkrdb/gokv/pkg/server/tcpSocket"
)

// this function handles the incomming gossip tcp connections and
// handles the response
func ReceiveGossip(c net.Conn) {
	defer c.Close()

	var decoder = json.NewDecoder(c)
	var req = &tcpSocket.TCPRequest{}
	var resp = tcpSocket.TCPResponse{}

	if err := decoder.Decode(&req); err != nil {

		log.Printf("[WRN] error decoding to json: %v\n", err)
		log.Printf("[WRN] dropping request\n")

		resp.RequestType = messaging.RT_Dropped
		resp.RequestCode = messaging.RC_Error
		resp.Load = "error decoding payload from json"

		tcpSocket.SendJSON(c, resp)
		return
	}

	log.Printf("[INF] received request: Type [%s], Code [%s]\n", req.RequestType.String(), req.RequestCode.String())

	resp.RequestType = req.RequestType
	resp.Load = `{"msg":"received your request successfully!"}`
	tcpSocket.SendJSON(c, resp)
	return
}
