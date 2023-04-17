package route

import (
	"github.com/gofiber/fiber/v2"
	controller "smhome/app/controllers"
	"smhome/pkg/middleware"
)

func UserRoute(r *fiber.App) {
	//r.Get("/api/user/getAll/:id", middleware.RequireUserID, controller.GetAllUser)
	r.Get("/api/user/logout", middleware.RequireUser, controller.Logout)
	r.Get("/api/user/getUserById/:id", middleware.RequireUserByID, controller.GetUserByID)
	r.Post("/api/user/login", controller.Login)
	r.Post("/api/user/new", controller.AddNewUser)
	r.Put("/api/user/changeAvatar/:id", middleware.RequireUserByID, controller.ChangeAvatar)
	r.Put("/api/user/update/:id", middleware.RequireUserByID, controller.UpdateInformation)
	r.Delete("/api/user/delete/:id", middleware.RequireUserByID, controller.DeleteUser)

}
