package event

import (
	"time"
)
// Event represents a generic event in the BPIP system.
// All specific event types embed this struct.
type Event struct {
	ID        string      `json:"id"`         // Unique identifier for the event
	Timestamp time.Time   `json:"timestamp"` // Time the event occurred
	Signature []byte      `json:"signature"`  // Cryptographic signature of the event data
	Metadata  interface{} `json:"metadata"`   // Economic metadata associated with the event
}
