package server

import (
	"encoding/json"
	"impostor/internal/domain"
	"impostor/internal/game"
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

		// 1. Get Lobby (Strict Mode: Must exist)
		lobby, ok := s.Hub.GetLobby(lobbyID)
		if !ok {
			// For strictly generated UUIDs, we should fail if not found.
			// sending close message or just closing.
			log.Printf("Lobby %s not found", lobbyID)
			c.Close()
			return
		}

		// Determine if this player should be the leader (if lobby has no players yet)
		// We need to check this safely.
		// Since we can't lock the lobby from here easily without exposing Mutex,
		// we should rely on AddPlayer logic or check safe copy.
		// BETTER: Update Hub/Lobby `AddPlayer` to return "isLeader" status or handle it internally.
		// For this MVP, let's assume if count == 0, isLeader = true.
		// But AddPlayer is void.
		// Let's modify logic: Just create the player struct here.
		// We need to know if they are the FIRST one.
		// HACK: We can check len(lobby.Players) but it's racy without lock. 
		// Since AddPlayer locks, we should move Leader assignment INSIDE AddPlayer?
		// or make AddPlayer smart.
		
		// Let's stick to safe logic using the public methods we have.
		// We can't change domain.Player inside AddPlayer easily if it's passed by pointer? 
		// Yes we can.
		
		// Let's trust the "CreateLobby" from API created an empty lobby.
		// But wait, `Hub.GetLobby` uses RLock.
		// If 2 players join simultaneously, they both might see count 0?
		// No, `AddPlayer` has a Write Lock.
		
		// Let's refine Handshake:
		p := &domain.Player{
			ID:       playerID,
			Name:     playerName,
			IsAlive:  true,
			IsLeader: false, // Default
		}
		
		isFirst := lobby.AddPlayerSafe(p) // We need this new method to be atomic
		if isFirst {
			p.IsLeader = true
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
		
		// Broadcast updated player list (includes the new player)
		// This ensures the new player gets the full list of existing players,
		// and existing players see the new one.
		lobby.BroadcastPlayerList()

		defer func() {
			// Cleanup on disconnect
			isEmpty := lobby.RemovePlayer(playerID)
			if isEmpty {
				log.Printf("Lobby %s empty, deleting...", lobbyID)
				s.Hub.DeleteLobby(lobbyID) // We need to expose this or use safe method
			} else {
				// Broadcast player left event to updating remaining clients
				// We need to fetch the updated list or just send the ID of who left.
				// Ideally, we also send the NEW leader if changed.
				// Let's enable a broadcast of the full player list to be safe.
				// This requires a helper "BroadcastPlayerList".
				// For now:
				leaveMsg := map[string]interface{}{
					"type": "PLAYER_LEFT",
					"player_id": playerID,
				}
				lobby.Broadcast(leaveMsg)
				
				// Also send an update for the new leader? 
				// The client can deduce it if we send the full list.
				// Let's modify RemovePlayer to return the new leader ID or boolean?
				// Simpler: Just force a full refresh of players.
				// Since we can't easily get the list safely here without another lock,
				// let's rely on the frontend to remove the player ID from its list,
				// AND check if the leader is gone.
				
				// Actually, `RemovePlayer` logic assigned a new leader.
				// We should probably broadcast the FULL list to ensure sync.
				// Let's implement `BroadcastPlayerList` on Lobby.
				lobby.BroadcastPlayerList() 
			}
		}()

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
					// Parse Start Options
					type StartPayload struct {
						Mode     string `json:"mode"`     // "easy" or "hard"
						Category string `json:"category"` // Category name
						Language string `json:"language"` // "en" or "es"
					}
					var startOpts StartPayload
					json.Unmarshal(msg, &startOpts)

					// Default to Hard (Standard)
					gameMode := domain.ModeHard
					if startOpts.Mode == "easy" {
						gameMode = domain.ModeEasy
					}

					// Default language to English
					language := startOpts.Language
					if language == "" {
						language = "en"
					}

					// Store language in lobby config
					lobby.Config.Language = language

					// Get selected category or default
					cat := game.GetCategoryByName(startOpts.Category, language)

					// Note: StartGame signature needs to accept GameMode if it doesn't already.
					// Reviewing match.go: func (l *Lobby) StartGame(category domain.Category) error
					// We need to update StartGame to take the mode!
					if err := lobby.StartGame(cat, gameMode); err != nil {
						log.Printf("Error starting game: %v", err)
					} else {
						log.Println("Game Started and broadcasted!")
					}
				} else if cmd.Action == "CHAT_MESSAGE" {
					// Broadcast chat message
					// Payload expected: { action: "CHAT_MESSAGE", message: "text" }
					type ChatPayload struct {
						Message string `json:"message"`
					}
					var chatPayload ChatPayload
					json.Unmarshal(msg, &chatPayload)

					chatMsg := map[string]interface{}{
						"type": "CHAT_MESSAGE",
						"from": playerName,
						"text": chatPayload.Message,
					}
					lobby.Broadcast(chatMsg)
				} else if cmd.Action == "CAST_VOTE" {
					type VotePayload struct {
						TargetID string `json:"target_id"`
					}
					var votePayload VotePayload
					json.Unmarshal(msg, &votePayload)
					
					lobby.CastVote(playerID, votePayload.TargetID)
				} else if cmd.Action == "RESET_GAME" {
					lobby.ResetGame()
				}
			}
		}
	}))
}
