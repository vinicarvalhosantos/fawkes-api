package reward

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinicarvalhosantos/fawkes-api/internal/handler/reward"
	constants "github.com/vinicarvalhosantos/fawkes-api/internal/util/constant"
	"github.com/vinicarvalhosantos/fawkes-api/internal/util/jwt"
)

func SetupRewardRoutes(router fiber.Router) {

	rewardRoute := router.Group("/reward")
	//Get all rewards
	rewardRoute.Get("/", jwt.Protected(), reward.GetReward)
	//Get reward by id
	rewardRoute.Get(constants.PathRewardIdParam, jwt.Protected(), reward.GetRewardByID)
	//Create a new Reward
	rewardRoute.Post("/", jwt.Protected(), reward.RegisterReward)
	//Update reward by id
	rewardRoute.Put(constants.PathRewardIdParam, jwt.Protected(), reward.UpdateReward)
	//Delete reward
	rewardRoute.Delete(constants.PathRewardIdParam, jwt.Protected(), reward.DeleteReward)
}
