package repositories

import (
	"github.com/punkestu/dodollaptoponline-go/internal/models"
)

type UserRepoMockImpl struct {
	users   []models.User
	counter int
}

func NewUserRepoMock() *UserRepoMockImpl {
	return &UserRepoMockImpl{
		users:   []models.User{},
		counter: 0,
	}
}

func (u *UserRepoMockImpl) GetUsers() ([]models.UserProfile, error) {
	userProfiles := make([]models.UserProfile, len(u.users))

	for i, user := range u.users {
		userProfiles[i] = user.ToUserProfile()
	}

	return userProfiles, nil
}

func (u *UserRepoMockImpl) GetUserByID(id int) (*models.User, error) {
	for _, user := range u.users {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, models.NewError("user not found", 404)
}

func (u *UserRepoMockImpl) GetUserByUsername(username string) (*models.User, error) {
	for _, user := range u.users {
		if user.Username == username {
			return &user, nil
		}
	}

	return nil, models.NewError("user not found", 404)
}

func (u *UserRepoMockImpl) InsertUser(user models.UserRegister) (int, error) {
	newUser := models.User{
		ID:       u.counter,
		Username: user.Username,
		Password: user.Password,
	}

	u.users = append(u.users, newUser)
	u.counter++

	return newUser.ID, nil
}
