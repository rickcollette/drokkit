// ./handlers/match.go (Ensure Match does not reference undefined types)
package handlers

import (
	"drokkit/models"
	"encoding/json"
	"net/http"
)

// GameState represents the state of the game within a match.
type GameState struct {
	TurnCount int          `json:"turn_count"`
	Moves     []PlayerMove `json:"moves"` // Record of moves in the game
}

// PlayerMove represents a single move by a player.
type PlayerMove struct {
	PlayerID uint   `json:"player_id"`
	Action   string `json:"action"`
}

// CreateMatch initiates a match between two players with an initial game state.
func CreateMatch(w http.ResponseWriter, r *http.Request) {
	var match models.Match
	if err := json.NewDecoder(r.Body).Decode(&match); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Initialize the game state
	initialGameState := GameState{
		TurnCount: 1,
		Moves:     []PlayerMove{},
	}

	// Serialize initial game state into JSON
	gameStateData, err := json.Marshal(initialGameState)
	if err != nil {
		http.Error(w, "Failed to create game state", http.StatusInternalServerError)
		return
	}
	match.GameState = gameStateData

	// Insert the match into the database
	if err := db.Create(&match).Error; err != nil {
		http.Error(w, "Failed to create match", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(match)
}

// PlayTurn allows a player to submit their turn, updates game state, and saves it.
func PlayTurn(w http.ResponseWriter, r *http.Request) {
	var turnData struct {
		MatchID  uint   `json:"match_id"`
		PlayerID uint   `json:"player_id"`
		Action   string `json:"action"`
	}

	if err := json.NewDecoder(r.Body).Decode(&turnData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Fetch the match from the database
	var match models.Match
	if err := db.First(&match, turnData.MatchID).Error; err != nil {
		http.Error(w, "Match not found", http.StatusNotFound)
		return
	}

	// Validate the player's turn
	if (match.Turn == 1 && match.PlayerOne != turnData.PlayerID) ||
		(match.Turn == 2 && match.PlayerTwo != turnData.PlayerID) {
		http.Error(w, "Not your turn", http.StatusForbidden)
		return
	}

	// Deserialize the existing game state
	var gameState GameState
	if err := json.Unmarshal(match.GameState, &gameState); err != nil {
		http.Error(w, "Failed to parse game state", http.StatusInternalServerError)
		return
	}

	// Update the game state with the new move
	newMove := PlayerMove{
		PlayerID: turnData.PlayerID,
		Action:   turnData.Action,
	}
	gameState.Moves = append(gameState.Moves, newMove)
	gameState.TurnCount++

	// Reserialize the updated game state
	updatedGameState, err := json.Marshal(gameState)
	if err != nil {
		http.Error(w, "Failed to update game state", http.StatusInternalServerError)
		return
	}
	match.GameState = updatedGameState

	// Toggle the turn between players
	match.Turn = 1 + (match.Turn % 2)

	// Save the updated match state
	if err := db.Save(&match).Error; err != nil {
		http.Error(w, "Failed to save match state", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(match)
}
