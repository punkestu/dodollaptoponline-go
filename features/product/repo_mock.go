package product

import "github.com/punkestu/dodollaptoponline-go/utils/models"

type ProductRepositoryMock struct {
	Products []Product
	counter  int
}

func NewProductRepositoryMock() *ProductRepositoryMock {
	return &ProductRepositoryMock{
		Products: []Product{},
		counter:  0,
	}
}

func (r *ProductRepositoryMock) GetProducts() ([]Product, error) {
	return r.Products, nil
}

func (r *ProductRepositoryMock) GetProduct(id int) (Product, error) {
	for _, product := range r.Products {
		if product.ID == id {
			return product, nil
		}
	}

	return Product{}, models.NewError("Product not found", 404)
}

func (r *ProductRepositoryMock) AddProduct(userID int, p ProductAdd) (Product, error) {
	r.counter++
	product := Product{
		ID:          r.counter,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		UserID:      userID,
	}

	r.Products = append(r.Products, product)

	return product, nil
}

func (r *ProductRepositoryMock) UpdateProduct(id int, p ProductUpdate) (Product, error) {
	for i, product := range r.Products {
		if product.ID == id {
			product.Name = p.Name
			product.Description = p.Description
			product.Price = p.Price
			product.Stock = p.Stock

			r.Products[i] = product

			return product, nil
		}
	}

	return Product{}, models.NewError("Product not found", 404)
}

func (r *ProductRepositoryMock) DeleteProduct(id int) error {
	for i, product := range r.Products {
		if product.ID == id {
			r.Products = append(r.Products[:i], r.Products[i+1:]...)
			return nil
		}
	}

	return models.NewError("Product not found", 404)
}

func (r *ProductRepositoryMock) UpdateStock(id int, stock int) (Product, error) {
	for i, product := range r.Products {
		if product.ID == id {
			product.Stock = stock

			r.Products[i] = product

			return product, nil
		}
	}

	return Product{}, models.NewError("Product not found", 404)
}
