package model

import (
	stringUtil "github.com/vinicarvalhosantos/fawkes-api/internal/util/string"
	"time"
)

type Reward struct {
	ID              uint
	Title           string
	Prompt          string
	Cost            int16
	BackgroundColor string
	IsEnabled       bool
	ShouldSkipQueue bool
	MaxUsePerStream uint
	MaxUserPerUser  uint
	GlobalCooldown  int32
	IsPaused        bool
	CreatedBy       int64
	CreatedAt       time.Time
	UpdateAt        time.Time
}

type UpdateReward struct {
	Title           string
	Prompt          string
	Cost            int16
	BackgroundColor string
	ShouldSkipQueue bool
	MaxUsePerStream uint
	MaxUsePerUser   uint
	GlobalCooldowm  int32
}

func CheckIfRewardEntityIsValid(reward *Reward) (bool, string) {
	if reward.Title == "" {
		return false, "Title"
	}
	if reward.Prompt == "" {
		return false, "Prompt"
	}
	if reward.Cost <= 0 {
		return false, "Cost"
	}
	if reward.BackgroundColor == "" {
		return false, "BackgroundColor"
	}
	if reward.MaxUsePerStream <= 0 {
		return false, "MaxUsePerStream"
	}
	if reward.MaxUserPerUser <= 0 {
		return false, "MaxUserPerUser"
	}
	if reward.GlobalCooldown <= 0 {
		return false, "GlobalCooldown"
	}
	if reward.CreatedBy <= 0 {
		return false, "CreatedBy"
	}
	return true, ""
}

/*func PrepareRewardToUpdate(reward **Reward, updateReward *UpdateReward) {

	if updateReward.Title != "" {
		(*reward).Title == updateReward.Title
	}
	if updateReward.Prompt != "" {
		(*reward).Prompt == updateReward.Prompt
	}
	if updateReward.Cost != 0 {
		(*reward).Cost == updateReward.Cost
	}
	if updateReward.BackgroundColor != "" {
		(*reward).BackgroundColor == updateReward.BackgroundColor
	}
	if updateReward.ShouldSkipQueue != 0 {
		(*reward).ShouldSkipQueue == updateReward.ShouldSkipQueue
	}
	if updateReward.MaxUsePerStream != 0 {
		(*reward).MaxUsePerStream == updateReward.MaxUsePerStream
	}
	if updateReward.MaxUsePerUser != 0 {
		(*reward).MaxUsePerUser == updateReward.MaxUsePerUser
	}
	if updateReward.GlobalCooldowm != 0 {
		(*reward).GlobalCooldown == reward.GlobalCooldown
	}
}*/

func MessageReward(genericMessage string) string {
	return stringUtil.FormatGenericMessagesString(genericMessage, "Reward")
}