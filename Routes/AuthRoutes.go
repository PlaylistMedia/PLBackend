package backend

import (
	"github.com/gofiber/fiber/v2"
	Auth "playlist.media/backend/Controllers"
)

func SetupRoutes(app* fiber.App) {
	app.Post("/signup", Auth.CreateUser)
	app.Post("/login", Auth.LoginUser)
	app.Delete("/delete/:id", Auth.DeleteUser)
}