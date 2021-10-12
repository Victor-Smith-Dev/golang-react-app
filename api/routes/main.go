package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/KlareTeam/interview-challenges/go/pricematic/controllers"
)

func MainRoute(route fiber.Router) {
	route.Get("", controllers.GetMain)
}