package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"

	"example/frontend"
)

func Start() error {
	server := fiber.New()

	// Serve static files
	server.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(frontend.FS),
		PathPrefix: "web/.output/public",
	}))

	// For all other routes serve the "index.html"
	server.Use("*", func(c *fiber.Ctx) error {
		return filesystem.SendFile(c, http.FS(frontend.FS), "web/.output/public/index.html")
	})

	return server.Listen(":3000")
}
