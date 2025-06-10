package apiv1_storage

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jnnkrdb/gokv/conf"
)

func ListKeys(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	b, ok := vars["bucket"]
	if !ok {
		log.Printf("[WRN] bucket var is missing\n")
		http.Error(w, "bucketVar is missing", http.StatusBadRequest)
		return
	}

	if keys, err := conf.STORAGE.ListKeys(b); err != nil {

		log.Printf("[ERR] couldn't receive list of buckets: %v\n", err)

		http.Error(w, err.Error(), http.StatusInternalServerError)

	} else {

		if err := json.NewEncoder(w).Encode(keys); err != nil {

			log.Printf("[ERR] couldn't parse list of buckets into json: %v\n", err)

			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
