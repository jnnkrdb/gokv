package apiv1_storage

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/gokv/conf"
	"github.com/jnnkrdb/gokv/pkg/gossip/functions"
	"github.com/jnnkrdb/gokv/pkg/messaging"
	websocket "github.com/jnnkrdb/gokv/pkg/server/webSocket"
)

func DeleteKey(w http.ResponseWriter, r *http.Request) {

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

	if err := conf.STORAGE.DeleteKey(b, k); err != nil {
		log.Printf("[ERR] couldn't delete key from bucket: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {

		websocket.Connections.Send(messaging.RC_SyncStorage, functions.SyncStorageLoad{
			Sync:   functions.SyncStorageLoad_Sync_Delete,
			Bucket: b,
			Key:    k,
		})

		w.Write([]byte("OK"))
	}
}
