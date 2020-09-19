package song

import (
	"github.com/gofiber/fiber/v2"
)

// GetSongs - get all songs
func GetSongs(c *fiber.Ctx) error {
	return c.SendString("All Songs")
}
