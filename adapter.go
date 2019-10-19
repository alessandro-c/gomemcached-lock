package lock

import (
	"errors"
	"time"
)

var (
	// ErrAdptNotStored is returned when an Add operation failed to respect its condtions
	ErrAdptNotStored = errors.New("adapter: not stored")

	// ErrAdptNotFound is returned when attempting to Get a key that does not exist
	ErrAdptNotFound = errors.New("adapter: not found")
)

// Adapter defines the interface for memchached client adapters to be implemented
type Adapter interface {

	// Add will attempt to add a new item in memcached.
	// If the item already exists returns ErrAdptNotStored.
	Add(key string, value string, expiration time.Duration) error

	// Get an existing item from memcached
	// returns ErrAdptNotFound if the correspondent key does not exist
	Get(key string) (value string, err error)

	// Delete an existing item in memcached.
	// returns ErrAdptNotFound if the correspondent key does not exist
	Delete(key string) error
}
