package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/punkestu/dodollaptoponline-go/utils/models"
)

type ProductService interface {
	GetProducts() ([]Product, error)
	GetProduct(id int) (Product, error)
	AddProduct(token string, p ProductAdd) (Product, error)
	UpdateProduct(token string, id int, p ProductUpdate) (Product, error)
	DeleteProduct(token string, id int) error
	UpdateStock(token string, id int, stock int) (Product, error)
}

type ProductHandlerImpl struct {
	Service ProductService
}

func NewProductHandlerImpl(s ProductService) *ProductHandlerImpl {
	return &ProductHandlerImpl{
		Service: s,
	}
}

func (h *ProductHandlerImpl) GetProducts(c *fiber.Ctx) error {
	products, err := h.Service.GetProducts()
	if err != nil {
		builtinError := models.ToError(err)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	response := models.NewSuccessResponse(200, "Successfully get products", products, &models.MetaResponse{
		Page:  0,
		Limit: -1,
		Total: len(products),
	})

	return c.JSON(response)
}

func (h *ProductHandlerImpl) GetProduct(c *fiber.Ctx) error {
	productID, err := c.ParamsInt("id")
	if err != nil {
		builtinError := models.NewError("Invalid product ID", 400)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	product, err := h.Service.GetProduct(productID)
	if err != nil {
		builtinError := models.ToError(err)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	response := models.NewSuccessResponse(200, "Successfully get product", product, nil)

	return c.JSON(response)
}

func (h *ProductHandlerImpl) AddProduct(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		builtinError := models.NewError("Unauthorized", 401)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	var product ProductAdd
	if err := c.BodyParser(&product); err != nil {
		builtinError := models.NewError("Invalid request body", 400)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	newProduct, err := h.Service.AddProduct(token, product)
	if err != nil {
		builtinError := models.ToError(err)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	response := models.NewSuccessResponse(201, "Successfully add product", newProduct, nil)

	return c.JSON(response)
}

func (h *ProductHandlerImpl) UpdateProduct(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		builtinError := models.NewError("Unauthorized", 401)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	productID, err := c.ParamsInt("id")
	if err != nil {
		builtinError := models.NewError("Invalid product ID", 400)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	var product ProductUpdate
	if err := c.BodyParser(&product); err != nil {
		builtinError := models.NewError("Invalid request body", 400)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	updatedProduct, err := h.Service.UpdateProduct(token, productID, product)
	if err != nil {
		builtinError := models.ToError(err)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	response := models.NewSuccessResponse(200, "Successfully update product", updatedProduct, nil)

	return c.JSON(response)
}

func (h *ProductHandlerImpl) DeleteProduct(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		builtinError := models.NewError("Unauthorized", 401)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	productID, err := c.ParamsInt("id")
	if err != nil {
		builtinError := models.NewError("Invalid product ID", 400)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	err = h.Service.DeleteProduct(token, productID)
	if err != nil {
		builtinError := models.ToError(err)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	response := models.NewSuccessResponse(200, "Successfully delete product", nil, nil)

	return c.JSON(response)
}

func (h *ProductHandlerImpl) UpdateStock(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		builtinError := models.NewError("Unauthorized", 401)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	productID, err := c.ParamsInt("id")
	if err != nil {
		builtinError := models.NewError("Invalid product ID", 400)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	var stock ProductUpdateStock
	if err := c.BodyParser(&stock); err != nil {
		builtinError := models.NewError("Invalid request body", 400)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	updatedProduct, err := h.Service.UpdateStock(token, productID, stock.Stock)
	if err != nil {
		builtinError := models.ToError(err)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	response := models.NewSuccessResponse(200, "Successfully update stock", updatedProduct, nil)

	return c.JSON(response)
}
