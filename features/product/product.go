package product

import (
	"github.com/gofiber/fiber/v2"
)

func Init() *fiber.App {
	productRoutes := NewProductRoutes(NewProductHandlerImpl(NewProductServiceImpl(NewProductRepositoryMysql())))

	return productRoutes
}
