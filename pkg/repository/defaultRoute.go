package repo

var DefaultRoutes = []string{
	// user default routes
	"/api/user/login",
	"/api/user/new",
	"/api/user/logout",
	"/api/user/changeAvatar",
	"/api/user/getAll",
	"/api/user/delete",
	"/api/user/update",
	"/api/user/getUserById",
	// sensor default routes
	"/api/sensor/temperature",
	"/api/sensor/humidity",
	"/api/sensor/light",
	// action default routes
	"/api/action/get",
	"/api/action/log",
}

//r.Post("/api/user/login", controller.Login)
//r.Post("/api/user/new", controller.AddNewUser)
//r.Get("/api/user/logout", middleware.RequireUser, controller.Logout)
//r.Put("/api/user/changeAvatar/:id", middleware.RequireUser, controller.ChangeAvatar)
//r.Get("/api/user/getAll", middleware.RequireUser, controller.GetAllUser)
//r.Delete("/api/user/delete/:username", middleware.RequireUser, controller.DeleteUser)
//r.Put("/api/user/update/:username", middleware.RequireUser, controller.UpdateInformation)
