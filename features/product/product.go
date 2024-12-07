package product

import (
	"github.com/gofiber/fiber/v2"
)

func Init() *fiber.App {
	api := fiber.New()

	productRoutes := NewProductRoutes(NewProductHandlerImpl(NewProductServiceImpl(NewProductRepositoryMock())))
	api.Mount("/product", productRoutes)

	return api
}
