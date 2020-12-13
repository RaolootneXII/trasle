package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raolootnexii/trasle/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	songRoutes := api.Group("/song")
	songRoutes.Get("/", handler.GetSongs)
	songRoutes.Get("/:id", handler.GetSong)
	songRoutes.Post("/", handler.CreateSong)
	songRoutes.Delete("/:id", handler.DeleteSong)
	songRoutes.Put("/:id", handler.UpdateSong)
}
