package routes

import (
	"github.com/KlareTeam/interview-challenges/go/pricematic/controllers"
	"github.com/gofiber/fiber/v2"
)

func ProductRoute(route fiber.Router) {
	route.Get("/", controllers.GetProducts)
	route.Post("/", controllers.NewProduct)
	route.Get("/:id", controllers.GetProduct)
	route.Put("/:id", controllers.ModifyProduct)
	route.Patch("/:id", controllers.ModifyProduct)
	route.Delete("/:id", controllers.DeleteProduct)
	route.Get("/:id/prices", controllers.GetPrices)
}