// ./handlers/alliance.go (Ensure CreateAlliance uses correct models)
package handlers

import (
	"drokkit/models"
	"encoding/json"
	"net/http"
	"time"
)

// CreateAlliance allows two faction leaders to form an alliance.
func CreateAlliance(w http.ResponseWriter, r *http.Request) {
	var allianceRequest struct {
		GameInstanceID uint   `json:"game_instance_id"`
		Name           string `json:"name"`
		FactionIDs     []uint `json:"faction_ids"`
	}

	if err := json.NewDecoder(r.Body).Decode(&allianceRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if len(allianceRequest.FactionIDs) < 2 || len(allianceRequest.FactionIDs) > 2 {
		http.Error(w, "Alliance must consist of exactly two factions", http.StatusBadRequest)
		return
	}

	// Check if factions exist and belong to the same game instance
	var factions []models.Faction
	if err := db.Where("id IN ? AND game_instance_id = ?", allianceRequest.FactionIDs, allianceRequest.GameInstanceID).Find(&factions).Error; err != nil {
		http.Error(w, "Factions not found", http.StatusNotFound)
		return
	}

	if len(factions) != 2 {
		http.Error(w, "Factions not found or do not belong to the same game instance", http.StatusBadRequest)
		return
	}

	// Create Alliance
	alliance := models.Alliance{
		GameInstanceID: allianceRequest.GameInstanceID,
		Name:           allianceRequest.Name,
		CreatedAt:      time.Now(),
	}

	if err := db.Create(&alliance).Error; err != nil {
		http.Error(w, "Failed to create alliance", http.StatusInternalServerError)
		return
	}

	// Add members to the alliance
	for _, faction := range factions {
		member := models.AllianceMember{
			AllianceID: alliance.ID,
			FactionID:  faction.ID,
			JoinedAt:   time.Now(),
		}
		if err := db.Create(&member).Error; err != nil {
			http.Error(w, "Failed to add faction to alliance", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(alliance)
}
