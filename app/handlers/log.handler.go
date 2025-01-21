package handlers

import (
	"net/http"

	"github.com/burakturnaa/mailoop.git/app/services"
	"github.com/burakturnaa/mailoop.git/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogHandler interface {
	GetAll(ctx *fiber.Ctx) error
	GetOne(ctx *fiber.Ctx) error
}

type logHandler struct {
	logService  services.LogService
	userService services.UserService
	jwtService  services.JWTService
}

func NewLogHandler(
	logService services.LogService,
	userService services.UserService,
	jwtService services.JWTService,
) LogHandler {
	return &logHandler{
		logService:  logService,
		userService: userService,
		jwtService:  jwtService,
	}
}

func (h *logHandler) GetAll(ctx *fiber.Ctx) error {
	logs, err := h.logService.GetAll()
	if err != nil {
		if err.Error() != mongo.ErrNoDocuments.Error() {
			response := utils.BuildResponse(5001, "database error", nil, nil)
			return ctx.Status(http.StatusInternalServerError).JSON(response)
		}
	}
	response := utils.BuildResponse(2001, "success", nil, logs)
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *logHandler) GetOne(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		response := utils.BuildResponse(4002, "Validation error", fiber.Map{"id": "id must be a valid primitive object id"}, nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}
	logs, err := h.logService.GetOne(id)
	if err != nil {
		if err.Error() != mongo.ErrNoDocuments.Error() {
			response := utils.BuildResponse(5001, "database error", nil, nil)
			return ctx.Status(http.StatusInternalServerError).JSON(response)
		} else {
			response := utils.BuildResponse(4041, "log not found", nil, nil)
			return ctx.Status(http.StatusNotFound).JSON(response)
		}
	}
	response := utils.BuildResponse(2001, "success", nil, logs)
	return ctx.Status(http.StatusOK).JSON(response)
}
