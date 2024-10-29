// ./handlers/faction.go (Ensure Faction references Zone correctly)
package handlers

import (
	"drokkit/models"
	"encoding/json"
	"net/http"
)

// CreateFaction allows a player to create a new faction within a game instance.
func CreateFaction(w http.ResponseWriter, r *http.Request) {
	var factionRequest struct {
		GameInstanceID uint   `json:"game_instance_id"`
		FactionType    string `json:"faction_type"`
		LeaderID       uint   `json:"leader_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&factionRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate FactionType
	validTypes := map[string]bool{
		"Industrialists": true,
		"Warriors":       true,
		"Technologists":  true,
		"Traders":        true,
	}

	if !validTypes[factionRequest.FactionType] {
		http.Error(w, "Invalid faction type", http.StatusBadRequest)
		return
	}

	// Create Faction
	faction := models.Faction{
		GameInstanceID: factionRequest.GameInstanceID,
		FactionType:    factionRequest.FactionType,
		LeaderID:       factionRequest.LeaderID,
	}

	// Set bonuses based on faction type
	switch factionRequest.FactionType {
	case "Industrialists":
		faction.ResourceBonus = 1.2
		faction.CombatBonus = 0.8
		faction.BuildingSpeed = 1.5
		faction.ResearchSpeed = 1.0
		faction.TradeRate = 1.0
		faction.DefenseStrength = 1.0
	case "Warriors":
		faction.ResourceBonus = 0.8
		faction.CombatBonus = 1.5
		faction.BuildingSpeed = 1.0
		faction.ResearchSpeed = 1.0
		faction.TradeRate = 1.0
		faction.DefenseStrength = 1.2
	case "Technologists":
		faction.ResourceBonus = 1.0
		faction.CombatBonus = 1.0
		faction.BuildingSpeed = 1.0
		faction.ResearchSpeed = 1.5
		faction.TradeRate = 1.0
		faction.DefenseStrength = 1.0
	case "Traders":
		faction.ResourceBonus = 1.0
		faction.CombatBonus = 0.8
		faction.BuildingSpeed = 1.0
		faction.ResearchSpeed = 1.0
		faction.TradeRate = 1.5
		faction.DefenseStrength = 0.8
	}

	if err := db.Create(&faction).Error; err != nil {
		http.Error(w, "Failed to create faction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(faction)
}
