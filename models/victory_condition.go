package models

import (
	"gorm.io/gorm"
)

// VictoryCondition represents a specific condition that can lead to a game’s victory.
type VictoryCondition struct {
	gorm.Model
	GameInstanceID uint   `json:"game_instance_id"`
	Type           string `gorm:"type:enum('Domination','Economic','Military');not null" json:"type"`
	Details        string `json:"details"` // JSON-encoded thresholds and requirements
	IsMet          bool   `json:"is_met"`
}
