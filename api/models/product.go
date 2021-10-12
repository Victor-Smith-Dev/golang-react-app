package models

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// TODO: estrutura modelo Product:
//  - ID Clave primaria
//  - Name Nombre del producto
//  - Price Precio del producto
//  - CreatedAt Fecha de creación
//  - UpdatedAt Fecha de última actualización
//  - DeletedAt FEcha de borrado
type Product struct {
	ID   uuid.UUID `gorm:"type:uuid"`
	Name string   `gorm:"size:50;not null;unique" json:"name"`
	Price float32 `json:"price"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// TODO: Obtener productos:
// 	- db referncia a la conexión con la db
// 	- c referencia a la solicitud recibida del cliente
func GetProducts(db *gorm.DB, c *fiber.Ctx)  error{
	var products []Product
	db.Find(&products)
	return c.JSON(fiber.Map { 
		"status": "Ok", 
		"data": products,
	})
}

// TODO: Obtiene un producto, definido por el parametro id:
//	- db referncia a la conexión con la db
//	- c referencia a la solicitud recibida del cliente
func GetProduct(db *gorm.DB, c *fiber.Ctx)  error {
	var product Product

	id := c.Params("id")
	db.Find(&product, "id = ?", id)

	if product.ID == uuid.Nil {
    return c.Status(404).JSON(fiber.Map {
			"status": "error", 	
			"data": nil,
		})
	}

	return c.JSON(fiber.Map {
		"status": "Ok", 
		"data": product,
	})
}

// TODO: Crea un nuevo producto a partir de c.Bodyparser():
//	- db referncia a la conexión con la db
//	- c referencia a la solicitud recibida del cliente
func NewProduct(db *gorm.DB, c *fiber.Ctx) error  {
	product := new(Product)

	err := c.BodyParser(product)
	if err != nil {
		return c.Status(500).JSON(fiber.Map {
			"status": "error", 
			"message": "Review your input", 
			"data": err,
		})	
 	}

	product.ID = uuid.NewV4()

	err = db.Create(&product).Error
  if err != nil {
		return c.Status(500).JSON(fiber.Map {
			"status": "error", 
			"data": err,
		})
	}

	return c.JSON(fiber.Map {
		"status": "Ok", 
		"data": product,
	})
}

// TODO: Elimina un producto, definido por el parametro id:
//	- db referncia a la conexión con la db
//	- c referencia a la solicitud recibida del cliente
func DeleteProduct(db *gorm.DB, c *fiber.Ctx) error {
	var product Product

	id := c.Params("id")
	db.Find(&product, "id = ?", id)

	if product.ID == uuid.Nil {
    return c.Status(404).JSON(fiber.Map {
			"status": "error", 	
			"data": nil,
		})
	}
	err := db.Delete(&product, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map {
			"status": "error", 
			"data": nil,
		})
	}

	return c.JSON(fiber.Map {
		"status": "OK",
		"message": "Producto Borrado",
	})
}

// TODO: Modifica los datos de un producto, dependiendo el verbo HTTP puede agregar un registro en la tabla prices:
//	- db referncia a la conexión con la db
//	- c referencia a la solicitud recibida del cliente
func ModifyProduct(db *gorm.DB, c *fiber.Ctx) error {
	type updateProduct struct {
		Name string   `json:"name"`
		Price float32 `json:"price"`
	}

	var product Product

	db.Find(&product, "id = ?", c.Params("id"))

	if product.ID == uuid.Nil  {
		return c.Status(404).JSON(fiber.Map {
			"status": "error", 
			"data": nil,
		})
	}
	
	var updateProductData updateProduct
	err := c.BodyParser(&updateProductData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map {
			"status": "error", 
			"message": "Review your input", 
			"data": err,
		})
	}

	if c.Method() == "PATCH" && updateProductData.Price != product.Price { 
		var price Price
		price.ID = uuid.NewV4()
		price.ProductId = product.ID
		price.Price = updateProductData.Price
		db.Create(&price)
	}

	product.Name = updateProductData.Name
	product.Price = updateProductData.Price

	db.Save(&product) 
	return c.JSON(fiber.Map {
		"status": "success", 
		"data": product,
	})
}

// TODO: Obtiene el historial de precios de un producto definido por el parametro id:
//	- db referncia a la conexión con la db
//	- c referencia a la solicitud recibida del cliente
func GetPrices(db *gorm.DB, c *fiber.Ctx)  error{
	var prices []Price
	db.Find(&prices, "product_id = ?", c.Params("id"))
	return c.JSON(fiber.Map { 
		"status": "success", 
		"data": prices,
	})
}