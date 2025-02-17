package handlers

import (
	"log"
	"net/http"

	"github.com/burakturnaa/mailoop.git/app/dto"
	"github.com/burakturnaa/mailoop.git/app/services"
	"github.com/burakturnaa/mailoop.git/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MailTemplateHandler interface {
	GetAll(ctx *fiber.Ctx) error
	GetOne(ctx *fiber.Ctx) error
	CreateMailTemplate(ctx *fiber.Ctx) error
	UpdateMailTemplate(ctx *fiber.Ctx) error
	DeleteMailTemplate(ctx *fiber.Ctx) error
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

func (h *mailTemplateHandler) GetAll(ctx *fiber.Ctx) error {
	mailTemplates, err := h.mailTemplateService.GetAll()
	if err != nil {
		if err.Error() != mongo.ErrNoDocuments.Error() {
			response := utils.BuildResponse(5001, "database error", nil, nil)
			return ctx.Status(http.StatusInternalServerError).JSON(response)
		}
	}
	response := utils.BuildResponse(2001, "success", nil, mailTemplates)
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *mailTemplateHandler) GetOne(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		response := utils.BuildResponse(4002, "Validation error", fiber.Map{"id": "id must be a valid primitive object id"}, nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}
	mailTemplates, err := h.mailTemplateService.GetOne(id)
	if err != nil {
		if err.Error() != mongo.ErrNoDocuments.Error() {
			response := utils.BuildResponse(5001, "database error", nil, nil)
			return ctx.Status(http.StatusInternalServerError).JSON(response)
		} else {
			response := utils.BuildResponse(4041, "mail template not found", nil, nil)
			return ctx.Status(http.StatusNotFound).JSON(response)
		}
	}
	response := utils.BuildResponse(2001, "success", nil, mailTemplates)
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *mailTemplateHandler) CreateMailTemplate(ctx *fiber.Ctx) error {
	// check the user id in the token
	userId, _ := primitive.ObjectIDFromHex(ctx.Locals("userIdClaim").(string))
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
	var templateId *primitive.ObjectID = mailTemplate.Id
	mailTemplate.Id = nil
	if err != nil {
		response := utils.BuildResponse(5001, "database error", nil, nil)
		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	log.Println("Mail template is created:", templateId.Hex(), "by", userId.Hex())
	response := utils.BuildResponse(2001, "success", nil, mailTemplate)
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *mailTemplateHandler) UpdateMailTemplate(ctx *fiber.Ctx) error {
	// check the user id in the token
	userId, _ := primitive.ObjectIDFromHex(ctx.Locals("userIdClaim").(string))
	user, _ := h.userService.FindUserByID(userId)
	if user == nil {
		response := utils.BuildResponse(4041, "user not found", nil, nil)
		return ctx.Status(http.StatusUnauthorized).JSON(response)
	}
	// parse the request body
	var updateMailTemplateRequest dto.UpdateMailTemplateBody
	err := ctx.BodyParser(&updateMailTemplateRequest)
	if err != nil {
		response := utils.BuildResponse(4001, "Invalid request body format", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	mailTemplateId, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		response := utils.BuildResponse(4002, "Validation error", fiber.Map{"id": "id must be a valid primitive object id"}, nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)

	}

	// find mail template by id
	findMailTemplate, _ := h.mailTemplateService.FindMailTemplateByID(mailTemplateId)
	if findMailTemplate == nil {
		response := utils.BuildResponse(4042, "mail template not found", nil, nil)
		return ctx.Status(http.StatusUnauthorized).JSON(response)
	}
	// clean html and string content via SanitizeHTML()
	updateMailTemplateRequest.Content = utils.SanitizeHTML(updateMailTemplateRequest.Content)

	updateMailTemplateRequest.Id = mailTemplateId
	// update
	mailTemplate, err := h.mailTemplateService.UpdateMailTemplate(updateMailTemplateRequest)
	mailTemplate.Id = nil
	if err != nil {
		response := utils.BuildResponse(5001, "database error", nil, nil)
		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	log.Println("Mail template is updated:", mailTemplateId.Hex(), "by", userId.Hex())
	response := utils.BuildResponse(2001, "success", nil, mailTemplate)
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *mailTemplateHandler) DeleteMailTemplate(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		response := utils.BuildResponse(4002, "Validation error", fiber.Map{"id": "id must be a valid primitive object id"}, nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}
	result, err := h.mailTemplateService.DeleteMailTemplate(id)
	if err != nil || !result {
		if err != nil && err.Error() != mongo.ErrNoDocuments.Error() {
			response := utils.BuildResponse(5001, "database error", nil, nil)
			return ctx.Status(http.StatusInternalServerError).JSON(response)
		} else {
			response := utils.BuildResponse(4041, "mail template not found", nil, nil)
			return ctx.Status(http.StatusNotFound).JSON(response)
		}
	}
	response := utils.BuildResponse(2001, "success", nil, nil)
	return ctx.Status(http.StatusOK).JSON(response)
}
