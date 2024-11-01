package main

import (
	"embed"
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/gofiber/fiber/v2"
	cli "github.com/jawher/mow.cli"
)

type ShowLine struct {
	Line   string `json:"line"`
	Sender string `json:"sender"`
	Type   string `json:"type"`
}

var (
	//go:embed resources
	resourceReader embed.FS
	pages          = map[string]string{
		"curl":       "resources/curl",
		"go":         "resources/go",
		"javascript": "resources/javascript",
	}
)

func provideExample(language string) {
	if page, found := pages[language]; found {
		content, err := resourceReader.ReadFile(page)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Got an error reading %s file: %v\n", language, err)
		} else {
			fmt.Println(string(content))
		}
	} else {
		fmt.Println("There is no example for the", language, "language. Please submit one!")
	}
}

func writeLineToFile(line string, outputFile string) {
	file, err := os.OpenFile(outputFile,
		os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Got error creating file: ", err)
		return
	}
	defer func() {
		err = errors.Join(err, file.Close())
	}()

	_, err = file.WriteString(line)
	if err != nil {
		fmt.Println("Got error writing to file: ", err)
	}
}

func runServer(quiet bool, port int, outputFile string) {
	fmt.Printf("Using colors: %v\n", !quiet)
	if outputFile != "" {
		fmt.Printf("Also writing lines to %s\n", outputFile)
	}

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

		if outputFile != "" {
			writeLineToFile(fmt.Sprintf("%s: %s %s\n", line.Sender, line.Type, line.Line), outputFile)
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

	err := app.Listen(fmt.Sprintf(`:%d`, port))
	if err != nil {
		panic(err)
	}
}

func main() {
	app := cli.App("post-debug-listener", "Listen on a port for debug POSTs, and print them")

	app.Spec = "[-q] [-p=<port>] [-o=<output_file>] [-e=<language>]"

	var (
		quiet      = app.BoolOpt("q quiet", false, "Run without colors")
		port       = app.IntOpt("p port", 3030, "Port to listen on")
		outputFile = app.StringOpt("o output", "", "Also write the posted lines to a file")
		language   = app.StringOpt("e example", "", "Provide a client example in the given language\n\t\t(curl, go, javascript)")
	)

	app.Action = func() {
		if *language != "" {
			provideExample(*language)
		} else {
			runServer(*quiet, *port, *outputFile)
		}
	}

	// Invoke the app passing in os.Args
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
