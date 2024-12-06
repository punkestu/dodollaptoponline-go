package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/punkestu/dodollaptoponline-go/internal/models"
)

type UserService interface {
	Login(credentials models.UserLogin) (*models.UserProfile, error)
	Register(user models.UserRegister) (int, error)
	GetProfile(id int) (*models.UserProfile, error)
}

type UserHandlerImpl struct {
	service UserService
}

func NewUserHandlerImpl(service UserService) *UserHandlerImpl {
	return &UserHandlerImpl{
		service: service,
	}
}

func (u *UserHandlerImpl) Login(c *fiber.Ctx) error {
	credentials := new(models.UserLogin)
	if err := c.BodyParser(credentials); err != nil {
		builtinErr := models.NewError("payload error", 400)
		return c.Status(builtinErr.Code).JSON(builtinErr)
	}

	user, err := u.service.Login(*credentials)
	if err != nil {
		builtinErr := models.ToError(err)
		return c.Status(builtinErr.Code).JSON(builtinErr)
	}

	response := models.NewSuccessResponse(200, "login success", user)
	return c.JSON(response)
}

func (u *UserHandlerImpl) Register(c *fiber.Ctx) error {
	user := new(models.UserRegister)
	if err := c.BodyParser(user); err != nil {
		builtinErr := models.NewError("payload error", 400)
		return c.Status(builtinErr.Code).JSON(builtinErr)
	}

	id, err := u.service.Register(*user)
	if err != nil {
		builtinErr := models.ToError(err)
		return c.Status(builtinErr.Code).JSON(builtinErr)
	}

	response := models.NewSuccessResponse(201, "register success", fiber.Map{
		"createdId": id,
	})

	return c.JSON(response)
}

func (u *UserHandlerImpl) GetProfile(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		builtinErr := models.NewError("payload error", 400)
		return c.Status(builtinErr.Code).JSON(builtinErr)
	}

	userProfile, err := u.service.GetProfile(id)
	if err != nil {
		builtinErr := models.ToError(err)
		return c.Status(builtinErr.Code).JSON(builtinErr)
	}

	response := models.NewSuccessResponse(200, "get profile success", userProfile)

	return c.JSON(response)
}
