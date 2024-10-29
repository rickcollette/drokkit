package models

import (
	"gorm.io/gorm"
)

// Stats represents player statistics, such as wins, losses, and experience.
type Stats struct {
	gorm.Model
	PlayerID    uint
	Wins        int `json:"wins"`
	Losses      int `json:"losses"`
	GamesPlayed int `json:"games_played"`
	Experience  int `json:"experience"`
}
