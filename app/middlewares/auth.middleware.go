package middlewares

import (
	"log"
	"net/http"

	"github.com/burakturnaa/mailoop.git/app/services"
	"github.com/burakturnaa/mailoop.git/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// AuthorizeJWT validates the token, return 401 if not valid
func AuthorizeJWT(jwtService services.JWTService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			response := utils.BuildResponse(4003, "No token provided", nil, nil)
			return ctx.Status(http.StatusBadRequest).JSON(response)
		}

		token := jwtService.ValidateToken(authHeader, ctx)
		if token == nil {
			response := utils.BuildResponse(4011, "Unauthorized", nil, nil)
			return ctx.Status(http.StatusUnauthorized).JSON(response)
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["iss"])
			ctx.Locals("userIdClaims", claims["user_id"])
		} else {
			response := utils.BuildResponse(4011, "Unauthorized", nil, nil)
			return ctx.Status(http.StatusUnauthorized).JSON(response)
		}
		return ctx.Next()
	}

}
