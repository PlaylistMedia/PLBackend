package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v4/stdlib"
	Routes "playlist.media/backend/Routes"
)

func GetFromDB(db *sql.DB) {
	// Get from database
	row := db.QueryRowContext(context.Background(), "SELECT password FROM users WHERE name='steve';")
	if err := row.Err(); err != nil {
		fmt.Println("db.QueryRowContext", err)
		return
	}

	var password string
	if err := row.Scan(&password); err != nil {
		fmt.Println("row.Scan", err)
		return
	}

}

func main() {
	app := fiber.New()

	Routes.SetupRoutes(app)

	dsn := url.URL{
		Scheme: "postgres",
		Host:   "localhost:5432",
		User:   url.User("postgres"),
		Path:   "postgres",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	db, err := sql.Open("pgx", dsn.String())

	if err != nil {
		fmt.Println("sql.Open", err)
		return
	}

	defer func() {
		_ = db.Close()
		fmt.Println("Closed connection to database.")
	}()

	// Set in database
	if _, err :=
		db.ExecContext(context.Background(), "INSERT INTO users(name, password) VALUES('steve', 'his password');"); err != nil {
		fmt.Println("db.ExecContext", err)
	}

	GetFromDB(db)

	app.Listen(":3000")
}
