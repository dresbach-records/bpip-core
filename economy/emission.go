package economy

import (
	"math/big"
)

// Emission represents the creation of new participation units.
// Emission is a function of the value entering the system.
type Emission struct {
	Amount *big.Int `json:"amount"` // The amount of participation units emitted
	// TODO: Implement the emission logic as defined in the BPIP economy documentation.
}
