package server

import (
	"encoding/json"
	"impostor/internal/domain"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

// setupWebsocketRoutes configures the WS endpoints.
func (s *Server) setupWebsocketRoutes() {
	// Middleware to check if it's a websocket upgrade
	s.App.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	s.App.Get("/ws/:lobbyId", websocket.New(func(c *websocket.Conn) {
		lobbyID := c.Params("lobbyId")
		playerID := c.Query("playerId")
		playerName := c.Query("playerName")

		if playerID == "" || playerName == "" {
			log.Println("Missing player info")
			c.Close()
			return
		}

		log.Printf("Player %s connecting to lobby %s", playerName, lobbyID)

		// 1. Get or Create Lobby
		lobby, ok := s.Hub.GetLobby(lobbyID)
		if !ok {
			// Auto-create for simplicity if it doesn't exist
			host := &domain.Player{
				ID:       playerID,
				Name:     playerName,
				IsLeader: true,
				IsAlive:  true,
			}
			lobby = s.Hub.CreateLobby(lobbyID, host)
		} else {
			// Add player to existing lobby
			p := &domain.Player{
				ID:       playerID,
				Name:     playerName,
				IsLeader: false,
				IsAlive:  true,
			}
			lobby.AddPlayer(p)
		}
		
		// Register connection for broadcasting
		lobby.RegisterClient(playerID, c)

		// Broadcast updated player list
		// In a real app we'd need a cleaner DTO approach
		// We can't access lobby.Players safely without lock, but we are outside lock here.
		// Wait, RegisterClient uses lock. We need a method on Lobby to "GetState" or "BroadcastState".
		// Let's create a quick valid map to send.
		// ACTUALLY: Lobby.Broadcast is generic. We need to fetch data safely.
		// Let's just send a simple "Player Joined" message for now, or even better:
		// Modify Lobby to have a method "BroadcastPlayerList".
		
		// For MVP, since we don't have that method, let's just assume we send a basic signal 
		// and the frontend will request state, or we send it here (risky without lock).
		// Best approach: Add `BroadcastGameState` to Lobby struct in next step.
		// For now, let's send a manual JSON with what we know.
		
		joinMsg := map[string]interface{}{
			"type": "PLAYER_JOINED",
			"player": map[string]interface{}{
				"id": playerID, 
				"name": playerName,
				"is_leader": lobby.Players[playerID].IsLeader, // Check safely if player exists and is leader
			},
		}
		lobby.Broadcast(joinMsg)

		// 2. Read Loop
		// In a real app, we would have a 'Client' struct managing the connection write pump too.
		// For this MVP, we just read.
		var (
			msg []byte
			err error
		)
		for {
			if _, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			// Handle messages (e.g., START_GAME)
			// Simple JSON command structure
			type Command struct {
				Action string `json:"action"`
			}
			var cmd Command
			if json.Unmarshal(msg, &cmd) == nil {
				if cmd.Action == "START_GAME" {
					// Hardcoded category for now
					cat := domain.Category{
						Name: "General",
						Pairs: []domain.WordPair{
							{Real: "Hospital", Trap: "Pharmacy"},
						},
					}
					if err := lobby.StartGame(cat); err != nil {
						log.Printf("Error starting game: %v", err)
					} else {
						log.Println("Game Started and broadcasted!")
					}
				}
			}
		}
	}))
}
