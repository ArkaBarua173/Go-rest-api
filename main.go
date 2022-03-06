package main

import (
	"log"

	"example.com/go-rest-api/database"
	"example.com/go-rest-api/routes"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	// user endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("api/users/:id", routes.GetUser)
	app.Put("api/users/:id", routes.UpdateUser)
	app.Delete("api/users/:id", routes.DeleteUser)
	// product endpoints
	app.Post("api/products", routes.CreateProduct)
	app.Get("api/products", routes.GetProducts)
	app.Get("api/products/:id", routes.GetProduct)
	app.Put("api/products/:id", routes.UpdateProduct)
	app.Delete("api/products/:id", routes.DeleteProduct)
}

func main() {

	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
