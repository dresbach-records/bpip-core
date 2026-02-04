package ledger

import (
	"fmt"
	"os"

	"github.com/dgraph-io/badger/v3"
)

// Ledger provides an interface to the underlying key-value store (BadgerDB).
type Ledger struct {
	db *badger.DB
}

// NewLedger initializes and returns a new Ledger instance.
// It takes the database path from the configuration.
func NewLedger(path string) (*Ledger, error) {
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, fmt.Errorf("failed to create ledger directory: %w", err)
	}

	opts := badger.DefaultOptions(path)
	opts.Logger = nil // Disable verbose logging from Badger

	db, err := badger.Open(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to open ledger database: %w", err)
	}

	return &Ledger{db: db}, nil
}

// Close gracefully closes the database connection.
func (l *Ledger) Close() error {
	return l.db.Close()
}

// Set stores a key-value pair in the ledger.
func (l *Ledger) Set(key, value []byte) error {
	err := l.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
	return err
}

// Get retrieves a value by its key from the ledger.
func (l *Ledger) Get(key []byte) ([]byte, error) {
	var value []byte
	err := l.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			value = append([]byte{}, val...)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return value, nil
}
