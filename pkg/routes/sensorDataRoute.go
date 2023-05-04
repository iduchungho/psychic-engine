package route

import (
	"github.com/gofiber/fiber/v2"
	controller "smhome/app/controllers"
)

func SensorDataRoute(r *fiber.App) {
	r.Get("/api/data/stat", controller.SensorStats)
}
