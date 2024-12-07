package user

import (
	"encoding/base64"
	"encoding/json"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserProfile struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (u *User) ToUserProfile() UserProfile {
	return UserProfile{
		ID:       u.ID,
		Username: u.Username,
	}
}

func (up *UserProfile) GetToken() string {
	jsonValue, err := json.Marshal(up)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString([]byte(jsonValue))
}
