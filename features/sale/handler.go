package sale

import (
	"github.com/gofiber/fiber/v2"
	"github.com/punkestu/dodollaptoponline-go/utils/models"
)

type SaleService interface {
	CreateSale(token string, sale SaleAdd) (int, error)
	GetMyPurchase(token string) ([]Sale, error)
	GetMySales(token string) ([]Sale, error)
	GetSale(id int) (*Sale, error)
}

type SaleHandlerImpl struct {
	SaleService SaleService
}

func NewSaleHandler(saleService SaleService) *SaleHandlerImpl {
	return &SaleHandlerImpl{
		SaleService: saleService,
	}
}

func (h *SaleHandlerImpl) CreateSale(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		builtinError := models.NewError("Token is required", 400)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	var sale SaleAdd
	if err := c.BodyParser(&sale); err != nil {
		builtinError := models.NewError("Invalid request body", 400)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	id, err := h.SaleService.CreateSale(token, sale)
	if err != nil {
		builtinError := models.ToError(err)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	response := models.NewSuccessResponse(201, "Sale created", &fiber.Map{"id": id}, nil)

	return c.Status(201).JSON(response)
}

func (h *SaleHandlerImpl) GetMyPurchase(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		builtinError := models.NewError("Token is required", 400)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	sales, err := h.SaleService.GetMyPurchase(token)
	if err != nil {
		builtinError := models.ToError(err)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	response := models.NewSuccessResponse(200, "My purchase", &fiber.Map{"sales": sales}, &models.MetaResponse{
		Page:  0,
		Limit: -1,
		Total: len(sales),
	})

	return c.Status(200).JSON(response)
}

func (h *SaleHandlerImpl) GetMySales(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		builtinError := models.NewError("Token is required", 400)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	sales, err := h.SaleService.GetMySales(token)
	if err != nil {
		builtinError := models.ToError(err)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	response := models.NewSuccessResponse(200, "My sales", &fiber.Map{"sales": sales}, &models.MetaResponse{
		Page:  0,
		Limit: -1,
		Total: len(sales),
	})

	return c.Status(200).JSON(response)
}

func (h *SaleHandlerImpl) GetSale(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		builtinError := models.NewError("Invalid sale ID", 400)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	sale, err := h.SaleService.GetSale(id)
	if err != nil {
		builtinError := models.ToError(err)
		return c.Status(builtinError.Code).JSON(builtinError)
	}

	response := models.NewSuccessResponse(200, "Sale detail", sale, nil)

	return c.Status(200).JSON(response)
}
