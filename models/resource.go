package models

import (
	"gorm.io/gorm"
)

// Resource represents a resource within a game instance.
type Resource struct {
	gorm.Model
	GameInstanceID uint   `json:"game_instance_id"`
	PlayerID       uint   `json:"player_id,omitempty"` // Nullable for global resources
	Type           string `gorm:"type:enum('Gold','Wood','Stone', etc.);not null" json:"type"`
	Amount         int    `json:"amount"`
	LastUpdated    string `json:"last_updated"`
}
