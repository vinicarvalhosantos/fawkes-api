package reward

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinicarvalhosantos/fawkes-api/internal/handler/reward"
	"github.com/vinicarvalhosantos/fawkes-api/internal/util/jwt"
)

func SetupRewardRoutes(router fiber.Router) {

	rewardRoute := router.Group("/reward")

	//Create a new Reward
	rewardRoute.Post("/", jwt.Protected(), reward.RegisterReward)

}
