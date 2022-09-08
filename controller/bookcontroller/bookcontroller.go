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
		"message": "All books",
		"data":    books,
	})
}

func CreateBook(c *fiber.Ctx) error {

	var book string

	coba := c.BodyParser(&book)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Create book",
		"data":    book,
		"coba":    coba,
	})
}
