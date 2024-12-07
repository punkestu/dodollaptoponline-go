package sale

import "github.com/gofiber/fiber/v2"

func Init() *fiber.App {
	saleRoutes := NewSaleRouter(NewSaleHandler(NewSaleService(NewSaleRepoMock())))

	return saleRoutes
}
