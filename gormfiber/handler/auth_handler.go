package handler

import (
	"gofiber/database"
	"gofiber/model/entity"
	"gofiber/model/request"
	"gofiber/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	loginRequest := request.LoginRequest{}
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "fail",
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : "failed",
			"error" : errValidate.Error(),
		})
	}

	var user entity.User
	err := database.DB.First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message" : "wrong credential",
		})
	}

	isValid := utils.CheckPasswordHas(loginRequest.Password, user.Password)
	if !isValid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "wrong credential",
		})
	}

	return c.JSON(fiber.Map{
		"token": "secret",
	})
}