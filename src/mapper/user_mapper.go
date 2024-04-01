package mapper

import (
	"pet_api/src/dto/request"
	"pet_api/src/dto/response"
	"pet_api/src/models"
)

func UserRequestToModel(req request.UserRequest) models.User {
	return models.User{
		Username: req.Username,
		Email: req.Email,
		Password: req.Password,
	}
}

func UserModelToResponse(user models.User) response.UserResponse {
	return response.UserResponse{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
	}
}