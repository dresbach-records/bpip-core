package server

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/dresbach-records/bpip-core/protocol"
)

// AppProcessor defines the interface for processing events.
// This allows us to decouple the server from the main application logic.
type AppProcessor interface {
	ProcessEvent(event *protocol.Event) error
}

// Server encapsulates the HTTP server logic.
type Server struct {
	port    string
	app     AppProcessor
	logger  *slog.Logger
	handler http.Handler
}

// NewServer creates and returns a new Server instance.
func NewServer(port string, app AppProcessor, logger *slog.Logger) *Server {
	srv := &Server{
		port:   port,
		app:    app,
		logger: logger,
	}

	router := http.NewServeMux()
	router.HandleFunc("/events", srv.eventsHandler)
	srv.handler = router

	return srv
}

// Start runs the HTTP server.
func (s *Server) Start() error {
	s.logger.Info("Starting HTTP server", "port", s.port)
	return http.ListenAndServe(":"+s.port, s.handler)
}

// eventsHandler handles incoming event submissions.
func (s *Server) eventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var event protocol.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		s.logger.Error("Failed to decode event", "error", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := s.app.ProcessEvent(&event); err != nil {
		s.logger.Error("Failed to process event", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "event processed", "id": event.ID})
}
