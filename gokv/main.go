package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jnnkrdb/gokv/conf"
	httpsocket "github.com/jnnkrdb/gokv/pkg/server/httpSocket"
	websocket "github.com/jnnkrdb/gokv/pkg/server/webSocket"
)

func main() {

	// activate logging
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	// print configs
	log.Printf("config:\n%v\n\n", *conf.NC)

	// load the storage backend
	conf.LoadStorage()

	// create the socket connection to existing nodes
	websocket.CreateWSConnections()

	// start the own websocket endpoint
	go websocket.RunWS()

	// running the test function printing the current storage
	go func() {
		if conf.NC.Debug {
			for {
				time.Sleep(30 * time.Second)

				// print the internal storage buckets and contents
				bucketList, err := conf.STORAGE.ListBuckets()
				if err != nil {
					log.Printf("[ERR] error receiving buckets from storage: %v\n", err)
					continue
				}
				for i := range bucketList {
					keyList, err := conf.STORAGE.ListKeys(bucketList[i])
					if err != nil {
						log.Printf("[ERR] error receiving keys from storage: %v\n", err)
						continue
					}

					log.Printf("\nBucket [%s]:\nKeys: %v", bucketList[i], keyList)
				}

				// request current nodes in cluster
				fmt.Printf("[INF] current nodes: %v\n", websocket.GetNodes())
			}
		}
	}()

	// run http socket
	httpsocket.RunHTTPSocket(conf.HTTP_PORT)
}
