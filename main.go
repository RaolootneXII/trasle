package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raolootnexii/trasle/song"
)

func main() {

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")

}

func setupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	songRoutes := api.Group("/song")
	songRoutes.Get("/", song.GetSongs)

}
