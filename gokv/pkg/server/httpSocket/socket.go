package httpsocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	apiv1_storage "github.com/jnnkrdb/gokv/pkg/server/httpSocket/api/v1/storage"
)

var router *mux.Router = mux.NewRouter().StrictSlash(true)

func RunHTTPSocket(port int) {

	log.Printf("[INF] starting http socket on :%d\n", port)

	// handle the storage api
	router.Handle("/api/v1/storage/export", http.HandlerFunc(apiv1_storage.Export)).Methods("GET", "OPTIONS")
	router.Handle("/api/v1/storage/buckets", http.HandlerFunc(apiv1_storage.ListBuckets)).Methods("GET", "OPTIONS")
	router.Handle("/api/v1/storage/buckets/{bucket}/keys", http.HandlerFunc(apiv1_storage.ListKeys)).Methods("GET", "OPTIONS")
	router.Handle("/api/v1/storage/buckets/{bucket}/keys/{key}/value", http.HandlerFunc(apiv1_storage.GetKey)).Methods("GET", "OPTIONS")
	router.Handle("/api/v1/storage/buckets/{bucket}/keys/{key}/value", http.HandlerFunc(apiv1_storage.WrityKey)).Methods("POST", "PUT", "PATCH", "OPTIONS")
	router.Handle("/api/v1/storage/buckets/{bucket}/keys/{key}", http.HandlerFunc(apiv1_storage.DeleteKey)).Methods("DELETE", "OPTIONS")
	router.Handle("/api/v1/storage/buckets/{bucket}", http.HandlerFunc(apiv1_storage.DeleteBucket)).Methods("DELETE", "OPTIONS")

	// handle healthz api
	router.Handle("/healthz/live", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("OK")) }))
	router.Handle("/healthz/ready", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("OK")) }))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		log.Panic("[ERR] fatal error with httpSocket: %v\n", err)
	}
}
