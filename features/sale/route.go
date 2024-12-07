package sale

import "github.com/gofiber/fiber/v2"

type SaleHandler interface {
	CreateSale(c *fiber.Ctx) error
	GetMyPurchase(c *fiber.Ctx) error
	GetMySales(c *fiber.Ctx) error
	GetSale(c *fiber.Ctx) error
}

func NewSaleRouter(saleHandler SaleHandler) *fiber.App {
	saleRouter := fiber.New()
	saleRouter.Post("/", saleHandler.CreateSale)
	saleRouter.Get("/purchase", saleHandler.GetMyPurchase)
	saleRouter.Get("/sales", saleHandler.GetMySales)
	saleRouter.Get("/:id", saleHandler.GetSale)

	return saleRouter
}
