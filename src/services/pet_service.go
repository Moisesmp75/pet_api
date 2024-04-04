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
	//Mapper PetsModel to PetsResponse

	pagination := common.GeneratePagination(totalItems, limit, int64(offset))

	return c.JSON(response.NewResponsePagination(pets, pagination))
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
	//Mapper PetModel to PetResponse
	return c.JSON(response.NewResponse(pet))
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
	user, err := repositories.GetUserById(pet.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}

	userResponse := mapper.UserModelToResponse(*user)
	petCreated, err := repositories.CreatePet(pet)

	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}

	resp := mapper.PetModelToResponse(*petCreated, userResponse)
	return c.JSON(response.NewResponse(resp))
}
