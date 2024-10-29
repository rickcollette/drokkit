package models

import (
	"gorm.io/gorm"
)

// Player represents a user in the game, including their credentials and statistics.
type Player struct {
	gorm.Model
	Username string          `gorm:"unique" json:"username"`
	Password string          `json:"password"`
	Stats    Stats           `json:"stats"`
	Matches  []Match         `json:"matches"`
	Factions []FactionMember `json:"factions"`
}
