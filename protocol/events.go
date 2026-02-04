package protocol

import (
	"encoding/json"
	"time"
)

// EventType é um enum para os tipos de eventos do protocolo.
type EventType string

const (
	EventTypeAssetCreation     EventType = "ASSET_CREATION"
	EventTypeContribution      EventType = "CONTRIBUTION"
	EventTypeDistribution      EventType = "DISTRIBUTION"
)

// EventHeader define os metadados comuns a todos os eventos.
type EventHeader struct {
	ID        string    `json:"id"`
	Type      EventType `json:"type"`
	Timestamp time.Time `json:"timestamp"`
}

// Event é a estrutura genérica para todos os eventos do protocolo.
// O Payload é um RawMessage para permitir a decodificação em tipos específicos.
type Event struct {
	EventHeader
	Payload json.RawMessage `json:"payload"`
}

// AssetCreationPayload é o payload para um evento de criação de ativo.
type AssetCreationPayload struct {
	AssetID     string `json:"asset_id"`
	Description string `json:"description"`
	Creator     string `json:"creator"`
}

// ContributionPayload é o payload para um evento de contribuição.
type ContributionPayload struct {
	AssetID       string  `json:"asset_id"`
	Contributor   string  `json:"contributor"`
	Contribution  string  `json:"contribution"`
	Hours         float64 `json:"hours"`
}
