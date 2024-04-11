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

func GetAllVisits(c *fiber.Ctx) error {
	offset, limit, errors := helpers.ValidatePaginationParams(c.Query("offset", "0"), c.Query("limit", "10"))
	if len(errors) > 0 {
		for _, v := range errors {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(errors))
	}
	totalItems := repositories.CountVisits()
	visits, err := repositories.GetAllVisits(offset, limit)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.VisitModelsToResponse(visits)
	pagination := common.GeneratePagination(totalItems, limit, int64(offset))

	return c.JSON(response.NewResponsePagination(resp, pagination))
}

func CreateVisit(c *fiber.Ctx) error {
	model := request.VisitRequest{}
	if _, err := helpers.ValidateRequest(c.Body(), &model); err != nil {
		for _, v := range err {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(err))
	}
	pet, err := repositories.GetPetById(model.PetID)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	user, err := repositories.GetUserById(model.UserID)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	if user.ID == pet.UserID {
		log.Println("you can't visit your own pet")
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse("you can't visit your own pet"))
	}
	newVisit := mapper.VisitRequestToModel(model)
	if helpers.IsFutureDate(newVisit.Date) {
		log.Println("the date must be a future date")
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse("the date must be a future date"))
	}
	newVisit.Pet = pet
	newVisit.User = user
	if _, err := repositories.CreateVisit(newVisit); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.VisitModelToResponse(newVisit)

	return c.JSON(response.NewResponse(resp))
}

func GetVisitById(c *fiber.Ctx) error {
	strid := c.Params("id")
	id, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	visit, err := repositories.GetVisitById(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.VisitModelToResponse(visit)
	return c.JSON(response.NewResponse(resp))
}