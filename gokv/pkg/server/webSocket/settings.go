package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/jnnkrdb/gokv/conf"
	"github.com/jnnkrdb/gokv/pkg/messaging"
)

const (
	WebsocketPath string = "/"
)

var (
	Connections = make(WsPool)
	WsHeader    = make(http.Header)
)

func init() {
	WsHeader.Add("x-gokv-node", conf.SELF_NAME)
	WsHeader.Add("x-gokv-auth", conf.SELF_UID)
}

type WsPool map[string]*websocket.Conn

func (wsp *WsPool) Send(request messaging.RequestCommand, vjson interface{}) error {

	var uid string = uuid.New().String()

	b, err := json.Marshal(vjson)
	if err != nil {
		log.Printf("[ERR] error parsing vjson: %v\n", err)
		return err
	}

	var e error = nil

	for n, conn := range *wsp {
		log.Printf("[INF][%s] error parsing vjson: %v\n", n, err)

		if err := conn.WriteJSON(messaging.WsRequest{
			Initiator:      conf.SELF_NAME,
			RequestCommand: request,
			RequestUID:     uid,
			Load:           b,
		}); err != nil {
			log.Printf("[WRN][%s] error sending ws request: %v\n", n, err)
			e = fmt.Errorf("%s%s", e.Error(), fmt.Errorf("[%s]%s", n, err.Error()))
		}
	}

	return e
}
