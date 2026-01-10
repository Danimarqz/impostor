package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// createLobbyHandler creates a new lobby with a backend-generated UUID.
func (s *Server) createLobbyHandler(c *fiber.Ctx) error {
	id := uuid.New().String()

	// Create empty lobby. The first player to connect will become the leader.
	// Hub.CreateLobby is thread-safe.
	// We pass nil as host because the host hasn't connected via WS yet.
	_ = s.Hub.CreateLobby(id, nil)

	return c.JSON(fiber.Map{
		"lobby_id": id,
	})
}
