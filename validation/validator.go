package validation

import (
	"bpip-core/event"
)

// Validator is responsible for validating events.
// This is a critical component for ensuring the integrity of the BPIP system.
type Validator interface {
	Validate(e event.Event) (bool, error)
}
