package tcpSocket

import "github.com/jnnkrdb/gokv/pkg/messaging"

type TCPRequest struct {

	// the request type and code determine the overall request
	// like the job identification and the current step or state
	// the job has

	RequestType messaging.RequestType `json:"rtype"`
	RequestCode messaging.RequestCode `json:"rcode"`

	// can be already parsed json, depends on the actual request
	// that was ordered
	Load string `json:"load"`
}
