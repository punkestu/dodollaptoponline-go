package models

import (
	"encoding/base64"
	"encoding/json"
)

type UserData struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func TokenToUserID(token string) (int, error) {
	var userData UserData
	jsonStr, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return -1, NewError("Unauthorized", 401)
	}
	if err := json.Unmarshal([]byte(jsonStr), &userData); err != nil {
		return -1, NewError("Unauthorized", 401)
	}

	return userData.ID, nil
}
