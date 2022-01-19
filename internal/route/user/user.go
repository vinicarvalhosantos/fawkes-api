package userRoutes

import (
	"github.com/gofiber/fiber/v2"
	userHandler "gitlab.com/vinicius.csantos/go-template-api/internal/handler/user"
	authRoutes "gitlab.com/vinicius.csantos/go-template-api/internal/route/user/auth"
	constants "gitlab.com/vinicius.csantos/go-template-api/internal/util/constant"
	"gitlab.com/vinicius.csantos/go-template-api/internal/util/jwt"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")

	//Get All Users
	user.Get("/", jwt.Protected(), userHandler.GetUsers)
	//Get User By ID
	user.Get(constants.PathUserIdParam, jwt.Protected(), userHandler.GetUser)
	//Update User
	user.Put(constants.PathUserIdParam, jwt.Protected(), userHandler.UpdateUser)
	//Delete User
	user.Delete(constants.PathUserIdParam, jwt.Protected(), userHandler.DeleteUser)

	//Register Route
	authRoutes.SetupRegisterRoute(user)
	//Authenticate Route
	authRoutes.SetupLoginRoute(user)

}
