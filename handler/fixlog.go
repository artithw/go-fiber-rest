package handler

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/model"

	"github.com/gofiber/fiber/v2"
)

// GetAllFixlogs query all fixlogs
func GetAllFixlogs(c *fiber.Ctx) error {
	db := database.DB
	var fixlogs []model.Fixlogs
	db.Find(&fixlogs)
	return c.JSON(fiber.Map{"status": "success", "message": "All fixlogs", "data": fixlogs})
}

// GetFixlog query fixlog
func GetFixlog(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var fixlog model.Fixlogs
	db.Find(&fixlog, id)
	if fixlog.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No fixlogs found with ID", "data": nil})

	}
	return c.JSON(fiber.Map{"status": "success", "message": "Fixlog found", "data": fixlog})
}

// CreateFixlog new fixlog
func CreateFixlog(c *fiber.Ctx) error {
	db := database.DB
	fixlog := new(model.Fixlogs)
	if err := c.BodyParser(fixlog); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create fixlog", "data": err})
	}
	db.Create(&fixlog)
	return c.JSON(fiber.Map{"status": "success", "message": "Created fixlog", "data": fixlog})
}

// DeleteFixlog delete fixlog
func DeleteFixlog(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var fixlog model.Fixlogs
	db.First(&fixlog, id)
	if fixlog.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No fixlog found with ID", "data": nil})

	}
	db.Delete(&fixlog)
	return c.JSON(fiber.Map{"status": "success", "message": "Fixlog successfully deleted", "data": nil})
}
