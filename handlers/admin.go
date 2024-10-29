// ./handlers/admin.go
package handlers

import (
	"encoding/json"
	"net/http"

	"drokkit/models"
	"time"
)

// CreateAdmin allows the creation of admin users. This should be restricted and possibly handled offline.
func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	var adminRequest struct {
		UserID      uint   `json:"user_id"`
		Permissions string `json:"permissions"` // Define permission levels as needed
	}

	if err := json.NewDecoder(r.Body).Decode(&adminRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	admin := models.Admin{
		UserID:      adminRequest.UserID,
		Permissions: adminRequest.Permissions,
		AssignedAt:  time.Now(),
	}

	if err := db.Create(&admin).Error; err != nil {
		http.Error(w, "Failed to create admin", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(admin)
}

// Example of a protected admin route
func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	// Verify if the requester has admin permissions
	// Extract user from JWT claims
	// Check permissions

	// Assuming permissions are verified
	var deleteRequest struct {
		PlayerID uint `json:"player_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&deleteRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := db.Delete(&models.Player{}, deleteRequest.PlayerID).Error; err != nil {
		http.Error(w, "Failed to delete player", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Player deleted successfully"})
}
