package apiv1_storage

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/gokv/conf"
	"github.com/jnnkrdb/gokv/pkg/gossip/functions"
	"github.com/jnnkrdb/gokv/pkg/messaging"
	websocket "github.com/jnnkrdb/gokv/pkg/server/webSocket"
)

func WrityKey(w http.ResponseWriter, r *http.Request) {

	var content string

	vars := mux.Vars(r)
	b, ok := vars["bucket"]
	if !ok {
		log.Printf("[WRN] bucket var is missing\n")
		http.Error(w, "bucketVar is missing", http.StatusBadRequest)
		return
	}
	k, ok := vars["key"]
	if !ok {
		log.Printf("[WRN] key var is missing\n")
		http.Error(w, "keyVar is missing", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		log.Printf("[ERR] couldn't get body content from json: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := conf.STORAGE.Write(b, k, content); err != nil {
		log.Printf("[ERR] couldn't write content into storage: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {

		websocket.Connections.Send(messaging.RC_SyncStorage, functions.SyncStorageLoad{
			Sync:   functions.SyncStorageLoad_Sync_Write,
			Bucket: b,
			Key:    k,
			Value:  content,
		})

		w.Write([]byte("OK"))
	}
}
