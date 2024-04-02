package services

import (
	"pet_api/src/common"
	"pet_api/src/dto/request"
	"pet_api/src/dto/response"
	"pet_api/src/mapper"
	"pet_api/src/repositories"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	offset, limit, errors := ValidatePaginationParams(c.Query("offset", "0"), c.Query("limit", "10"))
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(errors))
	}
	
	totalItems := repositories.CountUsers()
	users, err := repositories.GetAllUsers(offset, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.UsersModelsToResponse(*users)
	pagination := common.GeneratePagination(totalItems, limit, int64(offset))

	return c.JSON(response.NewResponsePagination(resp, pagination))
}

func CreateUser(c *fiber.Ctx) error {
	model := request.UserRequest{}
	if _, err := common.ValidateRequest(c.Body(), &model); err != nil {
		return c.Status(400).JSON(response.ErrorsResponse(err))
	}

	user := mapper.UserRequestToModel(model)
	userCreated, err := repositories.CreateUser(user)
	if err != nil {
		return c.Status(500).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.UserModelToResponse(*userCreated)
	return c.JSON(response.NewResponse(resp))
}
