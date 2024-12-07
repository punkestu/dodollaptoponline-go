package sale

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/punkestu/dodollaptoponline-go/config"
	"github.com/punkestu/dodollaptoponline-go/utils/models"
)

type SaleRepoMock struct {
	sales    []Sale
	products []SaleProduct
	counter  int
}

// CreateSale(userID int, sale SaleAdd) (int, error)
// GetMyPurchase(userID int) ([]Sale, error)
// GetMySales(userID int) ([]Sale, error)
// GetSale(id int) (*Sale, error)
func NewSaleRepoMock() *SaleRepoMock {
	return &SaleRepoMock{
		sales:    []Sale{},
		products: []SaleProduct{},
		counter:  0,
	}
}

func (r *SaleRepoMock) CreateSale(userID int, sale SaleAdd) (int, error) {
	err := r.CacheProduct()
	if err != nil {
		fmt.Printf("failed to cache product: %v", err)
		return -1, models.NewError("Failed to cache product", 500)
	}

	r.counter++
	newSale := Sale{
		ID:        r.counter,
		ProductID: sale.ProductID,
		UserID:    userID,
		Quantity:  sale.Quantity,
	}
	r.sales = append(r.sales, newSale)

	return r.counter, nil
}

func (r *SaleRepoMock) CacheProduct() error {
	resp, err := http.Get(config.GetServiceDomain("product"))
	if err != nil {
		return models.NewError(fmt.Sprintf("error when fetching product %v", err), 500)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return models.NewError(fmt.Sprintf("error when fetching product %v", resp.Status), 500)
	}

	var body []byte
	if body, err = io.ReadAll(resp.Body); err != nil {
		return models.NewError(fmt.Sprintf("error when read body %v", err), 500)
	}

	var apiProduct APIProduct
	if err := json.Unmarshal(body, &apiProduct); err != nil {
		return models.NewError(fmt.Sprintf("error when unmarshal %v", err), 500)
	}

	r.products = apiProduct.Data

	return nil
}

func (r *SaleRepoMock) GetMyPurchase(userID int) ([]Sale, error) {
	var sales []Sale = []Sale{}
	for _, sale := range r.sales {
		if sale.UserID == userID {
			sales = append(sales, sale)
		}
	}

	return sales, nil
}

func (r *SaleRepoMock) GetMySales(userID int) ([]Sale, error) {
	err := r.CacheProduct()
	if err != nil {
		fmt.Printf("failed to cache product: %v", err)
		return []Sale{}, models.NewError("Failed to cache product", 500)
	}

	// find product with user id
	var productIDs []int
	for _, product := range r.products {
		if product.UserID == userID {
			productIDs = append(productIDs, product.ID)
		}
	}

	// find sales with product id
	var sales []Sale = []Sale{}
	for _, sale := range r.sales {
		for _, productID := range productIDs {
			if sale.ProductID == productID {
				sales = append(sales, sale)
				break
			}
		}
	}

	return sales, nil
}

func (r *SaleRepoMock) GetSale(id int) (*Sale, error) {
	for _, sale := range r.sales {
		if sale.ID == id {
			return &sale, nil
		}
	}

	return nil, nil
}
