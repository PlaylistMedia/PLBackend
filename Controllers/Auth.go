package backend

import (
	"github.com/gofiber/fiber/v2"
)

// TODOs: Add a database and hook up the functions

type User struct {
	Username  string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
}

func CreateUser(ctx *fiber.Ctx) error {
	// Create a new user object for later population
	user := new(User)

	// Populate the user object
    if err := ctx.BodyParser(&user); err != nil {
		// Tell the client that the sent info isn't parseable.
		return fiber.NewError(fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message)
    }
	
	// Update the password with the hash
    hash, err := HashPassword(user);
	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message)
	} else {
		user.Password = hash;
	}
	
	user.Token, err = GenerateJWT(user);

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return fiber.NewError(fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message)
	}

	// TODO: Push user to database
    return ctx.JSON(user)

}

func LoginUser(ctx *fiber.Ctx) error {
	// Create a new user object for later population
	user := new(User)

	// Populate the user object
    if err := ctx.BodyParser(&user); err != nil {
		// Tell the client that the sent info isn't parseable.
		return fiber.NewError(fiber.ErrBadRequest.Code, fiber.ErrBadRequest.Message)
    }

	// TODO: Query the database for a user, get their info
	// Using demo data for now
	user_db := new(User);
	user_db.Username = "Steve"
	user_db.Password = "9c34c004fc342112de3e5c25a0a3b127aeb9d8f1a691d224ad10bf1f4b25936d0857cd26495c185284aacee5108245de7916e55279337e33f6cd5898e4f2dc8b"
	user_db.Email = "email2@email.com"

	// Hash the password (Request user)
	hash, err := HashPassword(user);

	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message)
	}
	// Check if the passwords match
	if hash != user_db.Password {
		return fiber.NewError(fiber.ErrUnauthorized.Code, fiber.ErrUnauthorized.Message);
	}
	
	// TODO: UPDATE THE TOKEN IN DB
	user_db.Token, err = GenerateJWT(user);

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return fiber.NewError(fiber.ErrInternalServerError.Code, fiber.ErrInternalServerError.Message)
	}

	return ctx.JSON(fiber.Map{
		"message": "User Logged in!",
		"token": user_db.Token,
	})
}

func DeleteUser(ctx* fiber.Ctx) error {
	return ctx.SendString("Delete User.")
}
