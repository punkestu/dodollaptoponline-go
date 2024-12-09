package user

import (
	"github.com/gofiber/fiber/v2"
)

func Init() *fiber.App {
	userRoutes := NewUserRoutes(NewUserHandlerImpl(NewUserService(NewUserRepoMysql())))

	return userRoutes
}
