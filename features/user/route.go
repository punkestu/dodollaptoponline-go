package user

import (
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	GetProfile(c *fiber.Ctx) error
}

func NewUserRoutes(handlers UserHandler) *fiber.App {
	userRoute := fiber.New()
	userRoute.Post("/login", handlers.Login)
	userRoute.Post("/register", handlers.Register)
	userRoute.Get("/profile/:id", handlers.GetProfile)

	return userRoute
}
