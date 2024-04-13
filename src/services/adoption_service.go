package services

import (
	"log"
	"pet_api/src/common"
	"pet_api/src/dto/request"
	"pet_api/src/dto/response"
	"pet_api/src/helpers"
	"pet_api/src/mapper"
	"pet_api/src/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllAdoptions(c *fiber.Ctx) error {
	offset, limit, errors := helpers.ValidatePaginationParams(c.Query("offset", "0"), c.Query("limit", "10"))
	if len(errors) > 0 {
		for _, v := range errors {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(errors))
	}
	totalItems := repositories.CountAdoptions()
	adoptions, err := repositories.GetAllAdoptions(offset, limit)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.AdoptionModelsToResponse(adoptions)
	pagination := common.GeneratePagination(totalItems, limit, int64(offset))
	return c.JSON(response.NewResponsePagination(resp, pagination))
}

func CreateAdoption(c *fiber.Ctx) error {
	model := request.AdoptionRequest{}
	if _, err := helpers.ValidateRequest(c.Body(), &model); err != nil {
		for _, v := range err {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(err))
	}
	pet, err := repositories.GetPetById(model.PetID)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	user, err := repositories.GetUserById(model.UserID)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}

	if user.ID == pet.UserID {
		log.Println("you can't adopt your own pet")
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse("you can't adopt your own pet"))
	}

	newAdoption := mapper.AdoptionRequestToModel(model)
	newAdoption.User = user
	newAdoption.Pet = pet

	if helpers.IsFutureDate(newAdoption.AdoptionDate) {
		log.Println("the adoptation must be a future date")
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse("the adoptation must be a future date"))
	}

	if _, err := repositories.CreateAdoption(newAdoption); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}

	resp := mapper.AdoptionModelToResponse(newAdoption)

	return c.JSON(response.NewResponse(resp))
}

func GetAdoptionById(c *fiber.Ctx) error {
	strid := c.Params("id")
	id, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	adoption, err := repositories.GetAdoptionById(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.AdoptionModelToResponse(adoption)
	return c.JSON(response.NewResponse(resp))
}
