package apiv1_connections

import (
	"encoding/json"
	"log"
	"net/http"

	websocket "github.com/jnnkrdb/gokv/pkg/server/webSocket"
)

func ListConnections(w http.ResponseWriter, r *http.Request) {
	var list = []string{}
	for n := range websocket.Connections {
		list = append(list, n)
	}
	if err := json.NewEncoder(w).Encode(list); err != nil {

		log.Printf("[ERR] couldn't parse list to json: %v, err: %v\n", list, err)

		http.Error(w, "couldn't parse list to json", http.StatusInternalServerError)
	}
}
