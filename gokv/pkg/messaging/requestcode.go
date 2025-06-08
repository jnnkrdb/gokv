package messaging

// answer codes

type RequestCode string

func (rc RequestCode) String() string {
	return string(rc)
}

const (

	// Marks the beginning of a request chain. Contains the first contact informations
	// and the actual request.
	RC_Open = "rc-open"

	// This is the default answer for the health check request, since no further infos
	// are necessary. Marks a request as finished.
	RC_Ok = "rc-ok"

	// With this code, the responding node declares, that it waits for further
	// actions from the sender or the node itself.
	RC_Waiting = "rc-waiting"

	// This code says, that the requested process was finished successfully.
	RC_Success = "rc-success"

	// This code says, that a request was received and probably processed, but there is
	// a possibility, that things went wrong while processing.
	RC_Warning = "rc-warning"

	// An error occured, further informations can be gathered from additional load.
	// This code says, that the request was dropped from the receiving node.
	RC_Error = "rc-error"

	// A critical error occured, there won't be any further infos on the payload.
	// Mostly happens, as an answer to requests, which could not be parsed from the
	// receiving node.
	RC_Critical = "rc-critical"
)
