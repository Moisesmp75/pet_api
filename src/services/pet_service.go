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

// GetAllPets godoc
//
//	@Summary		Lista a todas las mascotas
//	@Description	Lista todas las mascotas de la aplicación.
//	@Tags			pets
//	@Accept			json
//	@Produce		json
//	@Param			offset	query		int	false	"Offset de paginación"
//	@Param			limit	query		int	false	"Límite de resultados por página"
//	@Param			breed	query		int	false	"Filtrar mascota por raza"
//	@Param			color	query		int	false	"Filtrar mascota por color"
//	@Success		200		{object}	response.BaseResponsePag[response.PetResponse]
//	@Router			/pets [get]
func GetAllPets(c *fiber.Ctx) error {
	offset, limit, errors := helpers.ValidatePaginationParams(c.Query("offset", "0"), c.Query("limit", "10"))
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

// GetAllPets godoc
//
//	@Summary		Mostrar a una mascota
//	@Description	Muestra una mascota con el id.
//	@Tags			pets
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Pet id"
//	@Success		200	{object}	response.BaseResponse[response.PetResponse]
//	@Router			/pets/{id} [get]
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

// CreatePet godoc
//
//	@Summary		Crea una nueva mascota
//	@Description	Crea una nueva mascota en la aplicación.
//	@Tags			pets
//	@Accept			json
//	@Produce		json
//	@Param			petRequest	body		request.PetRequest	true	"Pet request payload"
//	@Success		200			{object}	response.BaseResponse[response.PetResponse]
//	@Router			/pets [post]
func CreatePet(c *fiber.Ctx) error {
	model := request.PetRequest{}
	if _, err := helpers.ValidateRequest(c.Body(), &model); err != nil {
		for _, v := range err {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(err))
	}
	petType, err := repositories.GetPetTypeById(model.PetTypeId)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	pet := mapper.PetRequestToModel(model)
	pet.PetType = petType
	user, err := repositories.GetUserById(model.UserID)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	pet.User = user
	petCreated, err := repositories.CreatePet(pet)

	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}

	resp := mapper.PetModelToResponse(petCreated)
	return c.JSON(response.NewResponse(resp))
}

// UpdatePetImages godoc
//
//	@Summary		Actualiza las imágenes de una mascota
//	@Description	Actualiza las imágenes de una mascota identificada por su ID.
//	@Tags			pets
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			id			path		int		true	"Pet id"
//
//	@Param			img_1	formData	file											true	"Imagen 1 de la mascota"
//	@Param			img_2	formData	file											false	"Imagen 2 de la mascota"
//	@Param			img_3	formData	file											false	"Imagen 3 de la mascota"
//	@Param			img_4	formData	file											false	"Imagen 4 de la mascota"
//	@Param			img_5	formData	file											false	"Imagen 5 de la mascota"
//
//	@Success		200			{object}	response.BaseResponse[response.PetResponse]
//	@Router			/pets/{id}/img [patch]
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
	images, err := CreatePetImages(pet.ID, form)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}
	pet.Images = images
	resp := mapper.OnlyPetModelToResponse(pet)
	return c.JSON(response.MessageResponse("images created successfully", resp))
}

// UpdatePet godoc
//
//	@Summary		Actualiza los detalles de una mascota
//	@Description	Actualiza los detalles de una mascota identificada por su ID.
//	@Tags			pets
//	@Accept			json
//	@Produce		json
//	@Param			id					path		int							true	"Pet id"
//	@Param			updatePetRequest	body		request.UpdatePetRequest	true	"Pet update request payload"
//	@Success		200					{object}	response.BaseResponse[response.PetResponse]
//	@Router			/pets/{id} [patch]
func UpdatePet(c *fiber.Ctx) error {
	strid := c.Params("id")
	id, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}
	model := request.UpdatePetRequest{}
	if _, err := helpers.ValidateRequest(c.Body(), &model); err != nil {
		for _, v := range err {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(err))
	}
	pet, err := repositories.GetPetById(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	updatePet := mapper.UpdatePetRequestToModel(model, pet)

	if _, err := repositories.UpdatePet(updatePet); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.OnlyPetModelToResponse(updatePet)
	return c.JSON(response.MessageResponse("pet updated successfully", resp))
}

// DeletePet godoc
//
//	@Summary		Elimina una mascota
//	@Description	Elimina una mascota identificada por su ID.
//	@Tags			pets
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Pet id"
//	@Success		200	{object}	response.BaseResponse[response.PetResponse]
//	@Router			/pets/{id} [delete]
func DeletePet(c *fiber.Ctx) error {
	strid := c.Params("id")
	id, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}
	pet, err := repositories.DeletePet(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.OnlyPetModelToResponse(pet)
	return c.JSON(response.MessageResponse("pet eliminated successfully", resp))
}
