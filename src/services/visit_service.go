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

// GetAllVisits godoc
//
//	@Summary		Lista todas las visitas
//	@Security		ApiKeyAuth
//	@Description	Obtiene una lista paginada de todas las visitas.
//	@Tags			visits
//	@Accept			json
//	@Produce		json
//	@Param			offset	query		int													false	"Offset para paginación"
//	@Param			limit	query		int													false	"Límite de resultados por página"
//	@Success		200		{object}	response.BaseResponsePag[response.VisitResponse]	"Respuesta exitosa"
//	@Router			/visits [get]
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

// CreateVisit godoc
//
//	@Summary		Crea una nueva visita
//	@Security		ApiKeyAuth
//	@Description	Crea una nueva visita en la aplicación.
//	@Tags			visits
//	@Accept			json
//	@Produce		json
//	@Param			visitRequest	body		request.VisitRequest							true	"Solicitud de visita"
//	@Success		200				{object}	response.BaseResponse[response.VisitResponse]	"Respuesta exitosa"
//	@Router			/visits [post]
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
	newVisit.UserID = user.ID
	visitResp, err := repositories.CreateVisit(newVisit)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.VisitModelToResponse(visitResp)

	return c.JSON(response.NewResponse(resp))
}

// GetVisitById godoc
//
//	@Summary		Obtiene una visita por ID
//	@Security		ApiKeyAuth
//	@Description	Obtiene los detalles de una visita según su ID.
//	@Tags			visits
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int												true	"ID de la visita"
//	@Success		200	{object}	response.BaseResponse[response.VisitResponse]	"Respuesta exitosa"
//	@Router			/visits/{id} [get]
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

// UpdateVisit godoc
//
//	@Summary		Actualizar una visita programada
//	@Security		ApiKeyAuth
//	@Description	Actualiza una visita identificada por su ID.
//	@Tags			visits
//	@Accept			json
//	@Produce		json
//	@Param			id					path		int							true	"Visit id"
//	@Param			updateVisitRequest	body		request.UpdateVisitRequest true	"Visit update request payload"
//	@Success		200					{object}	response.BaseResponse[response.VisitResponse]
//	@Router			/visits/{id} [patch]
func UpdateVisit(c *fiber.Ctx) error {
	strid := c.Params("id")
	id, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}
	model := request.UpdateVisitRequest{}
	if _, err := helpers.ValidateRequest(c.Body(), &model); err != nil {
		for _, v := range err {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(err))
	}
	visit, err := repositories.GetVisitById(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	userEmail := c.Locals("user_email").(string)
	user, _ := repositories.GetUserByEmail(userEmail)
	if visit.Pet.UserID != user.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse("You can't update this visit"))
	}
	updateVisit := mapper.UpdateVisitRequestToModel(model, visit)
	if _, err := repositories.UpdateVisit(updateVisit); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.VisitModelToResponse(updateVisit)
	return c.JSON(response.MessageResponse("pet visit successfully", resp))
}

// DeleteVisit godoc
//
//	@Summary		Elimina una visita programada
//	@Security		ApiKeyAuth
//	@Description	Elimina una visita identificada por su ID.
//	@Tags			visits
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Visit id"
//	@Success		200	{object}	response.BaseResponse[response.VisitResponse]
//	@Router			/visits/{id} [delete]
func DeleteVisit(c *fiber.Ctx) error {
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

	visit, err := repositories.GetVisitById(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}

	if visit.UserID != user.ID && user.Role.Name != "Admin" {
		log.Println("" + fmt.Sprintf("User with id '%v' can't delete this visit", user.ID))
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse("You can't delete this visit"))
	}

	deletedVisit, err := repositories.DeleteVisit(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.VisitModelToResponse(deletedVisit)
	return c.JSON(response.MessageResponse("visit eliminated successfuly", resp))
}
