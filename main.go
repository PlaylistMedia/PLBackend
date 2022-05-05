package main

import (
	"github.com/gofiber/fiber/v2"
	AuthRoutes "playlist.media/backend/Backend/Routes"
)

func main() {
    app := fiber.New();

    AuthRoutes.SetupRoutes(app);

    app.Listen(":3000")
}
