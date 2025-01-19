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

	UserRepository := repository.NewUserRepository(dbClientUsers)
	_ = UserRepository

	//auth
	var authService services.AuthService = services.NewAuthService(UserRepository)
	var jwtService services.JWTService = services.NewJWTService()
	var userService services.UserService = services.NewUserService(UserRepository)
	authHandler := handlers.NewAuthHandler(authService, jwtService, userService)

	authRoutes := server.Group("/api/auth")
	authRoutes.Post("/login", middlewares.AuthMiddleware(&dto.LoginBody{}), authHandler.Login)
	authRoutes.Post("/register", middlewares.AuthMiddleware(&dto.RegisterBody{}), authHandler.Register)

	server.Listen(":3000")
}
