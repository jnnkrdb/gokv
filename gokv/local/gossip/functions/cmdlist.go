package functions

import (
	"log"

	"github.com/jnnkrdb/gokv/pkg/messaging"
	"github.com/jnnkrdb/gokv/pkg/server/tcpSocket"
)

func Handle(req tcpSocket.TCPRequest, rsp *tcpSocket.TCPResponse) {

	log.Printf("[INF] received request from [%s]: [%s]\n", req.Initiator, req.RequestCmd.String())

	switch req.RequestCmd {

	case messaging.RC_SyncStorage: // --------------------------------- sync storage
		SyncStorage(rsp)
		break

	default: // ------------------------------------------------------- default
		log.Printf("[WRN] received unallowed request command\n")
		rsp.RequestState = messaging.RS_Error
		rsp.Load = []byte(`{"msg":"this request command is not allowed."}`)
		break
	}
}
