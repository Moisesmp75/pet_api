package services

import (
	"fmt"
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

// GetAllAdoptions godoc
//
//	@Summary		Lista todas las adopciones
//	@Description	Obtiene una lista paginada de todas las adopciones.
//	@Tags			adoptions
//	@Accept			json
//	@Produce		json
//	@Param			offset	query		int													false	"Offset para paginación"
//	@Param			limit	query		int													false	"Límite de resultados por página"
//	@Success		200		{object}	response.BaseResponsePag[response.AdoptionResponse]	"Respuesta exitosa"
//	@Router			/adoptions [get]
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

// CreateAdoption godoc
//
//	@Summary		Crea una nueva adopción
//	@Description	Crea una nueva adopción en la aplicación.
//	@Tags			adoptions
//	@Accept			json
//	@Produce		json
//	@Param			adoptionRequest	body		request.AdoptionRequest								true	"Solicitud de adopción"
//	@Success		200				{object}	response.BaseResponse[response.AdoptionResponse]	"Respuesta exitosa"
//	@Router			/adoptions [post]
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
	userEmail := c.Locals("user_email").(string)
	user, err := repositories.GetUserByEmail(userEmail)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse(err.Error()))
	}

	// user, err := repositories.GetUserById(model.UserID)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	// }

	if user.ID == pet.UserID {
		log.Println("you can't adopt your own pet")
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse("you can't adopt your own pet"))
	}

	newAdoption := mapper.AdoptionRequestToModel(model)
	newAdoption.User = user
	newAdoption.Pet = pet
	newAdoption.UserID = user.ID

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

// GetAdoptionById godoc
//
//	@Summary		Obtiene una adopción por ID
//	@Description	Obtiene los detalles de una adopción según su ID.
//	@Tags			adoptions
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int													true	"ID de la adopción"
//	@Success		200	{object}	response.BaseResponse[response.AdoptionResponse]	"Respuesta exitosa"
//	@Router			/adoptions/{id} [get]
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

// DeleteAdoption godoc
//
//	@Summary		Elimina una adopcion programada
//	@Description	Elimina una adopcion identificada por su ID.
//	@Tags			adoptions
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Adoption id"
//	@Success		200	{object}	response.BaseResponse[response.AdoptionResponse]
//	@Router			/adoptions/{id} [delete]
func DeleteAdoption(c *fiber.Ctx) error {
	strid := c.Params("id")
	id, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}
	userEmail := c.Locals("user_email").(string)
	user, err := repositories.GetUserByEmail(userEmail)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse(err.Error()))
	}
	adoption, err := repositories.GetAdoptionById(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}
	if adoption.UserID != user.ID && user.Role.Name != "Admin" {
		log.Println("" + fmt.Sprintf("User with id '%v' can't delete this adoption", user.ID))
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse("You can't delete this adoption"))
	}
	deletedAdoption, err := repositories.DeleteAdoption(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.AdoptionModelToResponse(deletedAdoption)
	return c.JSON(response.MessageResponse("adoption eliminated successfuly", resp))
}
