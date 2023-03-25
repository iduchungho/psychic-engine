package route

import (
	"github.com/gofiber/fiber/v2"
	"smhome/app/controllers"
	"smhome/pkg/middleware"
)

func UserRoute(r *fiber.App) {
	r.Post("/api/user/login", controller.Login)
	r.Post("/api/user/new", controller.AddNewUser)
	r.Get("/api/user/logout", middleware.RequireUser, controller.Logout)
}
