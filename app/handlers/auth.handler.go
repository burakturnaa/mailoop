package handlers

import (
	"net/http"

	"github.com/burakturnaa/mailoop.git/app/dto"
	"github.com/burakturnaa/mailoop.git/app/services"
	"github.com/burakturnaa/mailoop.git/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type authHandler struct {
	authService services.AuthService
	jwtService  services.JWTService
	userService services.UserService
}

func NewAuthHandler(
	authService services.AuthService,
	jwtService services.JWTService,
	userService services.UserService,
) AuthHandler {
	return &authHandler{
		authService: authService,
		jwtService:  jwtService,
		userService: userService,
	}
}

func (h *authHandler) Login(ctx *fiber.Ctx) error {
	var loginRequest dto.LoginBody
	err := ctx.BodyParser(&loginRequest)
	if err != nil {
		response := utils.BuildResponse(4001, "Invalid request body format", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	err = h.authService.VerifyCredential(loginRequest.Email, loginRequest.Password)
	if err != nil {
		response := utils.BuildResponse(4011, "Unauthorized", nil, nil)
		return ctx.Status(http.StatusUnauthorized).JSON(response)
	}

	user, _ := h.userService.FindUserByEmail(loginRequest.Email)

	token := h.jwtService.GenerateToken(user.Id)
	user.Token = token

	user.Id = nil
	response := utils.BuildResponse(2001, "success", nil, user)
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *authHandler) Register(ctx *fiber.Ctx) error {
	var registerRequest dto.RegisterBody
	err := ctx.BodyParser(&registerRequest)
	if err != nil {
		response := utils.BuildResponse(4001, "Invalid request body format", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	user, err := h.userService.CreateUser(registerRequest)
	if err != nil { // user already exists
		response := utils.BuildResponse(4041, "user already exists", nil, nil)
		return ctx.Status(http.StatusConflict).JSON(response)
	}

	// token := h.jwtService.GenerateToken(user.Id)
	// user.Token = token

	response := utils.BuildResponse(2001, "success", nil, user)
	return ctx.Status(http.StatusOK).JSON(response)
}
