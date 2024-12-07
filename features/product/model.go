package product

import (
	"encoding/base64"
	"encoding/json"

	"github.com/punkestu/dodollaptoponline-go/utils/models"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	UserID      int    `json:"user_id"`
}

type ProductAdd struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

type ProductUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

type ProductUpdateStock struct {
	Stock int `json:"stock"`
}

type UserData struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func TokenToUserID(token string) (int, error) {
	var userData UserData
	jsonStr, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return -1, models.NewError("Unauthorized", 401)
	}
	if err := json.Unmarshal([]byte(jsonStr), &userData); err != nil {
		return -1, models.NewError("Unauthorized", 401)
	}

	return userData.ID, nil
}
