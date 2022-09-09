package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"

	"github.com/baariqazhar/go-restapi-fiber/controller/bookcontroller"
	"github.com/baariqazhar/go-restapi-fiber/models"
)

func main() {
	fastergoding.Run()
	models.ConnectDatabase()

	app := fiber.New()

	api := app.Group("/api")
	book := api.Group("/book")

	book.Get("/", bookcontroller.GetBooks)
	book.Get("/:id", bookcontroller.GetBook)
	book.Post("/", bookcontroller.CreateBook)
	book.Put("/:id", bookcontroller.UpdateBook)

	app.Listen(":3001")
}
