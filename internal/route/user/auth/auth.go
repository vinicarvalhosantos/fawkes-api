package authRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinicarvalhosantos/fawkes-api/internal/handler/user/auth"
	"github.com/vinicarvalhosantos/fawkes-api/internal/util/jwt"
)

func SetupAuthRoutes(router fiber.Router) {

	authRoute := router.Group("/Login")

	authRoute.Post("/", jwt.Protected(), auth.LoginOrRegisterUser)

	authRoute.Get("/state", jwt.Protected(), auth.CreateState)

}
