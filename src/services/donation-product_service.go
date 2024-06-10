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

// GetAllDonationProducts godoc
//
// @Summary Lista todas las donaciones de productos
//
//	@Security		ApiKeyAuth
//	@Description	Obtiene una lista paginada de todas las donaciones.
//	@Tags			donations
//	@Accept			json
//	@Produce		json
//	@Param			offset	query		int													false	"Offset para paginación"
//	@Param			limit	query		int													false	"Límite de resultados por página"
//	@Param			ong_id	query		int													false	"Filtrar donaciones por ong"
//	@Success		200		{object}	response.BaseResponsePag[response.DonationProductResponse]	"Respuesta exitosa"
//	@Router			/donations/products [get]
func GetAllDonationsProduct(c *fiber.Ctx) error {
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
	totalItems := repositories.CountDonationsProduct(ong_id)
	donations, err := repositories.GetAllDonationsProduct(ong_id, offset, limit)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.DonationProductModelsToResponse(donations)
	pagination := common.GeneratePagination(totalItems, limit, int64(offset))
	return c.JSON(response.NewResponsePagination(resp, pagination))
}

// CreateDpnationProduct godoc
//
//	@Summary		Crea una nueva donacion
//	@Security		ApiKeyAuth
//	@Description	Crea una nueva donacion en la aplicación.
//	@Tags			donations
//	@Accept			json
//	@Produce		json
//	@Param			donationProductRequest	body		request.DonationProductRequest								true	"Donacion"
//	@Success		200				{object}	response.BaseResponse[request.DonationProductRequest]	"Respuesta exitosa"
//	@Router			/donations/products [post]
func CreateDonationProduct(c *fiber.Ctx) error {
	model := request.DonationProductRequest{}
	if _, err := helpers.ValidateRequest(c.Body(), &model); err != nil {
		for _, v := range err {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(err))
	}
	ong, err := repositories.GetUserById(model.OngId)
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
	newDonation := mapper.DonationProductRequestToModel(model)
	newDonation.User = user
	newDonation.UserID = user.ID

	if _, err := repositories.CreateDonationProduct(newDonation); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.DonationProductModelToResponse(newDonation)
	return c.JSON(response.NewResponse(resp))
}

// GetDonationProductById godoc
//
//	@Summary		Obtiene una donación de producto por ID
//	@Security		ApiKeyAuth
//	@Description	Obtiene los detalles de una donación de producto según su ID.
//	@Tags			donations
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int														true	"ID de la donación de productos"
//	@Success		200	{object}	response.BaseResponse[response.DonationProductResponse]	"Respuesta exitosa"
//	@Router			/donations/products/{id} [get]
func GetDonationProductById(c *fiber.Ctx) error {
	strid := c.Params("id")
	id, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	donation, err := repositories.GetDonationProductById(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.DonationProductModelToResponse(donation)
	return c.JSON(response.NewResponse(resp))
}
