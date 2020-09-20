package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/raolootnexii/trasle/database"
	"github.com/raolootnexii/trasle/song"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	songRoutes := api.Group("/song")
	songRoutes.Get("/", song.GetSongs)
	songRoutes.Post("/", song.CreateSong)
}

func initDatabase() {
	var err error
	database.DBConn, _ = gorm.Open("sqlite3", "trasle.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(&song.Song{})
	fmt.Println("Migrated")
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	initDatabase()

	app.Listen(":3000")
}
