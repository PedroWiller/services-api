package healthController

import (
	"github.com/gofiber/fiber/v2"

	"face-id-detection/internal/modules/health/dto"
)

func HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&dto.Health{
		Status: "UP",
	})
}

// app := fiber.New()
// app.Use(
// 	func(c *fiber.Ctx) error {
// 		err := c.Next()
// 		log.Printf("status: %v", c.Response().StatusCode())
// 		return err
// 	},
// 	recover.New(),
// )

// app.Listen(":3000")
