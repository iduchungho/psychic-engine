package route

import (
	"github.com/gofiber/fiber/v2"
	"smhome/app/controllers"
	"smhome/pkg/middleware"
)

func ActionRoute(r *fiber.App) {
	r.Get("/api/action/get", controller.GetActionByID)
	r.Post("/api/action/log", middleware.RequireUserByID, controller.PushActionLog)
}
