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
//	@Summary		Lista todos los roles
//	@Description	Obtiene todos los roles.
//	@Tags			roles
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	response.BaseResponse[[]response.RoleResponse]	"Respuesta exitosa"
//	@Router			/roles [get]
func GetAllRoles(c *fiber.Ctx) error {
	roles, err := repositories.GetAllRoles()
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}

	resp := mapper.RoleModelsToResponse(roles)
	return c.JSON(response.NewResponse(resp))
}
