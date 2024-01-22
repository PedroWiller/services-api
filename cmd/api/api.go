package api

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"face-id-detection/config/env"
	healthRoute "face-id-detection/internal/modules/health/route"
)

func Start() {
	app := fiber.New()
	healthRoute.Init(app)

	api := *fiber.New()
	app.Mount("/api/v1", &api)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", env.PORT)))
}
