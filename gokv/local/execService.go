package local

import (
	"log"
	"time"

	"github.com/jnnkrdb/gokv/conf"
	httpsocket "github.com/jnnkrdb/gokv/pkg/server/httpSocket"
	websocket "github.com/jnnkrdb/gokv/pkg/server/webSocket"
)

func RunService() {

	// load the storage backend
	conf.LoadStorage()

	// start tcp socket
	//go tcpSocket.RunTCPSocket(conf.GOSSIP_PORT, gossip.ReceiveGossip)

	go websocket.RunWS(conf.GOSSIP_PORT)

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
			}
		}
	}()

	// run http socket
	httpsocket.RunHTTPSocket(conf.HTTP_PORT)
}
