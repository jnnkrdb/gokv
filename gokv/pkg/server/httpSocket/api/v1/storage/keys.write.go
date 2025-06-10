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

		// sync struct contents
		var ssload = functions.SyncStorageLoad{
			Sync:   functions.SyncStorageLoad_Sync_Write,
			Bucket: b,
			Key:    k,
			Value:  content,
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
