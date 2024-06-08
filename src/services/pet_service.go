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
//	@Param			offset		query		int		false	"Offset de paginación"
//	@Param			limit		query		int		false	"Límite de resultados por página"
//	@Param			breed		query		string	false	"Filtrar mascotas por raza"
//	@Param			color		query		string	false	"Filtrar mascotas por color"
//	@Param			gender		query		string	false	"Filtrar mascotas por genero"
//	@Param			pet_type	query		string	false	"Filtrar mascotas por tipo"
//	@Success		200			{object}	response.BaseResponsePag[response.PetResponse]
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
	gender := c.Query("gender", "")
	petType := c.Query("pet_type", "")
	totalItems := repositories.CountPets(breed, color, gender, petType)
	pets, err := repositories.GetAllPets(offset, limit, breed, color, gender, petType)
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
//	@Security		ApiKeyAuth
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
	userEmail := c.Locals("user_email").(string)
	user, err := repositories.GetUserByEmail(userEmail)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	pet.User = user
	pet.UserID = user.ID
	pet.Image.Filename = helpers.GenerateUniqueFileName("user_" + strconv.FormatUint(user.ID, 10) + "_pet-image")
	pet.Image.URL = "https://firebasestorage.googleapis.com/v0/b/hairypets.appspot.com/o/pet_images%2Fdefault_pet.png?alt=media&token=a13e363e-e8bf-4eda-90f6-828e8b628050"
	pet.Adopted = false
	petCreated, err := repositories.CreatePet(pet)

	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}

	resp := mapper.PetModelToResponse(petCreated)
	return c.JSON(response.NewResponse(resp))
}

// UpdatePet godoc
//
//	@Summary		Actualiza los detalles de una mascota
//	@Security		ApiKeyAuth
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
	userEmail := c.Locals("user_email").(string)
	user, _ := repositories.GetUserByEmail(userEmail)
	if user.ID != pet.UserID {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse("You can't update this pet"))
	}
	updatePet := mapper.UpdatePetRequestToModel(model, pet)

	if _, err := repositories.UpdatePet(updatePet); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.OnlyPetModelToResponse(updatePet)
	return c.JSON(response.MessageResponse("pet updated successfully", resp))
}

// UpdatePetImages godoc
//
//	@Summary		Actualiza las imágenes de una mascota
//	@Security		ApiKeyAuth
//	@Description	Actualiza la imagen de una mascota identificada por su ID.
//	@Tags			pets
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			id		path		int		true	"Pet id"
//	@Param			pet_img	formData	file	true	"Imagen de la mascota"
//	@Success		200		{object}	response.BaseResponse[response.PetResponse]
//	@Router			/pets/{id}/img [patch]
func UpdatePetImage(c *fiber.Ctx) error {
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
	userEmail := c.Locals("user_email").(string)
	user, _ := repositories.GetUserByEmail(userEmail)
	if user.ID != pet.UserID {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse("You can't update this pet"))
	}

	file, err := c.FormFile("pet_img")
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}
	url_img, _, err := helpers.UploadFile(file, "pet_images/", false)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}

	pet_img := pet.Image
	pet_img.URL = url_img

	if _, err := repositories.UpdatePetImage(pet_img); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	pet.Image.URL = url_img
	resp := mapper.OnlyPetModelToResponse(pet)
	return c.JSON(response.MessageResponse("images created successfully", resp))
}

// DeletePet godoc
//
//	@Summary		Elimina una mascota
//	@Security		ApiKeyAuth
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
	pet, err := repositories.GetPetById(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	userEmail := c.Locals("user_email").(string)
	user, _ := repositories.GetUserByEmail(userEmail)
	if user.ID != pet.UserID {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse("You can't delete this pet"))
	}
	deletePet, err := repositories.DeletePet(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.OnlyPetModelToResponse(deletePet)
	return c.JSON(response.MessageResponse("pet eliminated successfully", resp))
}
