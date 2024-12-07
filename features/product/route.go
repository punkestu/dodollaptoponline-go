package product

import "github.com/gofiber/fiber/v2"

type ProductHandler interface {
	GetProducts(c *fiber.Ctx) error
	GetProduct(c *fiber.Ctx) error
	AddProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
	UpdateStock(c *fiber.Ctx) error
}

func NewProductRoutes(h ProductHandler) *fiber.App {
	productRoute := fiber.New()
	productRoute.Get("/", h.GetProducts)
	productRoute.Get("/:id", h.GetProduct)
	productRoute.Post("/", h.AddProduct)
	productRoute.Put("/:id", h.UpdateProduct)
	productRoute.Delete("/:id", h.DeleteProduct)
	productRoute.Patch("/:id/stock", h.UpdateStock)

	return productRoute
}
