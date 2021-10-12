package main

import (
	"os"

	"github.com/KlareTeam/interview-challenges/go/pricematic/database"
	"github.com/KlareTeam/interview-challenges/go/pricematic/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
)

/**
* Se importan las rutas creadas y se asignan a un grupo
 */
func setupRoutes(app *fiber.App) { 
  api := app.Group("/api/v1")
  routes.MainRoute(api.Group("/"))
  routes.ProductRoute(api.Group("/product"))
}
/**
* Se migran los modelos y se establece la data inicial
*/
func MigrationsAndSeed() {
  db := database.Connection()
  database.InitialMigration(db)
  database.DataSeed(db)
}
/**
*s
*/
func main() {
	app := fiber.New()
  app.Use(logger.New())
  app.Use(cors.New(cors.Config{
    AllowOrigins: "*",    
    AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
  }))

	setupRoutes(app)
  MigrationsAndSeed()

  port := os.Getenv("API_PORT")
	err := app.Listen(port)



	if err != nil {
		panic(err)
	}
  
}

