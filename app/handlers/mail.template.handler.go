package handlers

import (
	"net/http"

	"github.com/burakturnaa/mailoop.git/app/dto"
	"github.com/burakturnaa/mailoop.git/app/services"
	"github.com/burakturnaa/mailoop.git/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MailTemplateHandler interface {
	CreateMailTemplate(ctx *fiber.Ctx) error
}

type mailTemplateHandler struct {
	mailTemplateService services.MailTemplateService
	userService         services.UserService
	jwtService          services.JWTService
}

func NewMailTemplateHandler(
	mailTemplateService services.MailTemplateService,
	userService services.UserService,
	jwtService services.JWTService,
) MailTemplateHandler {
	return &mailTemplateHandler{
		mailTemplateService: mailTemplateService,
		userService:         userService,
		jwtService:          jwtService,
	}
}

func (h *mailTemplateHandler) CreateMailTemplate(ctx *fiber.Ctx) error {
	// check the user id in the token
	userId, _ := primitive.ObjectIDFromHex(ctx.Locals("userIdClaims").(string))
	user, _ := h.userService.FindUserByID(userId)
	if user == nil {
		response := utils.BuildResponse(4041, "user not found", nil, nil)
		return ctx.Status(http.StatusUnauthorized).JSON(response)
	}

	var mailTemplateRequest dto.MailTemplateBody
	err := ctx.BodyParser(&mailTemplateRequest)
	if err != nil {
		response := utils.BuildResponse(4001, "Invalid request body format", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}
	// clean html and string content via SanitizeHTML()
	mailTemplateRequest.Content = utils.SanitizeHTML(mailTemplateRequest.Content)

	mailTemplate, err := h.mailTemplateService.CreateMailTemplate(mailTemplateRequest)
	mailTemplate.Id = nil
	if err != nil {
		response := utils.BuildResponse(5001, "database error", nil, nil)
		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	response := utils.BuildResponse(2001, "success", nil, mailTemplate)
	return ctx.Status(http.StatusOK).JSON(response)
}
