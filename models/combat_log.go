package models

import (
	"gorm.io/gorm"
)

// CombatLog represents a combat event between two players or factions.
type CombatLog struct {
	gorm.Model
	GameInstanceID    uint          `json:"game_instance_id"`
	AttackerID        uint          `json:"attacker_id"`
	DefenderID        uint          `json:"defender_id"`
	UnitsLostAttacker int           `json:"units_lost_attacker"`
	UnitsLostDefender int           `json:"units_lost_defender"`
	Outcome           string        `gorm:"type:enum('Attacker Wins','Defender Wins','Draw');not null" json:"outcome"`
	Timestamp         string        `json:"timestamp"`
	CombatEvents      []CombatEvent `json:"combat_events"`
}

// CombatEvent represents a detailed step in a combat sequence.
type CombatEvent struct {
	gorm.Model
	CombatLogID uint   `json:"combat_log_id"`
	EventDetail string `json:"event_detail"` // JSON-encoded detailed event
}
