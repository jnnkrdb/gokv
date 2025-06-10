package messaging

// request type

type RequestCommand string

func (rc RequestCommand) String() string {
	return string(rc)
}

const (
	RC_SyncStorage RequestCommand = "req-syncstorage" // requests the node to synchronize the storage data
)
