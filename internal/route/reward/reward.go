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
	//Disable IsEnabled by id
	rewardRoute.Patch(constants.PathUpdateRewardDisable, jwt.Protected(), reward.DisableReward)
	//Enable IsEnabled by id
	rewardRoute.Patch(constants.PathUpdateRewardEnable, jwt.Protected(), reward.EnableReward)
	//Enable IsPaused by id
	rewardRoute.Patch(constants.PathUpdateRewardPause, jwt.Protected(), reward.PauseReward)
	//Disable IsPaused by id
	rewardRoute.Patch(constants.PathUpdateRewardUnpause, jwt.Protected(), reward.UnpauseReward)
	//Delete reward
	rewardRoute.Delete(constants.PathRewardIdParam, jwt.Protected(), reward.DeleteReward)
}
