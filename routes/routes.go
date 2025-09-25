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
    api.Post("/logout", controllers.Register)
	
	protected := api.Group("", jwtware.New(jwtware.Config{
		SigningKey: []byte("secret_jwt_key"),
	}))

    // Dashboard
    protected.Get("/dashboard", controllers.GetDashboard)
	
	// Books
    protected.Get("/books", controllers.GetBooks)
    protected.Post("/books", controllers.StoreBook)
    protected.Put("/books/:id", controllers.UpdateBook)
    protected.Delete("/books/:id", controllers.DeleteBook)

	//Profile
    protected.Get("/profile/:student_id", controllers.GetProfile)
    protected.Put("/profile/:student_id", controllers.UpdateProfile)
    protected.Put("/profile/change-password/:student_id", controllers.ChangePassword)
	
    //Lending History
    protected.Get("/lending-history", controllers.GetLendingHistory)
    protected.Post("/lending-history", controllers.PostLendingHistory)
    protected.Put("/lending-history/:history_id", controllers.PutLendingHistory)
    protected.Delete("/lending-history/:history_id", controllers.DeleteLendingHistory)
}
