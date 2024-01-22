package healthRoute

import (
	"github.com/gofiber/fiber/v2"

	healthController "face-id-detection/internal/modules/health/controller"
)

func Init(app *fiber.App) {
	app.Get("/health", healthController.HealthCheck)
}
