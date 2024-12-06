package services

import (
	"github.com/punkestu/dodollaptoponline-go/internal/models"
)

type UserRepository interface {
	GetUsers() ([]models.UserProfile, error)
	GetUserByID(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	InsertUser(user models.UserRegister) (int, error)
}

type UserServiceImpl struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (u *UserServiceImpl) Login(credentials models.UserLogin) (*models.UserProfile, error) {
	user, err := u.repo.GetUserByUsername(credentials.Username)
	if err != nil {
		return nil, err
	}

	if user.Password != credentials.Password {
		return nil, models.NewError("invalid password", 401)
	}

	userProfile := user.ToUserProfile()

	return &userProfile, nil
}

func (u *UserServiceImpl) Register(user models.UserRegister) (int, error) {
	_, err := u.repo.GetUserByUsername(user.Username)
	if err == nil {
		return 0, models.NewError("username already taken", 400)
	}

	return u.repo.InsertUser(user)
}

func (u *UserServiceImpl) GetProfile(id int) (*models.UserProfile, error) {
	user, err := u.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	userProfile := user.ToUserProfile()
	return &userProfile, nil
}
