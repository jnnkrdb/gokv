package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func RunWS(port int) {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}

func home(w http.ResponseWriter, r *http.Request) {

	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func echo(w http.ResponseWriter, r *http.Request) {

	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()

	for {

		mt, message, err := c.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("recv: %s", message)

		err = c.WriteMessage(mt, message)

		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
