// ./models/alliance.go (Updated)
package models

import (
	"gorm.io/gorm"
	"time"
)

// Alliance represents a formal alliance between factions within a game instance.
type Alliance struct {
	gorm.Model
	GameInstanceID  uint             `json:"game_instance_id"`
	Name            string           `json:"name"`
	CreatedAt       time.Time        `json:"created_at"`
	AllianceMembers []AllianceMember `json:"alliance_members"`
	AllianceChats   []AllianceChat   `json:"alliance_chats"`
}

// AllianceMember represents a faction within an alliance.
type AllianceMember struct {
	gorm.Model
	AllianceID uint      `json:"alliance_id"`
	FactionID  uint      `json:"faction_id"`
	JoinedAt   time.Time `json:"joined_at"`
}

// AllianceChat represents a chat message within an alliance.
type AllianceChat struct {
	gorm.Model
	AllianceID uint      `json:"alliance_id"`
	UserID     uint      `json:"user_id"`
	Message    string    `json:"message"`
	Timestamp  time.Time `json:"timestamp"`
}
