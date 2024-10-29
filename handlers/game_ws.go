package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Consider tightening this in production
	},
}

// PlayerConnection represents an active WebSocket connection for a player
type PlayerConnection struct {
	Conn     *websocket.Conn
	PlayerID uint
}

var (
	playerConnections = make(map[uint]*PlayerConnection)
	connectionsMutex  sync.Mutex
)

// WebSocketHandler establishes a WebSocket connection and manages turns
func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		http.Error(w, "Failed to upgrade WebSocket", http.StatusInternalServerError)
		return
	}

	// Extract and validate JWT from query parameters
	tokenStr := r.URL.Query().Get("token")
	if tokenStr == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	playerID := claims.UserID // Use a valid field that represents the player ID in `Claims`
	log.Printf("Player %d connected", playerID)

	// Store the player connection
	connectionsMutex.Lock()
	playerConnections[playerID] = &PlayerConnection{Conn: conn, PlayerID: playerID}
	connectionsMutex.Unlock()

	defer func() {
		conn.Close()
		connectionsMutex.Lock()
		delete(playerConnections, playerID)
		connectionsMutex.Unlock()
		log.Printf("Player %d disconnected", playerID)
	}()

	// Listen for moves from this player
	for {
		var move PlayerMove
		err := conn.ReadJSON(&move)
		if err != nil {
			log.Printf("Error reading JSON from player %d: %v", playerID, err)
			break
		}
		handlePlayerMove(playerID, move)
	}
}

func handlePlayerMove(playerID uint, move PlayerMove) {
	moveJSON, err := json.Marshal(move)
	if err != nil {
		log.Printf("Failed to marshal move: %v", err)
		return
	}

	// Broadcast move to other players in the match
	connectionsMutex.Lock()
	for _, pc := range playerConnections {
		if pc.PlayerID != playerID {
			err := pc.Conn.WriteMessage(websocket.TextMessage, moveJSON)
			if err != nil {
				log.Printf("Failed to send move to player %d: %v", pc.PlayerID, err)
			}
		}
	}
	connectionsMutex.Unlock()
}
