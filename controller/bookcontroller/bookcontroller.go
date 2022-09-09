package bookcontroller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/baariqazhar/go-restapi-fiber/models"
)

func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	models.DB.Find(&books)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Show all books",
		"data":    books,
	})
}

func CreateBook(c *fiber.Ctx) error {

	var book models.Book

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(&book).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "New book has created",
		"data":    book,
	})
}
