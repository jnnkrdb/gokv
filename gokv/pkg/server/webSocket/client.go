package websocket

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/jnnkrdb/gokv/conf"
)

// create the required connections to the other nodes
func CreateWSConnections() {

	for _, node := range conf.NC.HA.Nodes {

		if strings.Contains(node, conf.SELF_NAME) {
			continue
		}

		// create url
		u := url.URL{
			Scheme: "ws",
			Host:   fmt.Sprintf("%s:%d", node, conf.GOSSIP_PORT),
			Path:   WebsocketPath,
		}
		log.Printf("[INF] connecting to url: %s\n", u.String())

		if c, _, err := websocket.DefaultDialer.Dial(u.String(), WsHeader); err != nil {

			log.Printf("[ERR] couldn't connect to [%s]: %v\n", u.String(), err)

		} else {

			HandleWebSocketConnection(node, c)
		}
	}
}

/*
func runClient(port int) {

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
*/
