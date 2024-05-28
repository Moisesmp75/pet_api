package services

import (
	"log"
	"pet_api/src/dto/response"
	"pet_api/src/mapper"
	"pet_api/src/repositories"

	"github.com/gofiber/fiber/v2"
)

// GetAllRoles godoc
//
//	@Summary		Lista todos los tipos de mascotas
//	@Description	Obtiene todos los tipos de mascotas.
//	@Security		ApiKeyAuth
//	@Tags			pets
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.BaseResponse[[]response.PetTypeResponse]	"Respuesta exitosa"
//	@Router			/pets/types [get]
func GetAllPetTypes(c *fiber.Ctx) error {
	petTypes, err := repositories.GetAllPetTypes()
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.PetTypeModelsToResponse(petTypes)
	return c.JSON(response.NewResponse(resp))
}
