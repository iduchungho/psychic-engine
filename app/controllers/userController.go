package controller

import (
	"github.com/gofiber/fiber/v2"
	model "smhome/app/models"
	service "smhome/pkg/services"
	"smhome/pkg/utils"
	"smhome/platform/cache"
)

func GetUserByID(c *fiber.Ctx) error {
	id := c.Query("id", "none")
	if id == "none" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "require ?id = ...",
			"success": false,
		})
	}
	userService := service.NewUserService()
	res, err := userService.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    res,
		"status":  "ok",
		"success": true,
	})
}

func Login(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if c.BodyParser(&body) != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to parse body",
			"success": false,
		})
	}
	userService := service.NewUserService()
	res, err := userService.Login(c, body.Username, body.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"success": false,
		})
	}
	sess, err := cache.GetSessionStore().Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   err.Error(),
			"success": false,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"tokenString": sess.Get("Authorization"),
		"data":        res,
	})
}

func AddNewUser(c *fiber.Ctx) error {
	var userMd model.User
	if c.BodyParser(&userMd) != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "can't read body request",
		})
	}

	userService := service.NewUserService()
	userDoc, err := userService.RegisterUser(userMd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": *userDoc,
	})
}

func Logout(c *fiber.Ctx) error {

	/**
	* TODO: set authorization for more people
	*		key: Authorization1, Authorization2, ....
	*		api src : /api/user/logout/:id
	*		find id to get session_id then delete it
	**/

	msg, err := utils.RemoveDataSession(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"success": false,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    msg,
		"success": true,
	})
}

func ChangeAvatar(c *fiber.Ctx) error {
	var avtar struct {
		Avt string `json:"avt" form:"avt"`
	}

	if c.BodyParser(&avtar) != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse body",
		})
	}

	// get file header
	fileHeader, errFile := c.FormFile("avt")
	if errFile != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errFile.Error(),
		})
	}

	id := c.Query("id", "none")
	if id == "none" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "require ?id = ...",
			"success": false,
		})
	}
	userService := service.NewUserService()
	userRepo, err := userService.ChangeAvatarByID(id, fileHeader)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    *userRepo,
	})
}

//	func GetAllUser(c *fiber.Ctx) error {
//		user, err := service.NewEntityContext("user")
//		if err != nil {
//			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//				"error": err.Error(),
//			})
//		}
//		res, errRes := user.GetEntity("")
//		if errRes != nil {
//			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//				"error": errRes.Error(),
//			})
//		}
//		return c.Status(fiber.StatusOK).JSON(fiber.Map{
//			"data": res,
//		})
//	}
func DeleteUser(c *fiber.Ctx) error {
	id := c.Query("id", "none")
	if id == "none" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "require ?id = ...",
			"success": false,
		})
	}
	userService := service.NewUserService()
	err := userService.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"id":      id,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "user was deleted",
		"id":      id,
	})
}

//

func UpdateInformation(c *fiber.Ctx) error {
	var body struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}

	if c.BodyParser(&body) != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "can't read body request",
			"success": false,
		})
	}

	id := c.Query("id", "none")
	if id == "none" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "require ?id = ...",
			"success": false,
		})
	}

	userService := service.NewUserService()
	res, err := userService.UpdateInfo(id, body.FirstName, body.LastName, body.Email, body.Phone)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"success": false,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": res,
	})
}

func UpdatePassword(c *fiber.Ctx) error {
	id, err := utils.RequireID(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "require ?id = ...",
			"success": false,
		})
	}
	userService := service.NewUserService()
	msg, err := userService.UpdatePass(c, *id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"success": false,
		})
	}
	ses, err := utils.RemoveDataSession(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"success": false,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": *msg,
		"status":  *ses,
	})
}
