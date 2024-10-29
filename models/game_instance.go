package models

import (
	"gorm.io/gorm"
)

// GameInstance represents an active game with its settings and state.
type GameInstance struct {
	gorm.Model
	LobbyID           uint               `json:"lobby_id"`
	Status            string             `gorm:"type:enum('Pending','Active','Completed');default:'Pending'" json:"status"`
	StartedAt         string             `json:"started_at"`
	EndedAt           string             `json:"ended_at"`
	RandomSeed        string             `json:"random_seed"`
	Players           []Player           `json:"players"`
	Factions          []Faction          `json:"factions"`
	Resources         []Resource         `json:"resources"`
	CombatLogs        []CombatLog        `json:"combat_logs"`
	Alliances         []Alliance         `json:"alliances"`
	VictoryConditions []VictoryCondition `json:"victory_conditions"`
}
