package main

import (
	"social/api/src/controllers/auth"
	"social/api/src/controllers/book"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"

	jwtware "github.com/gofiber/jwt/v3"
)

func setupMiddleware(app *fiber.App) {
	app.Use(helmet.New())
	app.Use(limiter.New())
	app.Use(logger.New())

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
}

func setupRoutes(app *fiber.App) {

	
	app.Get("/api/user",auth.Restricted);
	app.Get("/api/books", book.GetBooks)
	app.Post("/api/book",book.NewBook)
}

func main() {
    app := fiber.New()

	//configs.ConnectDB()

	app.Post("/api/login",auth.Login)
	//app.Post("/api/user",controllers.CreateUser)
	setupMiddleware(app)
	setupRoutes(app)
    app.Listen(":3000")
}