package reward

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vinicarvalhosantos/fawkes-api/database"
	"github.com/vinicarvalhosantos/fawkes-api/internal/model"
	constants "github.com/vinicarvalhosantos/fawkes-api/internal/util/constant"
	stringUtil "github.com/vinicarvalhosantos/fawkes-api/internal/util/string"
)

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
