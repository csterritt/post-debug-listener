package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/gofiber/fiber/v2"
)

type ShowLine struct {
	Line   string `json:"line"`
	Sender string `json:"sender"`
	Type   string `json:"type"`
}

func main() {
	quiet := len(os.Args) == 2 && os.Args[1] == "-q"
	if len(os.Args) != 1 && !quiet {
		fmt.Fprintln(os.Stderr, "Usage: post-debug-listener [-q]\n    -q -- don't use colors\n")
		os.Exit(1)
	}
	fmt.Printf("Using colors: %v\n", !quiet)

	//inverse := lipgloss.NewStyle().
	//	Bold(true).
	//	Foreground(lipgloss.AdaptiveColor{Light: "255", Dark: "0"}).
	//	Background(lipgloss.AdaptiveColor{Light: "245", Dark: "15"})

	green := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "15", Dark: "0"}).
		Background(lipgloss.AdaptiveColor{Light: "118", Dark: "118"})

	yellow := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "15", Dark: "0"}).
		Background(lipgloss.AdaptiveColor{Light: "226", Dark: "226"})

	red := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "0", Dark: "15"}).
		Background(lipgloss.AdaptiveColor{Light: "9", Dark: "9"})

	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		line := new(ShowLine)
		if err := c.BodyParser(line); err != nil {
			return err
		}

		if quiet {
			fmt.Printf("%s: %s %s\n", line.Sender, line.Type, line.Line)
		} else {
			var itsType string
			switch line.Type {
			case "INFO":
				itsType = green.Render(" " + line.Type + " ")
			case "WARN":
				itsType = yellow.Render(" " + line.Type + " ")
			case "ERROR":
				itsType = red.Render(" " + line.Type + " ")
			default:
				itsType = green.Render(" " + line.Type + " ")
			}

			fmt.Println(line.Sender + " " + itsType + " " + line.Line)
		}

		return c.SendString("good")
	})

	app.Listen(":3030")
}
