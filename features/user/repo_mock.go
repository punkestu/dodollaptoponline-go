package user

import "github.com/punkestu/dodollaptoponline-go/utils/models"

type UserRepoMockImpl struct {
	users   []User
	counter int
}

func NewUserRepoMock() *UserRepoMockImpl {
	return &UserRepoMockImpl{
		users:   []User{},
		counter: 0,
	}
}

func (u *UserRepoMockImpl) GetUsers() ([]UserProfile, error) {
	userProfiles := make([]UserProfile, len(u.users))

	for i, user := range u.users {
		userProfiles[i] = user.ToUserProfile()
	}

	return userProfiles, nil
}

func (u *UserRepoMockImpl) GetUserByID(id int) (*User, error) {
	for _, user := range u.users {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, models.NewError("user not found", 404)
}

func (u *UserRepoMockImpl) GetUserByUsername(username string) (*User, error) {
	for _, user := range u.users {
		if user.Username == username {
			return &user, nil
		}
	}

	return nil, models.NewError("user not found", 404)
}

func (u *UserRepoMockImpl) InsertUser(user UserRegister) (int, error) {
	newUser := User{
		ID:       u.counter,
		Username: user.Username,
		Password: user.Password,
	}

	u.users = append(u.users, newUser)
	u.counter++

	return newUser.ID, nil
}
