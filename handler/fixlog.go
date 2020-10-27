package handler

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/model"

	"github.com/gofiber/fiber/v2"
)

// GetAllProducts query all fixlogs
func GetAllFixlogs(c *fiber.Ctx) error {
	db := database.DB
	var fixlogs []model.Fixlogs
	db.Find(&fixlogs)
	return c.JSON(fiber.Map{"status": "success", "message": "All products", "data": fixlogs})
}

// GetProduct query fixlog
func GetFixlog(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var fixlog model.Fixlogs
	db.Find(&fixlog, id)
	if fixlog.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})

	}
	return c.JSON(fiber.Map{"status": "success", "message": "Product found", "data": fixlog})
}

// CreateProduct new fixlog
func CreateFixlog(c *fiber.Ctx) error {
	db := database.DB
	fixlog := new(model.Fixlogs)
	if err := c.BodyParser(fixlog); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create product", "data": err})
	}
	db.Create(&fixlog)
	return c.JSON(fiber.Map{"status": "success", "message": "Created product", "data": fixlog})
}

// DeleteProduct delete fixlog
func DeleteFixlog(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var fixlog model.Fixlogs
	db.First(&fixlog, id)
	if fixlog.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})

	}
	db.Delete(&fixlog)
	return c.JSON(fiber.Map{"status": "success", "message": "Product successfully deleted", "data": nil})
}
