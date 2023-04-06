package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"log"
	"os"
	"smhome/app/models"
	"smhome/pkg/services"
	"smhome/pkg/utils"
	"smhome/platform/cache"
	"smhome/platform/cloudinary"
)

func Login(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if c.BodyParser(&body) != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse body",
		})
	}

	user, _ := service.NewEntityContext("user")
	_, err := user.FindDocument("username", body.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	hashPass, _ := user.GetElement("password")
	err = utils.ComparePassword(*hashPass, body.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Generate a jwt token
	id, _ := user.GetElement("id")
	token := utils.GenerateToken(*id)

	// Sign and get the complete encode token as a string using the secret
	tokenString, errToken := token.SignedString([]byte(os.Getenv("SECRET")))
	if errToken != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create token",
		})
	}

	// send it back
	// c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	// c.JSON(http.StatusOK, gin.H{
	// 	"token": tokenString,
	// })
	sess, errSess := cache.GetSessionStore().Get(c)
	if errSess != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errSess.Error(),
		})
	}
	sess.Set("Authorization", tokenString)
	defer func(sess *session.Session) {
		err := sess.Save()
		if err != nil {
			log.Fatal(err)
		}
	}(sess)
	res, errRes := user.FindDocument("username", body.Username)
	if errRes != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errRes.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"tokenString": sess.Get("Authorization"),
		"data":        res,
	})
}

func AddNewUser(c *fiber.Ctx) error {
	var userMd model.User
	newUser, err := service.NewEntityContext("user")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err = newUser.SetElement("type", "user"); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if c.BodyParser(&userMd) != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "can't read body request",
		})
	}

	hashPass, err := utils.GenPassword(userMd.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userMd.Password = string(hashPass)
	if errIs := newUser.InsertData(userMd); errIs != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": errIs.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": newUser,
	})

}

func Logout(c *fiber.Ctx) error {
	sess, err := cache.GetSessionStore().Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	sess.Set("Authorization", -1)
	defer func(sess *session.Session) {
		err := sess.Save()
		if err != nil {
			log.Fatal(err)
		}
	}(sess)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "your session has been wiped",
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

	username := c.Params("id")
	user, err := service.NewEntityContext("user")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// get file header
	fileHeader, errFile := c.FormFile("avt")
	if errFile != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errFile.Error(),
		})
	}
	// open header file-header
	file, errOpen := fileHeader.Open()
	if errOpen != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errOpen.Error(),
		})
	}
	cld := cloud.GetConnCloudinary()
	resp, errCld := cloud.UpdateImages(cld, file)
	if errCld != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errCld.Error(),
		})
	}

	_, errFind := user.FindDocument("username", username)
	if errFind != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errFind.Error(),
		})
	}

	err = user.UpdateData("avatar", resp.SecureURL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err = user.SetElement("avatar", resp.SecureURL); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})
}

func GetAllUser(c *fiber.Ctx) error {
	user, err := service.NewEntityContext("user")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	res, errRes := user.GetEntity("")
	if errRes != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errRes.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	username := c.Params("username")
	user, err := service.NewEntityContext("user")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	err = user.DeleteEntity("username", username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "user was deleted",
		"user":    username,
	})
}

func UpdateInformation(c *fiber.Ctx) error {
	var body struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Password  string `json:"password"`
	}

	if c.BodyParser(&body) != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "can't read body request",
		})
	}

	username := c.Params("username")

	user, _ := service.NewEntityContext("user")
	_, err := user.FindDocument("username", username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	hashPass, err := utils.GenPassword(body.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err = user.UpdateData("password", string(hashPass)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err = user.UpdateData("firstname", body.FirstName); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err = user.UpdateData("lastname", body.LastName); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    body,
		"success": true,
		"message": "ok",
	})
}
