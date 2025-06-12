package functions

import (
	"log"

	"github.com/jnnkrdb/gokv/pkg/messaging"
	"github.com/jnnkrdb/gokv/pkg/server/tcpSocket"
)

type SyncStorageLoad struct {
	Sync   uint   `json:"s"`
	Bucket string `json:"b"`
	Key    string `json:"k"`
	Value  string `json:"v"`
}

const (
	SyncStorageLoad_Sync_Write  uint = 1
	SyncStorageLoad_Sync_Delete uint = 2
)

func SyncStorage(rsp *tcpSocket.TCPResponse) {

	log.Printf("[WRN] method not implemented\n")

	rsp.RequestState = messaging.RS_Warning

	rsp.Load = []byte("the syncStorage method is not implemented")
}
