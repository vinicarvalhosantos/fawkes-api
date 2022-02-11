package reward

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vinicarvalhosantos/fawkes-api/database"
	"github.com/vinicarvalhosantos/fawkes-api/internal/model"
	constants "github.com/vinicarvalhosantos/fawkes-api/internal/util/constant"
	stringUtil "github.com/vinicarvalhosantos/fawkes-api/internal/util/string"
)

func GetReward(c *fiber.Ctx) error {
	db := database.DB
	var reward []*model.Reward

	err := db.Find(&reward).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  constants.StatusInternalServerError,
			"message": constants.GenericInternalServerErrorMessage,
			"data":    err.Error()})
	}

	if len(reward) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  constants.StatusNotFound,
			"message": model.MessageReward(constants.GenericNotFoundMessage),
			"data":    nil})
	}
	return c.JSON(fiber.Map{
		"status":  constants.StatusSuccess,
		"message": model.MessageReward(constants.GenericFoundSuccessMessage),
		"data":    reward})
}

func GetRewardByID(c *fiber.Ctx) error {
	db := database.DB
	var reward *model.Reward

	id := c.Params("rewardId")

	err := db.Find(&reward, constants.IdCondition, id).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  constants.StatusInternalServerError,
			"message": constants.GenericInternalServerErrorMessage,
			"data":    err.Error()})
	}

	if reward.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  constants.StatusNotFound,
			"message": model.MessageReward(constants.GenericNotFoundMessage),
			"data":    nil})
	}

	return c.JSON(fiber.Map{
		"status": constants.StatusSuccess,
		"model":  model.MessageReward(constants.GenericFoundSuccessMessage),
		"data":   reward})
}

func RegisterReward(c *fiber.Ctx) error {
	db := database.DB
	var reward *model.Reward
	var searchReward *model.Reward

	err := c.BodyParser(&reward)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constants.StatusInternalServerError, "message": constants.GenericInternalServerErrorMessage, "data": err.Error()})
	}

	isValid, invalidField := model.CheckIfRewardEntityIsValid(reward)

	if !isValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": constants.StatusBadRequest, "message": stringUtil.FormatGenericMessagesString(constants.GenericInvalidFieldMessage, invalidField), "data": nil})
	}

	err = db.Find(&searchReward, constants.IdCondition, reward.ID).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constants.StatusInternalServerError, "message": constants.GenericInternalServerErrorMessage, "data": err.Error()})
	}

	if searchReward.ID != uuid.Nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": constants.StatusBadRequest, "message": "This ID is already exists in our database", "data": nil})
	}

	err = db.Create(&reward).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constants.StatusInternalServerError, "message": constants.GenericInternalServerErrorMessage, "data": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": constants.StatusSuccess, "message": model.MessageReward(constants.GenericCreateSuccessMessage), "data": reward})
}

func UpdateReward(c *fiber.Ctx) error {
	db := database.DB
	var updateReward *model.UpdateReward
	var reward *model.Reward

	err := c.BodyParser(&updateReward)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  constants.StatusInternalServerError,
			"message": model.MessageReward(constants.GenericInternalServerErrorMessage),
			"data":    err.Error()})
	}

	id := c.Params("rewardId")

	err = db.Find(&reward, constants.IdCondition, id).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  constants.StatusInternalServerError,
			"message": model.MessageReward(constants.GenericInternalServerErrorMessage),
			"data":    err.Error()})
	}

	if reward.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  constants.StatusNotFound,
			"message": model.MessageReward(constants.GenericNotFoundMessage),
			"data":    nil})
	}

	model.PrepareRewardToUpdate(&reward, updateReward)
	err = db.Save(reward).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  constants.StatusInternalServerError,
			"message": model.MessageReward(constants.GenericInternalServerErrorMessage),
			"data":    err.Error()})
	}

	return c.JSON(fiber.Map{
		"status":  constants.StatusSuccess,
		"message": model.MessageAddress(constants.GenericFoundSuccessMessage),
		"data":    reward})
}

func DeleteReward(c *fiber.Ctx) error {
	db := database.DB
	var reward *model.Reward

	id := c.Params("rewardId")

	err := db.Find(&reward, constants.IdCondition, id).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  constants.StatusInternalServerError,
			"message": model.MessageReward(constants.GenericInternalServerErrorMessage),
			"data":    err.Error()})
	}

	if reward.ID == uuid.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  constants.StatusNotFound,
			"message": model.MessageReward(constants.GenericNotFoundMessage),
			"data":    nil})
	}

	err = db.Delete(&reward).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  constants.StatusInternalServerError,
			"message": model.MessageReward(constants.GenericInternalServerErrorMessage),
			"data":    err.Error()})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}
