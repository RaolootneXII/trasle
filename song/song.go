package song

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/raolootnexii/trasle/database"
)

type Song struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

// GetSongs - get all songs
func GetSongs(c *fiber.Ctx) error {
	db := database.DBConn
	var songs []Song
	db.Find(&songs)
	return c.JSON(songs)
}
