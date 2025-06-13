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

func DeleteBucket(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	b, ok := vars["bucket"]
	if !ok {
		log.Printf("[WRN] bucket var is missing\n")
		http.Error(w, "bucketVar is missing", http.StatusBadRequest)
		return
	}

	if err := conf.STORAGE.DeleteBucket(b); err != nil {
		log.Printf("[ERR] couldn't delete bucket from storage: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {

		websocket.Connections.Send(messaging.RC_SyncStorage, functions.SyncStorageLoad{
			Sync:   functions.SyncStorageLoad_Sync_Delete,
			Bucket: b,
		})

		w.Write([]byte("OK"))
	}
}
