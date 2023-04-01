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
	r.Put("/api/user/changeAvatar/:id", middleware.RequireUser, controller.ChangeAvatar)
	r.Get("/api/user/getAll", middleware.RequireUser, controller.GetAllUser)
	r.Delete("/api/user/delete/:username", middleware.RequireUser, controller.DeleteUser)
	r.Put("/api/user/update/:username", middleware.RequireUser, controller.UpdateInformation)
}
