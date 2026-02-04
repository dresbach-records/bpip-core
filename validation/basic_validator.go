package validation

import (
	"bpip-core/event"
	"fmt"
)

// BasicValidator is a simple implementation of the Validator interface.
// It performs basic checks on events.
type BasicValidator struct{}

// NewBasicValidator creates a new BasicValidator.
func NewBasicValidator() *BasicValidator {
	return &BasicValidator{}
}

// Validate performs basic validation on an event.
func (v *BasicValidator) Validate(e event.Event) (bool, error) {
	if e.ID == "" {
		return false, fmt.Errorf("event ID is missing")
	}
	if e.Timestamp.IsZero() {
		return false, fmt.Errorf("event timestamp is missing")
	}
	// TODO: Add more specific validation rules for different event types.
	return true, nil
}
