package user

import (
	"github.com/gofiber/fiber/v2"
)

func Init() *fiber.App {
	api := fiber.New()

	productRoutes := NewUserRoutes(NewUserHandlerImpl(NewUserService(NewUserRepoMock())))
	api.Mount("/user", productRoutes)

	return api
}
