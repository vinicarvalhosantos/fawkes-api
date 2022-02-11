package reward

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinicarvalhosantos/fawkes-api/internal/handler/reward"
	"github.com/vinicarvalhosantos/fawkes-api/internal/util/jwt"
)

func SetupRewardRoutes(router fiber.Router) {

	rewardRoute := router.Group("/reward")
	//Get all rewards
	rewardRoute.Get("/", jwt.Protected(), reward.GetReward)
	//Get reward by id
	rewardRoute.Get("/:rewardId", jwt.Protected(), reward.GetRewardByID)
	//Create a new Reward
	rewardRoute.Post("/", jwt.Protected(), reward.RegisterReward)
	//Update reward by id
	rewardRoute.Put("/:rewardId", jwt.Protected(), reward.UpdateReward)
	//Delete reward
	rewardRoute.Delete("/:rewardId", jwt.Protected(), reward.DeleteReward)
}
