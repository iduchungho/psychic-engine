package route

import (
	"github.com/gofiber/fiber/v2"
	controller "smhome/app/controllers"
)

func ActionRoute(r *fiber.App) {
	r.Get("/api/action/get/:username", controller.GetActionByUsername)
	r.Post("/api/action/log", controller.PushActionLog)
}
