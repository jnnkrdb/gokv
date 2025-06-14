package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/jnnkrdb/gokv/conf"
)

var (
	upgrader = websocket.Upgrader{}

	wsRouter *mux.Router = mux.NewRouter()
)

func RunWS() {

	// handle the connections api
	wsRouter.HandleFunc("/api/v1/connections", func(w http.ResponseWriter, r *http.Request) {

		// adding the own node to the list, since the self-node does
		// not have a websocket connection to itself
		var list = []string{conf.SELF_NAME}

		// range over the existing connections and add them to the
		// nodes list
		// these connections should be active
		for n := range Connections {

			list = append(list, n)
		}

		// parse as json and send
		if err := json.NewEncoder(w).Encode(NodePool{Nodes: list}); err != nil {

			log.Printf("[ERR] couldn't parse list to json: %v, err: %v\n", list, err)

			http.Error(w, "couldn't parse list to json", http.StatusInternalServerError)
		}

	}).Methods("GET", "OPTIONS")

	// handle the websocket registration
	wsRouter.HandleFunc(WebsocketPath, func(w http.ResponseWriter, r *http.Request) {

		// if debug is enabled, then print the received headers
		if conf.NC.Debug {
			log.Printf("[INF][%s] received headers: %v\n", r.URL.String(), r.Header)
		}

		// upgrade the connection to a websocket
		if c, err := upgrader.Upgrade(w, r, WsHeader); err != nil {

			log.Printf("[WRN][%s] error upgrading to websocket conn\n", r.URL.String())

		} else {

			go HandleWebSocketConnection(r.Header.Get("gokv-node"), c)
		}
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", conf.GOSSIP_PORT), wsRouter))
}
