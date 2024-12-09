package sale

import (
	"database/sql"

	"github.com/punkestu/dodollaptoponline-go/utils"
	"github.com/punkestu/dodollaptoponline-go/utils/models"
)

type SaleRepoMysql struct {
	db *sql.DB
}

func NewSaleRepoMysql() *SaleRepoMysql {
	return &SaleRepoMysql{
		db: utils.DB(),
	}
}

func (r *SaleRepoMysql) CreateSale(userID int, sale SaleAdd) (int, error) {
	rows, err := r.db.Exec("INSERT INTO sales (product_id, user_id, quantity) VALUES (?, ?, ?)", sale.ProductID, userID, sale.Quantity)
	if err != nil {
		return -1, models.NewError("internal server error", 500)
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return -1, models.NewError("internal server error", 500)
	}

	return int(id), nil
}

func (r *SaleRepoMysql) GetMyPurchase(userID int) ([]Sale, error) {
	rows, err := r.db.Query("SELECT * FROM sales WHERE user_id = ?", userID)
	if err != nil {
		return nil, models.NewError("internal server error", 500)
	}

	sales := []Sale{}
	for rows.Next() {
		var sale Sale
		err = rows.Scan(&sale.ID, &sale.ProductID, &sale.UserID, &sale.Quantity)
		if err != nil {
			return nil, models.NewError("internal server error", 500)
		}

		sales = append(sales, sale)
	}

	return sales, nil
}

func (r *SaleRepoMysql) GetMySales(userID int) ([]Sale, error) {
	rows, err := r.db.Query("SELECT sales.* FROM sales JOIN products ON sales.product_id = products.id WHERE products.user_id = ?", userID)
	if err != nil {
		return nil, models.NewError("internal server error", 500)
	}

	sales := []Sale{}
	for rows.Next() {
		var sale Sale
		err = rows.Scan(&sale.ID, &sale.ProductID, &sale.UserID, &sale.Quantity)
		if err != nil {
			return nil, models.NewError("internal server error", 500)
		}

		sales = append(sales, sale)
	}

	return sales, nil
}

func (r *SaleRepoMysql) GetSale(id int) (*Sale, error) {
	rows, err := r.db.Query("SELECT * FROM sales WHERE id = ?", id)
	if err != nil {
		return nil, models.NewError("internal server error", 500)
	}

	if rows.Next() {
		var sale Sale
		err = rows.Scan(&sale.ID, &sale.ProductID, &sale.UserID, &sale.Quantity)
		if err != nil {
			return nil, models.NewError("internal server error", 500)
		}

		return &sale, nil
	}

	return nil, models.NewError("sale not found", 404)
}
