package storage

import (
	"bpip-core/event"
)

// Store is an interface for storing and retrieving events.
// This allows for different storage backends to be used.
type Store interface {
	Get(id string) (event.Event, error)
	Put(e event.Event) error
}
