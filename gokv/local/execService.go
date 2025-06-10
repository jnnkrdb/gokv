package local

import (
	"log"
	"time"

	"github.com/jnnkrdb/gokv/conf"
	"github.com/jnnkrdb/gokv/local/gossip"
	httpsocket "github.com/jnnkrdb/gokv/pkg/server/httpSocket"
	"github.com/jnnkrdb/gokv/pkg/server/tcpSocket"
)

func RunService() {

	// load the storage backend
	conf.LoadStorage()

	// start tcp socket
	go tcpSocket.RunTCPSocket(conf.GOSSIP_PORT, gossip.ReceiveGossip)

	// start gosspi spreading
	go gossip.SpreadGossip()

	// run http socket
	go httpsocket.RunHTTPSocket(conf.HTTP_PORT)

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
