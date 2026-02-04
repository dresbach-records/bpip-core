package ledger

import (
	"math/big"
)

// Distribution represents the distribution of participation units to accounts.
type Distribution struct {
	AccountID string   `json:"account_id"` // The account receiving the distribution
	Amount    *big.Int `json:"amount"`      // The amount of participation units distributed
	// TODO: Add other distribution-related fields as necessary.
}
