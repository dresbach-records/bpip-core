package economy

import (
	"bpip-core/ledger"
)

// Redistribution moves participation units from one pool to another, or to accounts.
type Redistribution struct {
	From   *ledger.Pool   `json:"from"`   // The source of the redistribution
	To     []*ledger.Pool `json:"to"`     // The destination(s) of the redistribution
	// TODO: Implement the redistribution logic as defined in the BPIP economy documentation.
}
