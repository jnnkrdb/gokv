package httpsocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router = mux.NewRouter().StrictSlash(true)

func RunHTTPSocket(port int) {

	log.Printf("[INF] starting http socket on :%d\n", port)

	router.Handle("/healthz/live", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("OK")) }))
	router.Handle("/healthz/ready", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("OK")) }))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		log.Panic("[ERR] fatal error with httpSocket: %v\n", err)
	}
}
