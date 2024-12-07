package sale

import "github.com/punkestu/dodollaptoponline-go/utils/models"

type Sale struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	UserID    int `json:"user_id"`
	Quantity  int `json:"quantity"`
}

type SaleAdd struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type SaleProduct struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
}

type APIProduct struct {
	Data []SaleProduct       `json:"data"`
	Meta models.MetaResponse `json:"meta"`
}
