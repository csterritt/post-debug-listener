package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ShowLine struct {
	Line   string `json:"line"`
	Sender string `json:"sender"`
	Type   string `json:"type"`
}

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		line := new(ShowLine)
		if err := c.BodyParser(line); err != nil {
			return err
		}

		fmt.Printf("%s: %s %s\n", line.Sender, line.Type, line.Line)

		return c.SendString("good")
	})

	app.Listen(":3030")
}
