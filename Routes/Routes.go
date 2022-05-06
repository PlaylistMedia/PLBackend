package backend

import (
	"github.com/gofiber/fiber/v2"
	Controllers "playlist.media/backend/Controllers"
)

func SetupRoutes(app *fiber.App) {
	// Auth routes
	app.Post("/signup", Controllers.CreateUser)       // Create a user
	app.Post("/login", Controllers.LoginUser)         // Log in
	app.Delete("/delete/:id", Controllers.DeleteUser) // Delete a user

	// Room routes
	app.Get("/room/get_uid", Controllers.GetRoomUID)                 // Get the UID of the current room
	app.Get("/room/current", Controllers.CurrentlyPlaying)           // Get the URL, timestamp of the video current playing
	app.Post("/room/create_room", Controllers.CreateRoom)            // Create a new room
	app.Post("/room/queue_add", Controllers.QueueAdd)                // Add a new link to the queue
	app.Put("/room/queue_remove", Controllers.QueueRemove)           // Remove a link from the queue
	app.Patch("/room/update_timestamp", Controllers.UpdateTimestamp) // Update the timestamp of a video (requires URL)
	app.Delete("/room/delete_room/:id", Controllers.DeleteRoom)      // Delete a room
}
