package route

import (
	"smhome/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SenSorRoute(r *fiber.App) {
	r.Get("/api/sensor/temperature", controller.GetTemperature)
	r.Get("/api/sensor/humidity", controller.GetHumidity)
	r.Get("/api/sensor/light", controller.GetLight)
}
