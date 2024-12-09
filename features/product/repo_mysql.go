package product

import (
	"database/sql"

	"github.com/punkestu/dodollaptoponline-go/utils"
	"github.com/punkestu/dodollaptoponline-go/utils/models"
)

type ProductRepositoryMysql struct {
	db *sql.DB
}

func NewProductRepositoryMysql() *ProductRepositoryMysql {
	return &ProductRepositoryMysql{
		db: utils.DB(),
	}
}

func (r *ProductRepositoryMysql) GetProducts() ([]Product, error) {
	rows, err := r.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, models.NewError("internal server error", 500)
	}

	products := []Product{}
	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.UserID)
		if err != nil {
			return nil, models.NewError("internal server error", 500)
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepositoryMysql) GetProduct(id int) (Product, error) {
	rows, err := r.db.Query("SELECT * FROM products WHERE id = ?", id)
	if err != nil {
		return Product{}, models.NewError("internal server error", 500)
	}

	if rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.UserID)
		if err != nil {
			return Product{}, models.NewError("internal server error", 500)
		}

		return product, nil
	}

	return Product{}, models.NewError("Product not found", 404)
}

func (r *ProductRepositoryMysql) AddProduct(userID int, p ProductAdd) (Product, error) {
	rows, err := r.db.Exec("INSERT INTO products (name, description, price, stock, user_id) VALUES (?, ?, ?, ?, ?)", p.Name, p.Description, p.Price, p.Stock, userID)
	if err != nil {
		return Product{}, models.NewError("internal server error", 500)
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return Product{}, models.NewError("internal server error", 500)
	}

	product, err := r.GetProduct(int(id))
	if err != nil {
		return Product{}, models.NewError("internal server error", 500)
	}

	return product, nil
}

func (r *ProductRepositoryMysql) UpdateProduct(id int, p ProductUpdate) (Product, error) {
	rows, err := r.db.Exec("UPDATE products SET name = ?, description = ?, price = ?, stock = ? WHERE id = ?", p.Name, p.Description, p.Price, p.Stock, id)
	if err != nil {
		return Product{}, models.NewError("internal server error", 500)
	}

	_, err = rows.RowsAffected()
	if err != nil {
		return Product{}, models.NewError("internal server error", 500)
	}

	product, err := r.GetProduct(id)
	if err != nil {
		return Product{}, models.NewError("internal server error", 500)
	}

	return product, nil
}

func (r *ProductRepositoryMysql) DeleteProduct(id int) error {
	rows, err := r.db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return models.NewError("internal server error", 500)
	}

	_, err = rows.RowsAffected()
	if err != nil {
		return models.NewError("internal server error", 500)
	}

	return nil
}

func (r *ProductRepositoryMysql) UpdateStock(id int, stock int) (Product, error) {
	rows, err := r.db.Exec("UPDATE products SET stock = ? WHERE id = ?", stock, id)
	if err != nil {
		return Product{}, models.NewError("internal server error", 500)
	}

	_, err = rows.RowsAffected()
	if err != nil {
		return Product{}, models.NewError("internal server error", 500)
	}

	product, err := r.GetProduct(id)
	if err != nil {
		return Product{}, models.NewError("internal server error", 500)
	}

	return product, nil
}
