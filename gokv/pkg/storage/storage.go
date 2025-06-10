package storage

// this is the typical interface to describe the structure
// of a storage implementation
type Storage interface {

	// receive a specific key-value from a bucket
	Get(string, string) (string, error)

	// write a defined key-value into a given bucket
	Write(string, string, string) error

	// delete a specific key-value from a bucket
	Delete(string, string) error
}
