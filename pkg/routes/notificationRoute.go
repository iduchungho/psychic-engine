package route

import (
	"github.com/gofiber/fiber/v2"
	"smhome/app/controllers"
	"smhome/pkg/middleware"
)

func NotifyRoute(r *fiber.App) {
	r.Get("/api/noty/get", middleware.RequireUserByID, controller.GetNotyByUserID)
	r.Post("/api/noty/push", middleware.RequireUserByID, controller.PushNoty)
}
