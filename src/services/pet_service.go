package services

import (
	"log"
	"pet_api/src/common"
	"pet_api/src/dto/request"
	"pet_api/src/dto/response"
	"pet_api/src/mapper"
	"pet_api/src/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllPets(c *fiber.Ctx) error {
	offset, limit, errors := ValidatePaginationParams(c.Query("offset", "0"), c.Query("limit", "10"))
	if len(errors) > 0 {
		for _, v := range errors {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(errors))
	}

	totalItems := repositories.CountPets()
	pets, err := repositories.GetAllPets(offset, limit)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}

	users_id := []uint{}
	for _, p := range pets {
		users_id = append(users_id, p.UserID)
	}

	users, err := repositories.GetUsersById(users_id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}

	resp := mapper.PetsModelsToResponse(pets, users)

	pagination := common.GeneratePagination(totalItems, limit, int64(offset))

	return c.JSON(response.NewResponsePagination(resp, pagination))
}

func GetPetById(c *fiber.Ctx) error {
	strid := c.Params("id")
	id, err := strconv.Atoi(strid)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	pet, err := repositories.GetPetById(uint(id))
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	user, _ := repositories.GetUserById(pet.UserID)
	resp := mapper.PetModelToResponse(pet, user)
	return c.JSON(response.NewResponse(resp))
}

func CreatePet(c *fiber.Ctx) error {
	model := request.PetRequest{}
	if _, err := common.ValidateRequest(c.Body(), &model); err != nil {
		for _, v := range err {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(err))
	}
	pet := mapper.PetRequestToModel(model)
	user, err := repositories.GetUserById(model.UserID)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}

	petCreated, err := repositories.CreatePet(pet)

	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}

	resp := mapper.PetModelToResponse(petCreated, user)
	return c.JSON(response.NewResponse(resp))
}
