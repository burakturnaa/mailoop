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

	// collections
	dbClientUsers := configs.GetCollection(configs.DB, "users")
	dbClientMailTemplates := configs.GetCollection(configs.DB, "mail_templates")
	dbClientCompanies := configs.GetCollection(configs.DB, "companies")
	dbClientLogs := configs.GetCollection(configs.DB, "logs")

	// repositories
	UserRepository := repository.NewUserRepository(dbClientUsers)
	MailTemplateRepository := repository.NewMailTemplateRepository(dbClientMailTemplates)
	CompanyRepository := repository.NewCompanyRepository(dbClientCompanies)
	LogRepository := repository.NewLogRepository(dbClientLogs)

	// services
	var authService services.AuthService = services.NewAuthService(UserRepository)
	var jwtService services.JWTService = services.NewJWTService()
	var userService services.UserService = services.NewUserService(UserRepository)
	var mailTemplateService services.MailTemplateService = services.NewMailTemplateService(MailTemplateRepository)
	var companyService services.CompanyService = services.NewCompanyService(CompanyRepository)
	var logService services.LogService = services.NewLogService(LogRepository)

	// handlers
	authHandler := handlers.NewAuthHandler(authService, jwtService, userService)
	mailTemplateHandler := handlers.NewMailTemplateHandler(mailTemplateService, userService, jwtService)
	companyHandler := handlers.NewCompanyHandler(companyService, userService, jwtService)
	mailSenderHandler := handlers.NewMailSenderHandler(logService, mailTemplateService, companyService, userService, jwtService)
	logHandler := handlers.NewLogHandler(logService, userService, jwtService)

	// auth routes
	authRoutes := server.Group("/api/auth")
	authRoutes.Post("/login", middlewares.AuthValidation(&dto.LoginBody{}), authHandler.Login)
	authRoutes.Post("/register", middlewares.AuthValidation(&dto.RegisterBody{}), authHandler.Register)

	// mail template routes
	mailTemplateRoutes := server.Group("/api/mailtemp")
	mailTemplateRoutes.Get("/", middlewares.AuthorizeJWT(jwtService), mailTemplateHandler.GetAll)
	mailTemplateRoutes.Get("/:id", middlewares.AuthorizeJWT(jwtService), mailTemplateHandler.GetOne)
	mailTemplateRoutes.Post("/", middlewares.AuthorizeJWT(jwtService), middlewares.MailTemplateValidation(&dto.MailTemplateBody{}), mailTemplateHandler.CreateMailTemplate)
	mailTemplateRoutes.Put("/:id", middlewares.AuthorizeJWT(jwtService), middlewares.MailTemplateValidation(&dto.UpdateMailTemplateBody{}), mailTemplateHandler.UpdateMailTemplate)
	mailTemplateRoutes.Delete("/:id", middlewares.AuthorizeJWT(jwtService), mailTemplateHandler.DeleteMailTemplate)

	// company routes
	companyRoutes := server.Group("/api/company")
	companyRoutes.Get("/", middlewares.AuthorizeJWT(jwtService), companyHandler.GetAll)
	companyRoutes.Get("/:id", middlewares.AuthorizeJWT(jwtService), companyHandler.GetOne)
	companyRoutes.Post("/", middlewares.AuthorizeJWT(jwtService), middlewares.CompanyValidation(&dto.CompanyBody{}), companyHandler.CreateCompany)
	companyRoutes.Put("/:id", middlewares.AuthorizeJWT(jwtService), middlewares.CompanyValidation(&dto.UpdateCompanyBody{}), companyHandler.UpdateCompany)
	companyRoutes.Delete("/:id", middlewares.AuthorizeJWT(jwtService), companyHandler.DeleteCompany)

	// mail sender route
	mailSenderRoutes := server.Group("/api/mail")
	mailSenderRoutes.Post("/send", middlewares.AuthorizeJWT(jwtService), middlewares.MailSenderValidation(&dto.MailSenderBody{}), mailSenderHandler.CreateLog)

	// log route
	LogRoutes := server.Group("/api/log")
	LogRoutes.Get("/", middlewares.AuthorizeJWT(jwtService), logHandler.GetAll)

	server.Listen(":3000")
}
