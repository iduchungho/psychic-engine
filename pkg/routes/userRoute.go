package route

import (
	"github.com/gofiber/fiber/v2"
	"smhome/app/controllers"
	"smhome/pkg/middleware"
)

func UserRoute(r *fiber.App) {
	r.Get("/api/user/getAll/:id", middleware.RequireUserID, controller.GetAllUser)
	r.Get("/api/user/logout/:id", middleware.RequireUserID, controller.Logout)
	r.Get("/api/user/getUserById/:id", middleware.RequireUserID, controller.GetUserByID)
	r.Post("/api/user/login", controller.Login)
	r.Post("/api/user/new", controller.AddNewUser)
	r.Put("/api/user/changeAvatar/:id", middleware.RequireUserID, controller.ChangeAvatar)
	r.Put("/api/user/update/:id", middleware.RequireUserID, controller.UpdateInformation)
	r.Delete("/api/user/delete/:id", middleware.RequireUserID, controller.DeleteUser)

}
