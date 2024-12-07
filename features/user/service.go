package user

import "github.com/punkestu/dodollaptoponline-go/utils/models"

type UserRepository interface {
	GetUsers() ([]UserProfile, error)
	GetUserByID(id int) (*User, error)
	GetUserByUsername(username string) (*User, error)
	InsertUser(user UserRegister) (int, error)
}

type UserServiceImpl struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		repo: repo,
	}
}

func (u *UserServiceImpl) Login(credentials UserLogin) (*UserProfile, error) {
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

func (u *UserServiceImpl) Register(user UserRegister) (int, error) {
	_, err := u.repo.GetUserByUsername(user.Username)
	if err == nil {
		return 0, models.NewError("username already taken", 400)
	}

	return u.repo.InsertUser(user)
}

func (u *UserServiceImpl) GetProfile(id int) (*UserProfile, error) {
	user, err := u.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	userProfile := user.ToUserProfile()
	return &userProfile, nil
}
