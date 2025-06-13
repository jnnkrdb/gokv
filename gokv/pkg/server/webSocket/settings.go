package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/jnnkrdb/gokv/conf"
	"github.com/jnnkrdb/gokv/pkg/messaging"
)

const (
	WebsocketPath string = "/gossip"
)

var (
	Connections = make(WsPool)
)

type WsPool map[string]*websocket.Conn

func (wsp *WsPool) Send(request messaging.RequestCommand, vjson interface{}) error {

	// used for the request, to identify later on
	var uid string = uuid.New().String()

	b, err := json.Marshal(vjson)
	if err != nil {
		log.Printf("[ERR] error parsing vjson: %v\n", err)
		return err
	}

	var (
		e  error = nil
		wg       = sync.WaitGroup{}
	)

	// add every conn to the waitgroup
	wg.Add(len(*wsp))

	// run through every connection and send the json struct
	for n, conn := range *wsp {

		// run as goroutines to speed up requests
		go func() {
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
		}()
	}

	// wait for all requests to be finished
	wg.Wait()

	return e
}
