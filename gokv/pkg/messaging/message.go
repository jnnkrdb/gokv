package messaging

type WsRequest struct {
	// the initiator is the one starting the command chain
	Initiator string `json:"initiator"`

	// the uid creates a chain to identify a specific request
	RequestUID string `json:"reqid"`

	// the reqcommand contains the actual command that should
	// be executed
	RequestCommand RequestCommand `json:"reqc"`

	// [OPTIONAL] contains additional load for the command, if required
	Load []byte `json:"load"`
}
