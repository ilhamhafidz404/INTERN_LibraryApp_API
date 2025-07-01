package main

import (
	"library_app/database"
	_ "library_app/docs"
	"library_app/models"
	"library_app/routes"

	swagger "github.com/arsmn/fiber-swagger/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title Library App API
// @version 1.0
// @description REST API untuk manajemen buku dan user
// @host localhost:3000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Masukkan token Anda dengan format: Bearer <token>
func main() {
    app := fiber.New()

    // -------- Connect & Migrate DB
    database.Connect()
    database.DB.AutoMigrate(
        &models.Admin{},
        models.Book{},
        &models.Student{},
    )

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Library App API Service!")
    })


    // -------- Swagger Init
    app.Use(cors.New())
    app.Get("/swagger/*", swagger.HandlerDefault)


    // -------- Setup other routes
    routes.SetupRoutes(app)

    // Start server
    app.Listen(":3000")
}
