package ledger

import (
	"math/big"
)

// Participation represents a record of a user's participation in the BPIP economy.
type Participation struct {
	AccountID string   `json:"account_id"` // The account that participated
	Value     *big.Int `json:"value"`      // The value of the participation
	// TODO: Add other participation-related fields as necessary.
}
