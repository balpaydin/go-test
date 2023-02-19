package book

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct {
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {

	books:= []Book{
		Book{
			Title: "başlık 1",
			Author: "Kafka",
			Rating: 12,
		},
		Book{
			Title: "başlık 2",
			Author: "Kafka",
			Rating: 100,
		},
	}

	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	c.SendString("Single Book")
}

func NewBook(c *fiber.Ctx) error {
	book := new (Book)
	if err:=c.BodyParser(book); err !=nil{
		return c.Status(503).SendString(err.Error())
	}

	
	return c.JSON(&book)
}

func DeleteBook(c *fiber.Ctx) {
	c.SendString("Delete Book")
}
