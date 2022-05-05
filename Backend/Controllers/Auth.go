package backend

import "github.com/gofiber/fiber/v2"

func CreateUser(c* fiber.Ctx) error {
	return c.SendString("Hello, mate!")
}

func LoginUser() {

}

func DeleteUser() {

}
