package routes

import (
	"library_app/controllers"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api")
	
	// Auth
    api.Post("/login", controllers.Login)
    api.Post("/register", controllers.Register)
	
	protected := api.Group("", jwtware.New(jwtware.Config{
		SigningKey: []byte("secret_jwt_key"),
	}))
	
	// Books
    protected.Get("/books", controllers.GetBooks)
    protected.Post("/books", controllers.StoreBook)
}
