package services

import (
	"pet_api/src/common"
	"pet_api/src/dto/response"
	"pet_api/src/models"
	"pet_api/src/repositories"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	validate = validator.New()
)

func GetAllUsers(c *fiber.Ctx) error {
	offset := c.Query("offset", "0")
	limit := c.Query("limit", "10")

	errors := []string{}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		errors = append(errors, err.Error())
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		errors = append(errors, err.Error())
	}

	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(errors))
	}

	if offsetInt < 0 {
		offsetInt = 0
	}
	if limitInt%5 != 0 {
		limitInt = 10
	}

	totalItems := repositories.CountUsers()
	users, err := repositories.GetAllUsers(offsetInt, limitInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	pagination := common.GeneratePagination(totalItems, limitInt, int64(offsetInt))

	return c.JSON(response.NewResponsePagination(&users, pagination))
}

func CreateUser(c *fiber.Ctx) error {
	newUser := new(models.User)

	if err := c.BodyParser(newUser); err != nil {
		return c.Status(500).JSON(response.ErrorResponse(err.Error()))
	}
	if err := validate.Struct(newUser); err != nil {
		return c.Status(400).JSON(response.ErrorResponse(err.Error()))
	}
	_, err := repositories.CreateUser(*newUser)
	if err != nil {
		return c.Status(500).JSON(response.ErrorResponse(err.Error()))
	}
	return c.JSON(response.NewResponse(*newUser))
}
