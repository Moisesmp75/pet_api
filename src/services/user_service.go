package services

import (
	"log"
	"pet_api/src/auth"
	"pet_api/src/common"
	"pet_api/src/dto/request"
	"pet_api/src/dto/response"
	"pet_api/src/mapper"
	"pet_api/src/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	offset, limit, errors := ValidatePaginationParams(c.Query("offset", "0"), c.Query("limit", "10"))
	if len(errors) > 0 {
		for _, v := range errors {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(errors))
	}

	totalItems := repositories.CountUsers()
	users, err := repositories.GetAllUsers(offset, limit)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.UsersModelsToResponse(users)
	pagination := common.GeneratePagination(totalItems, limit, int64(offset))

	return c.JSON(response.NewResponsePagination(resp, pagination))
}

func GetUserById(c *fiber.Ctx) error {
	strid := c.Params("id")
	id, err := strconv.Atoi(strid)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	user, err := repositories.GetUserById(uint(id))
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.UserModelToResponse(user)
	return c.JSON(response.NewResponse(resp))
}

func CreateUser(c *fiber.Ctx) error {
	model := request.UserRequest{}
	if _, err := common.ValidateRequest(c.Body(), &model); err != nil {
		for _, v := range err {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(err))
	}
	rol, err := repositories.GetRoleById(model.RoleID)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	user := mapper.UserRequestToModel(model)
	user.Role = rol
	user.Password = auth.Encrypt_password(user.Password)
	userCreated, err := repositories.CreateUser(user)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.UserModelToResponse(userCreated)
	return c.JSON(response.NewResponse(resp))
}

func LoginUser(c *fiber.Ctx) error {
	model := request.LoginRequest{}
	if _, err := common.ValidateRequest(c.Body(), &model); err != nil {
		for _, v := range err {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(err))
	}
	user, err := repositories.GetUserByEmailOrPhone(model.Identity)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	if user.Role.ID != model.RoleID {
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse("Access Denied: Your account does not have the necessary privileges. Please make sure you are using the correct credentials for your user type."))
	}
	if !auth.DecryptPasswordHash(user.Password, model.Password) {
		log.Println("Invalid Credentials")
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse("Invalid Credentials"))
	}

	resp := mapper.UserModelToResponse(user)
	token, err := auth.GenerateToken(resp)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	lgResp := mapper.UserResponseToLoginResponse(resp)
	lgResp.Token = token

	c.Cookie(&fiber.Cookie{
		Name:        "token",
		Value:       token,
		Secure:      true,
		SessionOnly: false,
	})
	return c.JSON(response.NewResponse(lgResp))
}
