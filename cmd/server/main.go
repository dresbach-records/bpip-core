package main

import (
	"log"

	"bpip-core/node"
	"bpip-core/storage"
	"bpip-core/validation"
)

func main() {
	// Create a new BadgerStore.
	store, err := storage.NewBadgerStore("./badger")
	if err != nil {
		log.Fatalf("failed to create badger store: %v", err)
	}

	// Create a new validator.
	validator := validation.NewBasicValidator() // Assumes a basic validator is implemented

	// Create a new BPIP node.
	n := node.NewNode(validator, store)

	// TODO: Start a server (e.g., HTTP, gRPC) to listen for incoming events.
	// For example:
	// http.HandleFunc("/event", handleEvent(n))
	// log.Fatal(http.ListenAndServe(":8080", nil))

	_ = n // a temporary fix to the unused variable error
}

// handleEvent would be a function that handles incoming HTTP requests and calls n.ProcessEvent.
// func handleEvent(n *node.Node) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// ... handle request and call n.ProcessEvent(...)
// 	}
// }
