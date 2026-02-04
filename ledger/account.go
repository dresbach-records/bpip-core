package ledger

import (
	"time"
)

// Account represents a user or entity in the BPIP system.
type Account struct {
	ID           string    `json:"id"`            // Unique identifier for the account
	CreationDate time.Time `json:"creation_date"` // Date the account was created
	// TODO: Add other account-related fields as necessary.
}
