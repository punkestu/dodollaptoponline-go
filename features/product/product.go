package product

import (
	"github.com/gofiber/fiber/v2"
)

func Init() *fiber.App {
	productRoutes := NewProductRoutes(NewProductHandlerImpl(NewProductServiceImpl(NewProductRepositoryMock())))

	return productRoutes
}
