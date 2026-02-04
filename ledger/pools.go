package ledger

import (
	"math/big"
)

// Pool represents a pool of participation units to be distributed.
type Pool struct {
	Name    string   `json:"name"`    // The name of the pool
	Balance *big.Int `json:"balance"` // The current balance of the pool
	// TODO: Add other pool-related fields as necessary.
}
