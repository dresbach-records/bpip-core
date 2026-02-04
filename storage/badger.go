package storage

import (
	"bpip-core/event"

	"github.com/dgraph-io/badger/v3"
)

// BadgerStore is a simple implementation of the Store interface using BadgerDB.
// BadgerDB is an embeddable, persistent, and fast key-value store.
type BadgerStore struct {
	db *badger.DB
}

// NewBadgerStore creates a new BadgerStore.
func NewBadgerStore(path string) (*BadgerStore, error) {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}
	return &BadgerStore{db: db}, nil
}

// Get retrieves an event from the BadgerStore.
func (s *BadgerStore) Get(id string) (event.Event, error) {
	// TODO: Implement the Get method.
	return event.Event{}, nil
}

// Put stores an event in the BadgerStore.
func (s *BadgerStore) Put(e event.Event) error {
	// TODO: Implement the Put method.
	return nil
}
