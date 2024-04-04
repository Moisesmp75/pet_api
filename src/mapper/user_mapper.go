package mapper

import (
	"pet_api/src/dto/request"
	"pet_api/src/dto/response"
	"pet_api/src/models"
)

func UserRequestToModel(req request.UserRequest) models.User {
	return models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

func OnlyUserModelToResponse(user models.User) *response.UserResponse {
	return &response.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}

func UserModelToResponse(user models.User) response.UserResponse {
	return response.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Pets:     OnlyPetsModelsToResponse(user.Pets),
	}
}

func UsersModelsToResponse(users []models.User) []response.UserResponse {
	resp := make([]response.UserResponse, len(users))

	for i, v := range users {
		resp[i] = UserModelToResponse(v)
	}

	return resp
}

func UserResponseToLoginResponse(userResp response.UserResponse) response.LoginResponse {
	return response.LoginResponse{
		Email:    userResp.Email,
		Username: userResp.Username,
		Token:    "",
	}
}
