package main

import (
	"github.com/burakturnaa/mailoop.git/app/dto"
	"github.com/burakturnaa/mailoop.git/app/handlers"
	"github.com/burakturnaa/mailoop.git/app/middlewares"
	"github.com/burakturnaa/mailoop.git/app/repository"
	"github.com/burakturnaa/mailoop.git/app/services"
	"github.com/burakturnaa/mailoop.git/configs"
	"github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New()

	dbClientUsers := configs.GetCollection(configs.DB, "users")
	dbClientMailTemplates := configs.GetCollection(configs.DB, "mail_templates")

	UserRepository := repository.NewUserRepository(dbClientUsers)
	MailTemplateRepository := repository.NewMailTemplateRepository(dbClientMailTemplates)

	//auth
	var authService services.AuthService = services.NewAuthService(UserRepository)
	var jwtService services.JWTService = services.NewJWTService()
	var userService services.UserService = services.NewUserService(UserRepository)
	authHandler := handlers.NewAuthHandler(authService, jwtService, userService)

	authRoutes := server.Group("/api/auth")
	authRoutes.Post("/login", middlewares.AuthValidation(&dto.LoginBody{}), authHandler.Login)
	authRoutes.Post("/register", middlewares.AuthValidation(&dto.RegisterBody{}), authHandler.Register)

	//mail template
	var mailTemplateService services.MailTemplateService = services.NewMailTemplateService(MailTemplateRepository)
	mailTemplateHandler := handlers.NewMailTemplateHandler(mailTemplateService, userService, jwtService)

	mailTemplateRoutes := server.Group("/api/mailtemp")
	mailTemplateRoutes.Post("/", middlewares.AuthorizeJWT(jwtService), middlewares.MailTemplateValidation(&dto.MailTemplateBody{}), mailTemplateHandler.CreateMailTemplate)
	mailTemplateRoutes.Put("/:id", middlewares.AuthorizeJWT(jwtService), middlewares.MailTemplateValidation(&dto.UpdateMailTemplateBody{}), mailTemplateHandler.UpdateMailTemplate)
	// mailTemplateRoutes.Get("/", middlewares.AuthMiddleware(&dto.RegisterBody{}), authHandler.Register)
	// mailTemplateRoutes.Get("/:id", middlewares.AuthMiddleware(&dto.RegisterBody{}), authHandler.Register)

	server.Listen(":3000")
}
