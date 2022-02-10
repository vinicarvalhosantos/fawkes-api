package authRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinicarvalhosantos/fawkes-api/internal/handler/user/auth"
	"github.com/vinicarvalhosantos/fawkes-api/internal/util/jwt"
)

func SetupAuthRoutes(router fiber.Router) {

	authRoute := router.Group("/Login")

	//Login Or Register User
	authRoute.Post("/", jwt.Protected(), auth.LoginOrRegisterUser)
	//Create User Login/Register State
	authRoute.Get("/state", jwt.Protected(), auth.CreateState)

}
