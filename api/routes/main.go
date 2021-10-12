package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/KlareTeam/interview-challenges/go/pricematic/controllers"
)

//TODO: Ruta para renderizar cliente
func MainRoute(route fiber.Router) {
	route.Get("", controllers.GetMain)
}