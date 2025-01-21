package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"

	"github.com/burakturnaa/mailoop.git/app/dto"
	"github.com/burakturnaa/mailoop.git/app/services"
	"github.com/burakturnaa/mailoop.git/configs"
	"github.com/burakturnaa/mailoop.git/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MailSenderHandler interface {
	GetAll(ctx *fiber.Ctx) error
	GetOne(ctx *fiber.Ctx) error
	CreateLog(ctx *fiber.Ctx) error
}

type mailSenderHandler struct {
	logService          services.LogService
	mailTemplateService services.MailTemplateService
	companyService      services.CompanyService
	userService         services.UserService
	jwtService          services.JWTService
}

func NewMailSenderHandler(
	logService services.LogService,
	mailTemplateService services.MailTemplateService,
	companyService services.CompanyService,
	userService services.UserService,
	jwtService services.JWTService,
) MailSenderHandler {
	return &mailSenderHandler{
		logService:          logService,
		mailTemplateService: mailTemplateService,
		companyService:      companyService,
		userService:         userService,
		jwtService:          jwtService,
	}
}

func (h *mailSenderHandler) GetAll(ctx *fiber.Ctx) error {
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

func (h *mailSenderHandler) GetOne(ctx *fiber.Ctx) error {
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

func (h *mailSenderHandler) CreateLog(ctx *fiber.Ctx) error {
	// check the user id in the token
	userId, _ := primitive.ObjectIDFromHex(ctx.Locals("userIdClaim").(string))
	user, _ := h.userService.FindUserByID(userId)
	if user == nil {
		response := utils.BuildResponse(4041, "user not found", nil, nil)
		return ctx.Status(http.StatusUnauthorized).JSON(response)
	}

	var mailSenderRequest dto.MailSenderBody
	err := ctx.BodyParser(&mailSenderRequest)
	if err != nil {
		response := utils.BuildResponse(4001, "Invalid request body format", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	mailTemplateResult, _ := h.mailTemplateService.FindMailTemplateByID(mailSenderRequest.MailTemplateId)
	if mailTemplateResult == nil { // mail template doesn't exsist
		response := utils.BuildResponse(4042, "mail template not found", nil, nil)
		return ctx.Status(http.StatusUnauthorized).JSON(response)
	}
	emailList := []string{}
	for _, email := range mailSenderRequest.EmailList {
		companyUserResult, _ := h.companyService.FindCompanyByEmail(email)
		if companyUserResult != nil {
			emailList = append(emailList, email)
		}
	}

	// send email
	emails, errSendEmail := sendEmail(emailList, mailTemplateResult.Subject, mailTemplateResult.Content)
	if errSendEmail != nil {
		log.Println("EMAİL SENDER:", errSendEmail)
	}

	if len(emails) < 1 {
		response := utils.BuildResponse(5001, "database error", nil, nil)
		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	LogBody := dto.MailSenderBody{
		MailTemplateId: *mailTemplateResult.Id,
		EmailList:      emails,
	}
	logResult, err := h.logService.CreateLog(LogBody)
	_ = logResult
	if err == nil {
		log.Println("logs saved..")
	}

	response := utils.BuildResponse(2001, "success", nil, nil)
	return ctx.Status(http.StatusOK).JSON(response)
}

func sendEmail(to []string, subject, body string) ([]string, error) {
	log.Printf("emails are being sent to : %s", to)
	envSmtp := configs.EnvSMTP()
	smtpHost := envSmtp["host"]
	smtpPort := envSmtp["port"]
	username := envSmtp["username"]
	password := envSmtp["password"]

	from := envSmtp["email"]

	var emailList []string

	auth := smtp.PlainAuth("", username, password, smtpHost)
	for _, recipient := range to {
		message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n%s", from, recipient, subject, body)
		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{recipient}, []byte(message))
		if err != nil {
			return nil, fmt.Errorf("failed to send email to %s: %w", recipient, err)
		}
		emailList = append(emailList, recipient)
		log.Printf("email sent to %s from %s", recipient, from)
	}
	return emailList, nil
}
