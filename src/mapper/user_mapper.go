package mapper

import (
	"pet_api/src/dto/request"
	"pet_api/src/dto/response"
	"pet_api/src/models"
	"time"
)

func UserRequestToModel(req request.UserRequest) models.User {
	return models.User{
		Name:           req.Name,
		LastName:       req.LastName,
		Username:       req.UserName,
		PhoneNumber:    req.PhoneNumber,
		Email:          req.Email,
		Password:       req.Password,
		Dni:            req.Dni,
		Address:        req.Address,
		City:           req.City,
		MotherLastName: req.MotherLastName,
		About:          req.About,
	}
}

func OnlyUserModelToResponse(user models.User) *response.UserResponse {
	return &response.UserResponse{
		ID:             user.ID,
		Name:           user.Name,
		LastName:       user.LastName,
		MotherLastName: user.MotherLastName,
		UserName:       user.Username,
		PhoneNumber:    user.PhoneNumber,
		Dni:            user.Dni,
		Address:        user.Address,
		City:           user.City,
		Email:          user.Email,
		Role:           user.Role.Name,
		ImageUrl:       user.Image.URL,
		About:          user.About,
	}
}

func UserModelToResponse(user models.User) response.UserResponse {
	return response.UserResponse{
		ID:             user.ID,
		Name:           user.Name,
		LastName:       user.LastName,
		MotherLastName: user.MotherLastName,
		UserName:       user.Username,
		PhoneNumber:    user.PhoneNumber,
		Dni:            user.Dni,
		Address:        user.Address,
		City:           user.City,
		Email:          user.Email,
		Pets:           OnlyPetsModelsToResponse(user.Pets),
		Role:           user.Role.Name,
		ImageUrl:       user.Image.URL,
		About:          user.About,
	}
}

func UsersModelsToResponse(users []models.User) []response.UserResponse {
	resp := make([]response.UserResponse, len(users))

	for i, v := range users {
		resp[i] = UserModelToResponse(v)
	}

	return resp
}

func OnlyUserModelsToResponse(users []models.User) []response.UserResponse {
	resp := make([]response.UserResponse, len(users))

	for i, v := range users {
		resp[i] = *OnlyUserModelToResponse(v)
	}

	return resp
}

func UserResponseToLoginResponse(userResp response.UserResponse) response.LoginResponse {
	return response.LoginResponse{
		Email: userResp.Email,
		Role:  userResp.Role,
		Token: "",
		Iat:   time.Now(),
		Exp:   time.Now().Add(12 * time.Hour),
	}
}

func UpdateUserRequestToModel(req request.UpdateUserRequest, user models.User) models.User {
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.UserName != "" {
		user.Username = req.UserName
	}
	if req.PhoneNumber != "" {
		user.PhoneNumber = req.PhoneNumber
	}
	if req.Password != "" {
		user.Password = req.Password
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Address != "" {
		user.Address = req.Address
	}
	if req.City != "" {
		user.City = req.City
	}
	if req.MotherLastName != "" {
		user.MotherLastName = req.MotherLastName
	}
	if req.About != "" {
		user.About = req.About
	}
	return user
}
