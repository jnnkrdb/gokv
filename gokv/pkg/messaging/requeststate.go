package messaging

// answer codes

type RequestState string

func (rc RequestState) String() string {
	return string(rc)
}

const (

	// Marks the beginning of a request chain. Contains the first contact information
	// and the actual request.
	RS_Open RequestState = "rs-open"

	// This is the default answer for requests, which do not need further infos.
	// Marks a request as finished.
	RS_Ok RequestState = "rs-ok"

	// With this code, the responding node declares, that it waits for further
	// actions from the sender or the node itself.
	RS_Waiting RequestState = "rs-waiting"

	// This code says, that the requested process was finished successfully.
	RS_Success RequestState = "rs-success"

	// This code says, that a request was received and probably processed, but there is
	// a possibility, that things went wrong while processing.
	RS_Warning RequestState = "rs-warning"

	// An error occured, further informations can be gathered from additional load.
	// This code says, that the request was dropped from the receiving node.
	RS_Error RequestState = "rs-error"

	// A critical error occured, there won't be any further infos on the payload.
	// Mostly happens, as an answer to requests, which could not be parsed from the
	// receiving node.
	RS_Critical RequestState = "rs-critical"
)
