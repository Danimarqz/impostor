package server

import (
	"impostor/internal/game"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) getCategoriesHandler(c *fiber.Ctx) error {
	language := c.Query("lang", "en") // Default to English
	names := game.GetAllCategoryNames(language)
	return c.JSON(fiber.Map{
		"categories": names,
	})
}
