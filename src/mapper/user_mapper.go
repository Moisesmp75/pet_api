package mapper

import (
	"pet_api/src/dto/request"
	"pet_api/src/dto/response"
	"pet_api/src/models"
	"time"
)

func UserRequestToModel(req request.UserRequest) models.User {
	return models.User{
		Name: 			 req.Name,
		LastName: 	 req.LastName,
		PhoneNumber: req.PhoneNumber,
		Email:    	 req.Email,
		Password: 	 req.Password,
	}
}

func OnlyUserModelToResponse(user models.User) *response.UserResponse {
	return &response.UserResponse{
		ID:       	 user.ID,
		Name: 			 user.Name,
		LastName: 	 user.LastName,
		PhoneNumber: user.PhoneNumber,
		Email:    	 user.Email,
		Role:     	 user.Role.Name,
		ImageUrl:    user.ImageUrl,
	}
}

func UserModelToResponse(user models.User) response.UserResponse {
	return response.UserResponse{
		ID:      		 user.ID,
		Name: 			 user.Name,
		LastName: 	 user.LastName,
		PhoneNumber: user.PhoneNumber,
		Email:    	 user.Email,
		Pets:     	 OnlyPetsModelsToResponse(user.Pets),
		Role:     	 user.Role.Name,
		ImageUrl:    user.ImageUrl,
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
		Email:  userResp.Email,
		Role:   userResp.Role,
		Token:	"",
		Iat: 		time.Now(),
		Exp: 		time.Now().Add(12 * time.Hour),
	}
}
