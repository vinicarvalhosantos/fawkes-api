package authRoutes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/vinicius.csantos/fawkes-api/internal/handler/user/auth"
	"gitlab.com/vinicius.csantos/fawkes-api/internal/util/jwt"
)

func SetupAuthRoutes(router fiber.Router) {

	authRoute := router.Group("/Login")

	authRoute.Post("/", jwt.Protected(), auth.LoginOrRegisterUser)

	authRoute.Get("/state", jwt.Protected(), auth.CreateState)

}
