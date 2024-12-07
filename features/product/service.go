package product

import "github.com/punkestu/dodollaptoponline-go/utils/models"

type ProductRepository interface {
	GetProducts() ([]Product, error)
	GetProduct(id int) (Product, error)
	AddProduct(userID int, p ProductAdd) (Product, error)
	UpdateProduct(id int, p ProductUpdate) (Product, error)
	DeleteProduct(id int) error
	UpdateStock(id int, stock int) (Product, error)
}

type ProductServiceImpl struct {
	Repository ProductRepository
}

func NewProductServiceImpl(r ProductRepository) *ProductServiceImpl {
	return &ProductServiceImpl{
		Repository: r,
	}
}

func (s *ProductServiceImpl) GetProducts() ([]Product, error) {
	return s.Repository.GetProducts()
}

func (s *ProductServiceImpl) GetProduct(id int) (Product, error) {
	return s.Repository.GetProduct(id)
}

func (s *ProductServiceImpl) AddProduct(token string, p ProductAdd) (Product, error) {
	userID, err := models.TokenToUserID(token)
	if err != nil {
		return Product{}, err
	}

	return s.Repository.AddProduct(userID, p)
}

func (s *ProductServiceImpl) UpdateProduct(token string, id int, p ProductUpdate) (Product, error) {
	userID, err := models.TokenToUserID(token)
	if err != nil {
		return Product{}, err
	}

	product, err := s.Repository.GetProduct(id)
	if err != nil {
		return Product{}, err
	}

	if product.UserID != userID {
		return Product{}, models.NewError("Unauthorized", 401)
	}

	return s.Repository.UpdateProduct(id, p)
}

func (s *ProductServiceImpl) DeleteProduct(token string, id int) error {
	userID, err := models.TokenToUserID(token)
	if err != nil {
		return err
	}

	product, err := s.Repository.GetProduct(id)
	if err != nil {
		return err
	}

	if product.UserID != userID {
		return models.NewError("Unauthorized", 401)
	}

	return s.Repository.DeleteProduct(id)
}

func (s *ProductServiceImpl) UpdateStock(token string, id int, stock int) (Product, error) {
	userID, err := models.TokenToUserID(token)
	if err != nil {
		return Product{}, err
	}

	product, err := s.Repository.GetProduct(id)
	if err != nil {
		return Product{}, err
	}

	if product.UserID != userID {
		return Product{}, models.NewError("Unauthorized", 401)
	}

	return s.Repository.UpdateStock(id, stock)
}
