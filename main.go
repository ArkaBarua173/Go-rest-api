package main

import (
	"log"

	"example.com/go-rest-api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectDb()

	app := fiber.New()

	log.Fatal(app.Listen(":3000"))

}
