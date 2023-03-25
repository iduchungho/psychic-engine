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

	// get file header
	fileHeader, errFile := c.FormFile("avatar")
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
	userMd.Avatar = resp.SecureURL

	userMd.Password = string(hashPass)
	if errIs := newUser.InsertData(userMd); errIs != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errIs.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": userMd,
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
	return nil
}
