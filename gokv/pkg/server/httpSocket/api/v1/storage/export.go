package apiv1_storage

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jnnkrdb/gokv/conf"
)

func Export(w http.ResponseWriter, r *http.Request) {

	if err := json.NewEncoder(w).Encode(conf.STORAGE); err != nil {

		log.Printf("[ERR] couldn't parse list of buckets into json: %v\n", err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
