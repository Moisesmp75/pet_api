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

// GetAllDonationMoney godoc
//
// @Summary Lista todas las donaciones de dinero
//
//	@Security		ApiKeyAuth
//	@Description	Obtiene una lista paginada de todas las donaciones de dinero.
//	@Tags			donations
//	@Accept			json
//	@Produce		json
//	@Param			offset	query		int													false	"Offset para paginación"
//	@Param			limit	query		int													false	"Límite de resultados por página"
//	@Param			ong_id	query		int													false	"Filtrar donaciones por ong"
//	@Success		200		{object}	response.BaseResponsePag[response.DonationMoneyResponse]	"Respuesta exitosa"
//	@Router			/donations/money [get]
func GetAllDonationsMoney(c *fiber.Ctx) error {
	offset, limit, errors := helpers.ValidatePaginationParams(c.Query("offset", "0"), c.Query("limit", "10"))
	if len(errors) > 0 {
		for _, v := range errors {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(errors))
	}
	ong_idStr := c.Query("ong_id", "0")
	ong_id, err := strconv.ParseUint(ong_idStr, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	totalItems := repositories.CountDonationsMoney(ong_id)
	donations, err := repositories.GetAllDonationsMoney(ong_id, offset, limit)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.DonationMoneyModelsToResponse(donations)
	pagination := common.GeneratePagination(totalItems, limit, int64(offset))
	return c.JSON(response.NewResponsePagination(resp, pagination))
}

// CreateDonationMoney godoc
//
//	@Summary		Crea una nueva donacion de dinero
//	@Security		ApiKeyAuth
//	@Description	Crea una nueva donacion de dinero en la aplicación.
//	@Tags			donations
//	@Accept			json
//	@Produce		json
//	@Param			donationMoneyRequest	body		request.DonationMoneyRequest								true	"Donacion"
//	@Success		200				{object}	response.BaseResponse[request.DonationMoneyRequest]	"Respuesta exitosa"
//	@Router			/donations/money [post]
func CreateDonationMoney(c *fiber.Ctx) error {
	model := request.DonationMoneyRequest{}
	if _, err := helpers.ValidateRequest(c.Body(), &model); err != nil {
		for _, v := range err {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(err))
	}
	ong, err := repositories.GetUserById(model.OngID)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	if ong.Role.Name != "ONG" {
		log.Println("this user is not an ONG")
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse("this user is not an ONG"))
	}
	userEmail := c.Locals("user_email").(string)
	user, err := repositories.GetUserByEmail(userEmail)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse(err.Error()))
	}
	if user.ID == ong.ID {
		log.Println("you can't donate yourself")
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse("you can't donate yourself"))
	}
	newDonation := mapper.DonationMoneyRequestToModel(model)
	newDonation.User = user
	newDonation.UserID = user.ID

	if _, err := repositories.CreateDonationMoney(newDonation); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.DonationMoneyModelToResponse(newDonation)
	return c.JSON(response.NewResponse(resp))
}

// GetDonationMoneyById godoc
//
//	@Summary		Obtiene una donación de dinero por ID
//	@Security		ApiKeyAuth
//	@Description	Obtiene los detalles de una donación de dinero según su ID.
//	@Tags			donations
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int														true	"ID de la donación de dinero"
//	@Success		200	{object}	response.BaseResponse[response.DonationMoneyResponse]	"Respuesta exitosa"
//	@Router			/donations/money/{id} [get]
func GetDonationMoneyById(c *fiber.Ctx) error {
	strid := c.Params("id")
	id, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	donation, err := repositories.GetDonationMoneyByID(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.DonationMoneyModelToResponse(donation)
	return c.JSON(response.NewResponse(resp))
}

// UpdateDonationMoney godoc
//
// @Summary Actualiza una donación de dinero
// @Security ApiKeyAuth
// @Description Actualiza una donación de dinero en la aplicación.
// @Tags donations
// @Accept json
// @Produce json
// @Param id path int true "ID de la donación de dinero"
// @Param updateDonationMoneyRequest body request.UpdateDonationRequest true "Datos de la actualización de la donación"
// @Success 200 {object} response.BaseResponse[response.DonationMoneyResponse] "Respuesta exitosa"
// @Router /donations/money/{id} [patch]
func UpdateDonationMoney(c *fiber.Ctx) error {
	strid := c.Params("id")
	id, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}
	model := request.UpdateDonationMoneyRequest{}
	if _, err := helpers.ValidateRequest(c.Body(), &model); err != nil {
		for _, v := range err {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(err))
	}
	donation, err := repositories.GetDonationMoneyByID(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	userEmail := c.Locals("user_email").(string)
	user, _ := repositories.GetUserByEmail(userEmail)
	if donation.OngID != user.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse("You can't update this donation"))
	}
	updateDonation := mapper.UpdateDonationMoneyRequestToModel(model, donation)
	if _, err := repositories.UpdateDonationMoney(updateDonation); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.DonationMoneyModelToResponse(updateDonation)
	if !updateDonation.Received {
		_, err := repositories.DeleteDonationMoney(updateDonation.ID)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
		}
	}
	return c.JSON(response.MessageResponse("update adoption successfully", resp))
}
