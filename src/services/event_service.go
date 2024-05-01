package services

import (
	"fmt"
	"log"
	"pet_api/src/common"
	"pet_api/src/dto/response"
	"pet_api/src/helpers"
	"pet_api/src/mapper"
	"pet_api/src/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetAllEvents godoc
//
//	@Summary		Lista todos los eventos
//	@Security		ApiKeyAuth
//	@Description	Obtiene una lista paginada de todos los eventos
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			offset	query		int													false	"Offset para paginación"
//	@Param			limit	query		int													false	"Límite de resultados por página"
//	@Success		200		{object}	response.BaseResponsePag[[]response.EventResponse]	"Respuesta exitosa"
//	@Router			/events [get]
func GetAllEvents(c *fiber.Ctx) error {
	offset, limit, errors := helpers.ValidatePaginationParams(c.Query("offset", "0"), c.Query("limit", "10"))
	if len(errors) > 0 {
		for _, v := range errors {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(errors))
	}
	totalItems := repositories.CountVisits()
	events, err := repositories.GetAllEvents(offset, limit)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.EventModelsToResponse(events)
	pagination := common.GeneratePagination(totalItems, limit, int64(offset))

	return c.JSON(response.NewResponsePagination(resp, pagination))

}

// GetEventById godoc
//
//	@Summary		Obtiene un Evento por ID
//	@Security		ApiKeyAuth
//	@Description	Obtiene los detalles de un Evento según su ID.
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int												true	"ID del Evento"
//	@Success		200	{object}	response.BaseResponse[response.EventResponse]	"Respuesta exitosa"
//	@Router			/events/{id} [get]
func GetEventById(c *fiber.Ctx) error {
	strid := c.Params("id")
	id, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	event, err := repositories.GetEventById(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.EventModelToResponse(event)
	return c.JSON(response.NewResponse(resp))
}

// DeleteVisit godoc
//
//	@Summary		Elimina un Evento programado
//	@Security		ApiKeyAuth
//	@Description	Elimina un Evento identificado por su ID.
//	@Tags			events
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"ID del Evento"
//	@Success		200	{object}	response.BaseResponse[response.EventResponse]
//	@Router			/events/{id} [delete]
func DeleteEvent(c *fiber.Ctx) error {
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

	event, err := repositories.GetEventById(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}

	if event.ONGID != user.ID {
		log.Println("" + fmt.Sprintf("User with id '%v' can't delete this event", user.ID))
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse("You can't delete this event"))
	}

	deletedEvent, err := repositories.DeleteEvent(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.EventModelToResponse(deletedEvent)
	return c.JSON(response.MessageResponse("event eliminated successfuly", resp))
}
