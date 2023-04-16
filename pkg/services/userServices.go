package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/pkg/errors"
	"log"
	"mime/multipart"
	"os"
	interfaces "smhome/app/interface"
	model "smhome/app/models"
	repo "smhome/pkg/repository"
	"smhome/pkg/utils"
	"smhome/platform/cache"
	cloud "smhome/platform/cloudinary"
	"smhome/platform/database"
)

type UserService struct {
	Factory interfaces.RepoFactory
}

func NewUserService() *UserService {
	return &UserService{
		Factory: NewFactory(database.GetCollection(repo.USER)),
	}
}

func (user *UserService) GetUserByID(id string) (*model.User, error) {
	users := user.Factory.NewUserRepo()
	return users.GetUserByID(id)
}

func (user *UserService) Login(ctx *fiber.Ctx, username string, pass string) (*model.User, error) {
	users := user.Factory.NewUserRepo()
	byUsername, err := users.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	// compare password
	hasPass := byUsername.Password
	err = utils.ComparePassword(hasPass, pass)
	if err != nil {
		return nil, nil
	}
	// generate token
	id := byUsername.Id
	token := utils.GenerateToken(id)
	// Sign and get the complete encode token as a string using the secret
	tokenString, errToken := token.SignedString([]byte(os.Getenv("SECRET")))
	if errToken != nil {
		return nil, errors.New("Failed to create token")
	}
	// send it session
	sess, errSess := cache.GetSessionStoreSlice(id).Get(ctx)
	if errSess != nil {
		return nil, errSess
	}
	// Save to session
	sess.Set("Authorization", tokenString)
	defer func(sess *session.Session) {
		err := sess.Save()
		if err != nil {
			log.Fatal(err)
		}
	}(sess)
	return byUsername, nil
}

func (user *UserService) RegisterUser(usr model.User) (*model.User, error) {
	userRepo := user.Factory.NewUserRepo()
	hashPass, err := utils.GenPassword(usr.Password)
	if err != nil {
		return nil, err
	}
	usr.Password = string(hashPass)
	usr.Type = repo.USER
	usr.Avatar = "None"
	find, _ := userRepo.GetUserByUsername(usr.UserName)
	if find == nil {
		res, err := userRepo.CreateUser(usr)
		res.Password = ""
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return nil, errors.New("username already exist")
}

func (user *UserService) ChangeAvatarByID(id string, fileHeader *multipart.FileHeader) (*model.User, error) {
	userRepo := user.Factory.NewUserRepo()
	// open header file-header
	file, errOpen := fileHeader.Open()
	if errOpen != nil {
		return nil, errOpen
	}
	cld := cloud.GetConnCloudinary()
	resp, errCld := cloud.UpdateImages(cld, file)
	if errCld != nil {
		return nil, errCld
	}
	update, erUpdate := userRepo.UpdateUser(id, "avatar", resp.SecureURL)
	if erUpdate != nil {
		return nil, erUpdate
	}
	return update, nil
}

func (user *UserService) UpdateInfo(id string, firstname string, lastname string, pass string) (*model.User, error) {
	userRepo := user.Factory.NewUserRepo()
	hashPass, err := utils.GenPassword(pass)
	if err != nil {
		return nil, err
	}
	var ret *model.User
	if ret, err = userRepo.UpdateUser(id, "password", string(hashPass)); err != nil {
		return nil, err
	}
	if ret, err = userRepo.UpdateUser(id, "firstname", firstname); err != nil {
		return nil, err
	}
	if ret, err = userRepo.UpdateUser(id, "lastname", lastname); err != nil {
		return nil, err
	}
	return ret, err
}

func (user *UserService) DeleteUser(id string) error {
	userRepo := user.Factory.NewUserRepo()
	err := userRepo.DeleteUserByID(id)
	return err
}
