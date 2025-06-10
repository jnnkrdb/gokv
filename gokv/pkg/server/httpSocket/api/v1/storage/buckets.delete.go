package apiv1_storage

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/gokv/conf"
	"github.com/jnnkrdb/gokv/local/gossip"
	"github.com/jnnkrdb/gokv/local/gossip/functions"
	"github.com/jnnkrdb/gokv/pkg/messaging"
	"github.com/jnnkrdb/gokv/pkg/server/tcpSocket"
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

		// sync struct contents
		var ssload = functions.SyncStorageLoad{
			Sync:   functions.SyncStorageLoad_Sync_Delete,
			Bucket: b,
		}

		// run synchro steps
		var tcpr = tcpSocket.TCPRequest{
			Initiator:    conf.SELF_NAME,
			RequestState: messaging.RS_Open,
			RequestCmd:   messaging.RC_SyncStorage,
		}

		if byt, err := json.Marshal(ssload); err != nil {
			log.Printf("[ERR] couldn't parse tcpr into bytes: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			tcpr.Load = byt
		}

		gossip.SpreadGossipTCP(tcpr)

		w.Write([]byte("OK"))
	}
}
