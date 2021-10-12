package controllers

import (
	"github.com/KlareTeam/interview-challenges/go/pricematic/database"
	"github.com/KlareTeam/interview-challenges/go/pricematic/models"
	"github.com/gofiber/fiber/v2"
)

// TODO: Obtiene todos los productos insertado en la base de datos
func GetProducts(c *fiber.Ctx) error {
	db := database.Connection()
	return models.GetProducts(db, c)
}

// TODO: Obtiene un producto, definido por el parametro {:id}
func GetProduct(c *fiber.Ctx) error {
	db := database.Connection()
	return models.GetProduct(db, c)
}

// TODO: Crea un nuevo producto a partir de c.Bodyparser()
func NewProduct(c *fiber.Ctx) error {
  db := database.Connection()
	return models.NewProduct(db, c)
}

// TODO: Elimina un producto, definido por el parametro {:id}
func DeleteProduct(c *fiber.Ctx) error {
	db := database.Connection()
	return models.DeleteProduct(db, c)
}

// TODO: Modifica los datos de un producto, dependiendo el verbo HTTP puede agregar un registro en la tabla prices
func ModifyProduct(c *fiber.Ctx) error {
	db := database.Connection()
	return models.ModifyProduct(db, c)
}

// TODO: Obtiene el historial de precios de un producto definido por el parametro {:id}
func GetPrices(c *fiber.Ctx) error {
	db := database.Connection()
	return models.GetPrices(db, c)
}
