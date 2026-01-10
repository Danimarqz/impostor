package server

import (
	"impostor/internal/game"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Server contains the Fiber instance and the Game Hub.
type Server struct {
	App *fiber.App
	Hub *game.Hub
}

// NewServer initializes the web server and its dependencies.
func NewServer() *Server {
	app := fiber.New(fiber.Config{
		AppName: "The Impostor Agent",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Initialize Hub
	hub := game.NewHub()

	s := &Server{
		App: app,
		Hub: hub,
	}

	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	// Health check
	s.App.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// REST API Routes
	s.App.Post("/api/lobby", s.createLobbyHandler)
	s.App.Get("/api/categories", s.getCategoriesHandler)

	s.setupWebsocketRoutes()
}

// Run starts the server on the specified port.
func (s *Server) Run(port string) {
	log.Printf("Server listening on %s", port)
	log.Fatal(s.App.Listen(port))
}
