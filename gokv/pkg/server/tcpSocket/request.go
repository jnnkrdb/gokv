package tcpSocket

import (
	"github.com/jnnkrdb/gokv/pkg/messaging"
)

type TCPRequest struct {

	// the initiator is the one starting the command chain
	Initiator string `json:"initiator"`

	// the request type and code determine the overall request
	// like the job identification and the current step or state
	// the job has

	RequestState messaging.RequestState   `json:"rstate"`
	RequestCmd   messaging.RequestCommand `json:"rcmd"`

	// can be already parsed json, depends on the actual request
	// that was ordered
	Load []byte `json:"load"`
}
