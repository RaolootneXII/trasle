package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raolootnexii/trasle/database"
	"github.com/raolootnexii/trasle/model"
)

// GetSongs - get all songs
func GetSongs(c *fiber.Ctx) error {
	db := database.DBConn
	var songs []model.Song
	db.Find(&songs)
	return c.JSON(songs)
}

// GetSong - get a song
func GetSong(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	song := new(model.Song)
	db.First(&song, id)
	if song.ID == 0 {
		c.JSON(fiber.Map{
			"message": "not found",
		})
		return c.SendStatus(404)
	}
	return c.JSON(song)
}

// CreateSong - create a song
func CreateSong(c *fiber.Ctx) error {
	db := database.DBConn
	song := new(model.Song)

	if err := c.BodyParser(song); err != nil {
		return err
	}

	db.Create(&song)
	c.JSON(song)
	return c.SendStatus(201)
}

// DeleteSong - delete a song
func DeleteSong(c *fiber.Ctx) error {
	if c.Params("id") != "" {
		id := c.Params("id")
		db := database.DBConn
		song := new(model.Song)
		db.Delete(&song, id)
		c.JSON(fiber.Map{
			"message": "deleted",
		})
		return c.SendStatus(204)
	}
	c.JSON(fiber.Map{
		"message": "not found",
	})
	return c.SendStatus(404)
}

// UpdateSong - update a song
func UpdateSong(c *fiber.Ctx) error {
	if c.Params("id") != "" {
		id := c.Params("id")
		db := database.DBConn
		song := new(model.Song)
		db.First(&song, id)
		if song.ID == 0 {
			c.JSON(fiber.Map{
				"message": "not found",
			})
			return c.SendStatus(404)
		}
		if err := c.BodyParser(song); err != nil {
			return err
		}
		db.Save(&song)
		c.JSON(fiber.Map{
			"message": "updated",
		})
		return c.SendStatus(200)
	}
	c.JSON(fiber.Map{
		"message": "id not is valid",
	})
	return c.SendStatus(400)
}
