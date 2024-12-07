package sale

import "github.com/punkestu/dodollaptoponline-go/utils/models"

type SaleRepo interface {
	CreateSale(userID int, sale SaleAdd) (int, error)
	GetMyPurchase(userID int) ([]Sale, error)
	GetMySales(userID int) ([]Sale, error)
	GetSale(id int) (*Sale, error)
}

type SaleServiceImpl struct {
	SaleRepo SaleRepo
}

func NewSaleService(saleRepo SaleRepo) *SaleServiceImpl {
	return &SaleServiceImpl{
		SaleRepo: saleRepo,
	}
}

func (s *SaleServiceImpl) CreateSale(token string, sale SaleAdd) (int, error) {
	userID, err := models.TokenToUserID(token)
	if err != nil {
		return -1, err
	}

	saleID, err := s.SaleRepo.CreateSale(userID, sale)
	if err != nil {
		return -1, err
	}

	return saleID, nil
}

func (s *SaleServiceImpl) GetMyPurchase(token string) ([]Sale, error) {
	userID, err := models.TokenToUserID(token)
	if err != nil {
		return nil, err
	}

	sales, err := s.SaleRepo.GetMyPurchase(userID)
	if err != nil {
		return nil, err
	}

	return sales, nil
}

func (s *SaleServiceImpl) GetMySales(token string) ([]Sale, error) {
	userID, err := models.TokenToUserID(token)
	if err != nil {
		return nil, err
	}

	sales, err := s.SaleRepo.GetMySales(userID)
	if err != nil {
		return nil, err
	}

	return sales, nil
}

func (s *SaleServiceImpl) GetSale(id int) (*Sale, error) {
	sale, err := s.SaleRepo.GetSale(id)
	if err != nil {
		return nil, err
	}

	if sale == nil {
		return nil, models.NewError("Sale not found", 404)
	}

	return sale, nil
}
