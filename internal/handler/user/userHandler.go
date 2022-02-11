package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vinicarvalhosantos/fawkes-api/database"
	"github.com/vinicarvalhosantos/fawkes-api/internal/model"
	constantUtils "github.com/vinicarvalhosantos/fawkes-api/internal/util/constant"
	"github.com/vinicarvalhosantos/fawkes-api/internal/util/cpfCNPJ"
	stringUtil "github.com/vinicarvalhosantos/fawkes-api/internal/util/string"
)

func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []*model.User

	err := db.Find(&users).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constantUtils.StatusInternalServerError, "message": constantUtils.GenericInternalServerErrorMessage, "data": err.Error()})
	}

	if len(users) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": constantUtils.StatusNotFound, "message": model.MessageUser(constantUtils.GenericNotFoundMessage), "data": nil})
	}

	for i := 0; i < len(users); i++ {

		err = FetchUserAddresses(&users[i])

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constantUtils.StatusInternalServerError, "message": constantUtils.GenericInternalServerErrorMessage, "data": err.Error()})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": constantUtils.StatusSuccess, "message": model.MessageUser(constantUtils.GenericFoundSuccessMessage), "data": users})
}

func GetUserById(c *fiber.Ctx) error {
	db := database.DB
	var user *model.User

	id := c.Params("userId")

	err := db.Find(&user, constantUtils.IdCondition, id).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constantUtils.StatusInternalServerError, "message": constantUtils.GenericInternalServerErrorMessage, "data": err.Error()})
	}

	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": constantUtils.StatusNotFound, "message": model.MessageUser(constantUtils.GenericNotFoundMessage), "data": nil})
	}

	err = FetchUserAddresses(&user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constantUtils.StatusInternalServerError, "message": constantUtils.GenericInternalServerErrorMessage, "data": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": constantUtils.StatusSuccess, "message": model.MessageUser(constantUtils.GenericFoundSuccessMessage), "data": user})
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DB
	var user *model.User
	var updateUser *model.UpdateUser

	err := c.BodyParser(&updateUser)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constantUtils.StatusInternalServerError, "message": constantUtils.GenericInternalServerErrorMessage, "data": err.Error()})
	}

	id := c.Params("userId")

	err = db.Find(&user, constantUtils.IdCondition, id).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constantUtils.StatusInternalServerError, "message": constantUtils.GenericInternalServerErrorMessage, "data": err.Error()})
	}

	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": constantUtils.StatusNotFound, "message": model.MessageUser(constantUtils.GenericNotFoundMessage), "data": nil})
	}

	userCpfCnpj := stringUtil.RemoveSpecialCharacters(updateUser.CpfCnpj)

	if !cpfCNPJ.ValidateCpfCNPJ(userCpfCnpj) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": constantUtils.StatusBadRequest, "message": constantUtils.CpfCnpjInvalidMessage, " data": nil})
	}

	err = FetchUserAddresses(&user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": constantUtils.StatusInternalServerError, "message": constantUtils.GenericInternalServerErrorMessage, "data": err.Error()})
	}

	model.PrepareUserToUpdate(&user, updateUser)

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": constantUtils.StatusSuccess, "message": model.MessageUser(constantUtils.GenericUpdateSuccessMessage), "data": user})
}

func FetchUserAddresses(user **model.User) error {
	db := database.DB
	var userAddresses []model.Address

	err := db.Find(&userAddresses, constantUtils.UserIdCondition, (*user).ID).Error

	if err != nil {
		return err
	}

	(*user).Address = userAddresses

	return nil
}
