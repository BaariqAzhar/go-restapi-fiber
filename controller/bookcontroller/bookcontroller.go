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

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	if models.DB.Find(&book, id).RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Book not found",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Show book",
		"data":    book,
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

func UpdateBook(c *fiber.Ctx) error {

	id := c.Params("id")

	var book models.Book

	if err := c.BodyParser(&book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if models.DB.Where("id =? ", id).Updates(&book).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Can't update book data",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Book has updated",
		"data":    book,
	})
}
