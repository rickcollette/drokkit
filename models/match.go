package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

// Match represents a game match between two players.
type Match struct {
	gorm.Model
	PlayerOne uint            `json:"player_one"`
	PlayerTwo uint            `json:"player_two"`
	GameState json.RawMessage `json:"game_state"` // JSON-encoded game state
	Turn      uint            `json:"turn"`
}
