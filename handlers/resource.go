package handlers

import (
	"drokkit/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

// UpdateResource allows updating a player's resources.
func UpdateResource(w http.ResponseWriter, r *http.Request) {
	var resourceUpdate struct {
		GameInstanceID uint   `json:"game_instance_id"`
		PlayerID       uint   `json:"player_id"`
		Type           string `json:"type"`
		Amount         int    `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&resourceUpdate); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Find existing resource or create a new one
	var resource models.Resource
	result := db.Where("game_instance_id = ? AND player_id = ? AND type = ?", resourceUpdate.GameInstanceID, resourceUpdate.PlayerID, resourceUpdate.Type).First(&resource)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			resource = models.Resource{
				GameInstanceID: resourceUpdate.GameInstanceID,
				PlayerID:       resourceUpdate.PlayerID,
				Type:           resourceUpdate.Type,
				Amount:         resourceUpdate.Amount,
			}
			if err := db.Create(&resource).Error; err != nil {
				http.Error(w, "Failed to create resource", http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
	} else {
		// Update existing resource
		resource.Amount += resourceUpdate.Amount
		resource.LastUpdated = "now" // Replace with actual timestamp
		if err := db.Save(&resource).Error; err != nil {
			http.Error(w, "Failed to update resource", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resource)
}
