package app

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/dresbach-records/bpip-core/config"
	"github.com/dresbach-records/bpip-core/ledger"
	"github.com/dresbach-records/bpip-core/protocol"
	"github.com/dresbach-records/bpip-core/server"
	"github.com/google/uuid"
)

// App encapsulates the core application logic.
type App struct {
	config *config.Config
	ledger *ledger.Ledger
	logger *slog.Logger
	server *server.Server
}

// NewApp creates and returns a new App instance.
func NewApp(cfg *config.Config, ledger *ledger.Ledger, logger *slog.Logger) *App {
	app := &App{
		config: cfg,
		ledger: ledger,
		logger: logger,
	}
	app.server = server.NewServer(cfg.APIPort, app, logger)
	return app
}

// Run starts the main application logic, which is now the HTTP server.
func (a *App) Run() error {
	a.logger.Info("Starting the BPIP Core application server.")
	return a.server.Start()
}

// ProcessEvent validates, processes, and stores an event in the ledger.
// It also assigns a server-side ID and timestamp.
func (a *App) ProcessEvent(event *protocol.Event) error {
	// Assign a new ID and timestamp on the server-side
	event.ID = uuid.NewString()
	event.Timestamp = time.Now()

	a.logger.Info("Processing event", "type", event.Type, "id", event.ID)

	// Serialize the full event to JSON for storage
	eventBytes, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	// Store the event in the ledger
	if err := a.ledger.Set([]byte(event.ID), eventBytes); err != nil {
		return fmt.Errorf("failed to store event in ledger: %w", err)
	}

	a.logger.Info("Successfully stored event in ledger", "id", event.ID)
	return nil
}
