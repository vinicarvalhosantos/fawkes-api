package userRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinicarvalhosantos/fawkes-api/internal/handler/user"
	authRoutes "github.com/vinicarvalhosantos/fawkes-api/internal/route/user/auth"
	constantUtils "github.com/vinicarvalhosantos/fawkes-api/internal/util/constant"
	"github.com/vinicarvalhosantos/fawkes-api/internal/util/jwt"
)

func SetupUserRoutes(router fiber.Router) {

	userRoute := router.Group("/user")

	//Get All Users
	userRoute.Get("/", jwt.Protected(), user.GetAllUsers)
	//Get User By ID
	userRoute.Get(constantUtils.PathUserIdParam, jwt.Protected(), user.GetUserById)

	authRoutes.SetupAuthRoutes(userRoute)

}
