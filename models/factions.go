// ./models/factions.go
package models

import (
	"gorm.io/gorm"
	"time"
)

// Faction represents a faction within a game instance.
type Faction struct {
	gorm.Model
	GameInstanceID  uint            `json:"game_instance_id"`
	FactionType     string          `gorm:"type:enum('Industrialists','Warriors','Technologists','Traders');not null" json:"faction_type"`
	LeaderID        uint            `json:"leader_id"`
	ResourceBonus   float64         `json:"resource_bonus"`
	CombatBonus     float64         `json:"combat_bonus"`
	BuildingSpeed   float64         `json:"building_speed"`
	ResearchSpeed   float64         `json:"research_speed"`
	TradeRate       float64         `json:"trade_rate"`
	DefenseStrength float64         `json:"defense_strength"`
	FactionMembers  []FactionMember `json:"faction_members"`
	ControlledZones []Zone          `json:"controlled_zones"`
}

// FactionMember represents a player within a faction.
type FactionMember struct {
	gorm.Model
	FactionID uint      `json:"faction_id"`
	PlayerID  uint      `json:"player_id"`
	JoinedAt  time.Time `json:"joined_at"`
}
