package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/burakturnaa/mailoop.git/app/dto"
	"github.com/burakturnaa/mailoop.git/app/services"
	"github.com/burakturnaa/mailoop.git/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CompanyHandler interface {
	GetAll(ctx *fiber.Ctx) error
	GetOne(ctx *fiber.Ctx) error
	CreateCompany(ctx *fiber.Ctx) error
	UpdateCompany(ctx *fiber.Ctx) error
	DeleteCompany(ctx *fiber.Ctx) error
}

type companyHandler struct {
	companyService services.CompanyService
	userService    services.UserService
	jwtService     services.JWTService
}

func NewCompanyHandler(
	companyService services.CompanyService,
	userService services.UserService,
	jwtService services.JWTService,
) CompanyHandler {
	return &companyHandler{
		companyService: companyService,
		userService:    userService,
		jwtService:     jwtService,
	}
}

func (h *companyHandler) GetAll(ctx *fiber.Ctx) error {
	companies, err := h.companyService.GetAll()
	if err != nil {
		if err.Error() != mongo.ErrNoDocuments.Error() {
			response := utils.BuildResponse(5001, "database error", nil, nil)
			return ctx.Status(http.StatusInternalServerError).JSON(response)
		}
	}
	response := utils.BuildResponse(2001, "success", nil, companies)
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *companyHandler) GetOne(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		response := utils.BuildResponse(4002, "Validation error", fiber.Map{"id": "id must be a valid primitive object id"}, nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}
	companies, err := h.companyService.GetOne(id)
	if err != nil {
		if err.Error() != mongo.ErrNoDocuments.Error() {
			response := utils.BuildResponse(5001, "database error", nil, nil)
			return ctx.Status(http.StatusInternalServerError).JSON(response)
		} else {
			response := utils.BuildResponse(4041, "company not found", nil, nil)
			return ctx.Status(http.StatusNotFound).JSON(response)
		}
	}
	response := utils.BuildResponse(2001, "success", nil, companies)
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *companyHandler) CreateCompany(ctx *fiber.Ctx) error {
	// check the user id in the token
	userId, _ := primitive.ObjectIDFromHex(ctx.Locals("userIdClaim").(string))

	user, _ := h.userService.FindUserByID(userId)
	if user == nil {
		response := utils.BuildResponse(4041, "user not found", nil, nil)
		return ctx.Status(http.StatusUnauthorized).JSON(response)
	}

	var companyRequest dto.CompanyBody
	err := ctx.BodyParser(&companyRequest)
	if err != nil {
		response := utils.BuildResponse(4001, "Invalid request body format", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	findCompany, _ := h.companyService.FindCompanyByEmail(companyRequest.Email)
	if findCompany != nil {
		response := utils.BuildResponse(4091, "company already exists", nil, nil)
		return ctx.Status(http.StatusConflict).JSON(response)
	}

	company, err := h.companyService.CreateCompany(companyRequest)
	var templateId *primitive.ObjectID = company.Id
	company.Id = nil
	if err != nil {
		response := utils.BuildResponse(5001, "database error", nil, nil)
		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	log.Println("Company is created:", templateId.Hex(), "by", userId.Hex())
	response := utils.BuildResponse(2001, "success", nil, company)
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *companyHandler) UpdateCompany(ctx *fiber.Ctx) error {
	// check the user id in the token
	userId, _ := primitive.ObjectIDFromHex(ctx.Locals("userIdClaim").(string))
	user, _ := h.userService.FindUserByID(userId)
	if user == nil {
		response := utils.BuildResponse(4041, "user not found", nil, nil)
		return ctx.Status(http.StatusUnauthorized).JSON(response)
	}
	// parse the request body
	var updateCompanyRequest dto.UpdateCompanyBody
	err := ctx.BodyParser(&updateCompanyRequest)
	if err != nil {
		response := utils.BuildResponse(4001, "Invalid request body format", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	companyId, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		response := utils.BuildResponse(4002, "Validation error", fiber.Map{"id": "id must be a valid primitive object id"}, nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)

	}

	// find company by id
	findCompany, _ := h.companyService.FindCompanyByID(companyId)
	if findCompany == nil {
		response := utils.BuildResponse(4042, "company not found", nil, nil)
		return ctx.Status(http.StatusUnauthorized).JSON(response)
	}

	updateCompanyRequest.Id = companyId
	// update
	company, err := h.companyService.UpdateCompany(updateCompanyRequest)
	company.Id = nil
	if err != nil {
		fmt.Println(err)
		response := utils.BuildResponse(5001, "database error", nil, nil)
		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	log.Println("Company is updated:", companyId.Hex(), "by", userId.Hex())
	response := utils.BuildResponse(2001, "success", nil, company)
	return ctx.Status(http.StatusOK).JSON(response)
}

func (h *companyHandler) DeleteCompany(ctx *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(ctx.Params("id"))
	if err != nil {
		response := utils.BuildResponse(4002, "Validation error", fiber.Map{"id": "id must be a valid primitive object id"}, nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}
	result, err := h.companyService.DeleteCompany(id)
	if err != nil || !result {
		if err != nil && err.Error() != mongo.ErrNoDocuments.Error() {
			response := utils.BuildResponse(5001, "database error", nil, nil)
			return ctx.Status(http.StatusInternalServerError).JSON(response)
		} else {
			response := utils.BuildResponse(4041, "company not found", nil, nil)
			return ctx.Status(http.StatusNotFound).JSON(response)
		}
	}
	response := utils.BuildResponse(2001, "success", nil, nil)
	return ctx.Status(http.StatusOK).JSON(response)
}
