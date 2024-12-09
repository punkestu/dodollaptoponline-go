package user

import (
	"database/sql"

	"github.com/punkestu/dodollaptoponline-go/utils"
	"github.com/punkestu/dodollaptoponline-go/utils/models"
)

type UserRepoMysqlImpl struct {
	db *sql.DB
}

func NewUserRepoMysql() *UserRepoMysqlImpl {
	return &UserRepoMysqlImpl{
		db: utils.DB(),
	}
}

func (u *UserRepoMysqlImpl) GetUsers() ([]UserProfile, error) {
	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, models.NewError("internal server error", 500)
	}

	userProfiles := []UserProfile{}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			return nil, models.NewError("internal server error", 500)
		}

		userProfiles = append(userProfiles, user.ToUserProfile())
	}

	return userProfiles, nil
}

func (u *UserRepoMysqlImpl) GetUserByID(id int) (*User, error) {
	rows, err := u.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, models.NewError("internal server error", 500)
	}

	if rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			return nil, models.NewError("internal server error", 500)
		}

		return &user, nil
	}

	return nil, models.NewError("user not found", 404)
}

func (u *UserRepoMysqlImpl) GetUserByUsername(username string) (*User, error) {
	rows, err := u.db.Query("SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		return nil, models.NewError("internal server error", 500)
	}

	if rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			return nil, models.NewError("internal server error", 500)
		}

		return &user, nil
	}

	return nil, models.NewError("user not found", 404)
}

func (u *UserRepoMysqlImpl) InsertUser(user UserRegister) (int, error) {
	rows, err := u.db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		return 0, models.NewError("internal server error", 500)
	}

	lastInsertedId, err := rows.LastInsertId()
	if err != nil {
		return 0, models.NewError("internal server error", 500)
	}

	return int(lastInsertedId), nil
}
