// ./models/zone.go (New File)
package models

import (
	"gorm.io/gorm"
	"time"
)

// Zone represents a territory or control area within a game instance.
type Zone struct {
	gorm.Model
	GameInstanceID        uint      `json:"game_instance_id"`
	Coordinates           string    `json:"coordinates"` // e.g., "x,y"
	ControlledByFactionID uint      `json:"controlled_by_faction_id,omitempty"`
	LastControlChange     time.Time `json:"last_control_change"`
}
