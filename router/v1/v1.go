package v1

import (
	"github.com/gofiber/fiber/v2"
	addressRoutes "github.com/vinicarvalhosantos/fawkes-api/internal/route/address"
	authRoutes "github.com/vinicarvalhosantos/fawkes-api/internal/route/auth"
	rewardRoutes "github.com/vinicarvalhosantos/fawkes-api/internal/route/reward"
	userRoutes "github.com/vinicarvalhosantos/fawkes-api/internal/route/user"
)

func SetupV1Routes(router fiber.Router) {

	api := router.Group("/v1")

	//User Routes
	userRoutes.SetupUserRoutes(api)

	//Auth Routes
	authRoutes.SetupAuthRoutes(api)

	//Address Routes
	addressRoutes.SetupAddressRoutes(api)

	//Reward Routes
	rewardRoutes.SetupRewardRoutes(api)

}
