package apiv1_storage

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jnnkrdb/gokv/conf"
)

func ListBuckets(w http.ResponseWriter, r *http.Request) {

	if buckets, err := conf.STORAGE.ListBuckets(); err != nil {

		log.Printf("[ERR] couldn't receive list of buckets: %v\n", err)

		http.Error(w, err.Error(), http.StatusInternalServerError)

	} else {

		if err := json.NewEncoder(w).Encode(buckets); err != nil {

			log.Printf("[ERR] couldn't parse list of buckets into json: %v\n", err)

			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
