package backend

import (
	"github.com/gofiber/fiber/v2"
	Auth "playlist.media/backend/Backend/Controllers"
)

func SetupRoutes(app* fiber.App) {
	app.Get("/signup", Auth.CreateUser)
}