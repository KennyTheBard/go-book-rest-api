package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"hellkite.eu/go-api/controllers"
	"hellkite.eu/go-api/middlewares"
	"hellkite.eu/go-api/models"
)

func main() {
	var err error

	// database connection
	err = models.InitDatabaseConnection()
	if err != nil {
		panic(err)
	}

	// configure server
	app := fiber.New()
	app.Use(cors.New())

	// configure routes
	publicGroup := app.Group("/public")
	publicGroup.Get("/authenticate", controllers.Authenticate)
	publicGroup.Post("/register", controllers.CreateUser)
	publicGroup.Post("/getAllUsers", controllers.GetAllUsers)
	publicGroup.Post("/getUserById", controllers.GetUserById)
	publicGroup.Post("/updateUserName", controllers.UpdateUserName)
	publicGroup.Post("/createOrder", controllers.CreateOrder)
	publicGroup.Post("/getAllOrders", controllers.GetAllOrders)
	publicGroup.Static("/", "./public/index.html")

	apiGroup := app.Group("/api", middlewares.JwtAuthMiddleware)
	apiGroup.Get("/getStuff", controllers.GetStuff)
	apiGroup.Post("/pushThingy", controllers.PushThingy)

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
