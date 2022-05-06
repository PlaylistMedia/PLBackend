package backend

import "github.com/gofiber/fiber/v2"

type Media struct {
	URL       string `json:"url"`
	Timestamp int64  `json:"timestamp"`
}

type Room struct {
	UID     string  `json:"uid"`     // UID of the room
	Members []User  `json:"members"` // Array of uid(s)
	Current Media   // URL of the video currently being played
	Queue   []Media // List of URLs in the queue
}

func GetRoomUID(ctx *fiber.Ctx) error {

	return fiber.NewError(fiber.ErrNotImplemented.Code, fiber.ErrNotImplemented.Message)
}

func CreateRoom(ctx *fiber.Ctx) error {

	return fiber.NewError(fiber.ErrNotImplemented.Code, fiber.ErrNotImplemented.Message)
}

func QueueAdd(ctx *fiber.Ctx) error {

	return fiber.NewError(fiber.ErrNotImplemented.Code, fiber.ErrNotImplemented.Message)
}

func QueueRemove(ctx *fiber.Ctx) error {

	return fiber.NewError(fiber.ErrNotImplemented.Code, fiber.ErrNotImplemented.Message)
}

func UpdateTimestamp(ctx *fiber.Ctx) error {

	return fiber.NewError(fiber.ErrNotImplemented.Code, fiber.ErrNotImplemented.Message)
}

func CurrentlyPlaying(ctx *fiber.Ctx) error {

	room := new(Room)

	// Populate the room object
	if err := ctx.BodyParser(&room); err != nil {
		// Tell the client that the sent info isn't parseable.
		return fiber.NewError(fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message)
	}

	return ctx.JSON(fiber.Map{
		"url":       room.Current.URL,
		"timestamp": room.Current.Timestamp, // In seconds
	})
}

func DeleteRoom(ctx *fiber.Ctx) error {

	return fiber.NewError(fiber.ErrNotImplemented.Code, fiber.ErrNotImplemented.Message)
}
