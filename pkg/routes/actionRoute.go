package route

import (
	"github.com/gofiber/fiber/v2"
	controller "smhome/app/controllers"
	"smhome/pkg/middleware"
)

func ActionRoute(r *fiber.App) {
	r.Get("/api/action/get/:id", controller.GetActionByID)
	r.Post("/api/action/log/:id", middleware.RequireUserID, controller.PushActionLog)
}
