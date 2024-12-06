package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/punkestu/dodollaptoponline-go/internal/handlers"
	"github.com/punkestu/dodollaptoponline-go/internal/repositories"
	"github.com/punkestu/dodollaptoponline-go/internal/routes"
	"github.com/punkestu/dodollaptoponline-go/internal/services"
)

func main() {
	api := fiber.New()

	userRoutes := routes.NewUserRoutes(handlers.NewUserHandlerImpl(services.NewUserService(repositories.NewUserRepoMock())))
	api.Mount("/user", userRoutes)

	api.Listen(":3000")
}
