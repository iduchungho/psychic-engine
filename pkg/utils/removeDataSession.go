package utils

import (
	"github.com/gofiber/fiber/v2"
	"smhome/platform/cache"
)

func RemoveDataSession(c *fiber.Ctx) (*string, error) {
	sess, err := cache.GetSessionStore().Get(c)
	if err != nil {
		return nil, err
	}
	sess.Set("Authorization", -1)
	err = sess.Save()
	if err != nil {
		return nil, err
	}

	msg := "your session has been wiped"
	return &msg, nil
}
