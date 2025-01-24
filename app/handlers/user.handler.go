package handlers

import (
	"net/http"

	"github.com/burakturnaa/mailoop.git/app/services"
	"github.com/burakturnaa/mailoop.git/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	CheckToken(ctx *fiber.Ctx) error
}

type userHandler struct {
	userService services.UserService
	jwtService  services.JWTService
}

func NewUserHandler(
	userService services.UserService,
	jwtService services.JWTService,
) UserHandler {
	return &userHandler{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (h *userHandler) CheckToken(ctx *fiber.Ctx) error {
	response := utils.BuildResponse(2001, "success", nil, nil)
	return ctx.Status(http.StatusOK).JSON(response)
}
