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

func (s *Server) getRandomWordHandler(c *fiber.Ctx) error {
	category := c.Query("category", "General")
	lang := c.Query("lang", "en")

	word, err := game.GetRandomWord(category, lang)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get random word",
		})
	}

	return c.JSON(word)
}
