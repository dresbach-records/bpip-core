package node

import (
	"bpip-core/storage"
	"bpip-core/validation"
)

// Node represents a BPIP node.
// It is responsible for processing events, validating them, and storing them.
type Node struct {
	validator validation.Validator
	store     storage.Store
}

// NewNode creates a new BPIP node.
func NewNode(validator validation.Validator, store storage.Store) *Node {
	return &Node{
		validator: validator,
		store:     store,
	}
}

// ProcessEvent processes an incoming event.
func (n *Node) ProcessEvent( /* TODO: pass event data */ ) error {
	// TODO: Implement the event processing logic.
	// 1. Unmarshal the event data.
	// 2. Validate the event.
	// 3. Store the event.
	// 4. Trigger economic calculations.
	return nil
}
