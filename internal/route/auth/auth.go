package authRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinicarvalhosantos/fawkes-api/internal/handler/user/auth"
	"github.com/vinicarvalhosantos/fawkes-api/internal/util/jwt"
)

func SetupAuthRoutes(router fiber.Router) {

	authRoute := router.Group("/login")

	//Login Or Register User
	authRoute.Get("/", jwt.Protected(), auth.LoginOrRegisterTwitchUser)
	//Twitch Callback
	authRoute.Get("/user-callback", auth.UserLoginCallback)

}
