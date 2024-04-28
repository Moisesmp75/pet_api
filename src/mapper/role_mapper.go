package mapper

import (
	"pet_api/src/dto/response"
	"pet_api/src/models"
)

func RoleModelToResponse(role models.Role) response.RoleResponse {
	role_name := role.Name
	if role.Name == "Duenio" {
		role_name = "Dueño"
	}
	return response.RoleResponse{
		ID:          role.ID,
		Name:        role_name,
		Description: role.Description,
	}
}

func RoleModelsToResponse(roles []models.Role) []response.RoleResponse {
	resp := make([]response.RoleResponse, len(roles))

	for i, v := range roles {
		resp[i] = RoleModelToResponse(v)
	}

	return resp
}
