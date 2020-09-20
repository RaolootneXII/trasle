package song

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/raolootnexii/trasle/database"
)

type Song struct {
	gorm.Model
	Title  string `json:"title";gorm:"not null"`
	Lyrics string `json:"lyrics";gorm:"not null"`
	Author string `json:"author";gorm:"not null"`
	Rating int    `json:"rating";gorm:"not null"`
}

// GetSongs - get all songs
func GetSongs(c *fiber.Ctx) error {
	db := database.DBConn
	var songs []Song
	db.Find(&songs)
	return c.JSON(songs)
}

// CreateSong - create a song
func CreateSong(c *fiber.Ctx) error {
	db := database.DBConn
	song := new(Song)

	if err := c.BodyParser(song); err != nil {
		return err
	}

	db.Create(&song)
	c.JSON(song)
	return c.SendStatus(201)
}
