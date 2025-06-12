package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func RunWS(port int) {
	http.HandleFunc(WebsocketPath, func(w http.ResponseWriter, r *http.Request) {

		// upgrade the connection to a websocket
		c, err := upgrader.Upgrade(w, r, WsHeader)
		if err != nil {
			log.Printf("[WRN][%s] error upgrading to websocket conn\n", r.URL.String())
			return
		}

		HandleWebSocketConnection(r.Header.Get("gokv.jnnkrdb.de/node"), c)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
