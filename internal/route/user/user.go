package userRoutes

import (
	"github.com/gofiber/fiber/v2"
	authRoutes "gitlab.com/vinicius.csantos/fawkes-api/internal/route/user/auth"
)

func SetupUserRoutes(router fiber.Router) {

	userRoute := router.Group("/user")

	authRoutes.SetupAuthRoutes(userRoute)

}
