package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
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
	Factory interfaces.IRepoFactory
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
		return nil, errors.New("username or password incorrect")
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
	sess, errSess := cache.GetSessionStore().Get(ctx)
	if errSess != nil {
		return nil, errSess
	}
	// Save to session
	sess.Set("Authorization", tokenString)
	err = sess.Save()
	if err != nil {
		return nil, err
	}
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

func (user *UserService) UpdateInfo(props ...string) (*model.User, error) {
	var (
		id        = props[0]
		firstname = props[1]
		lastname  = props[2]
	)
	userRepo := user.Factory.NewUserRepo()
	var ret *model.User
	var err = errors.New("")
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

func (user *UserService) UpdatePass(c *fiber.Ctx, id string) (*string, error) {
	var body struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return nil, err
	}
	userRepo := user.Factory.NewUserRepo()
	hashNewPass, errNewPass := utils.GenPassword(body.NewPassword)
	if errNewPass != nil {
		return nil, errNewPass
	}
	userOj, err := userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	err = utils.ComparePassword(userOj.Password, body.OldPassword)
	if err != nil {
		return nil, errors.New("password incorrect")
	}
	_, err = userRepo.UpdateUser(id, "password", string(hashNewPass))
	if err != nil {
		return nil, err
	}
	return &body.NewPassword, nil
}
