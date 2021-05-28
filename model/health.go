package model

import "errors"

// EventGetHealth evento de health check via nats
const EventGetHealth = "EVENT_GET_HEALTH"

// Health modelo padrão para o health check
type Health struct {
	Version        string `json:"version"`
	ServerStatedAt string `json:"server_started_at"`
	DatabaseStatus string `json:"database_status" db:"database_status"`
}

// ToHealth converte uma interface{} para *Health
func ToHealth(data interface{}) (*Health, error) {
	value, ok := data.(*Health)
	if !ok {
		return nil, errors.New("não foi possível converter interface{} para *Health")
	}
	return value, nil
}
