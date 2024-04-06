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
	breed := c.Query("breed", "")
	color := c.Query("color", "")
	totalItems := repositories.CountPets(breed, color)
	pets, err := repositories.GetAllPets(offset, limit, breed, color)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}

	resp := mapper.PetsModelsToResponse(pets)

	pagination := common.GeneratePagination(totalItems, limit, int64(offset))

	return c.JSON(response.NewResponsePagination(resp, pagination))
}

func GetPetById(c *fiber.Ctx) error {
	strid := c.Params("id")
	id, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	pet, err := repositories.GetPetById(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}

	resp := mapper.PetModelToResponse(pet)
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

	_, err := repositories.GetUserById(model.UserID)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}

	petCreated, err := repositories.CreatePet(pet)
	// petCreated.User = user
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}

	resp := mapper.PetModelToResponse(petCreated)
	return c.JSON(response.NewResponse(resp))
}

func UpdatePetImages(c *fiber.Ctx) error {
	strid := c.Params("id")
	id, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	pet, err := repositories.GetPetById(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}

	form, err := c.MultipartForm()
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}

	if _, err := CreatePetImages(pet.ID, form); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))	
	}

	return c.JSON(response.MessageResposne("images created successfully"))
}
