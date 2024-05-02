package services

import (
	"log"
	"pet_api/src/auth"
	"pet_api/src/common"
	"pet_api/src/dto/request"
	"pet_api/src/dto/response"
	"pet_api/src/helpers"
	"pet_api/src/mapper"
	"pet_api/src/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetAllUsers godoc
//
//	@Summary		Lista todos los usuarios
//	@Security		ApiKeyAuth
//	@Description	Lista todos los usuarios de la aplicación.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			offset	query		int												false	"Offset de paginación"
//	@Param			limit	query		int												false	"Límite de resultados por página"
//	@Success		200		{object}	response.BaseResponsePag[response.UserResponse]	"Respuesta exitosa"
//	@Router			/users [get]
func GetAllUsers(c *fiber.Ctx) error {
	offset, limit, errors := helpers.ValidatePaginationParams(c.Query("offset", "0"), c.Query("limit", "10"))
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

// GetUserById godoc
//
//	@Summary		Muestra un usuario
//	@Security		ApiKeyAuth
//	@Description	Muestra un usuario con el ID especificado.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int												true	"ID del usuario"
//	@Success		200	{object}	response.BaseResponse[response.UserResponse]	"Respuesta exitosa"
//	@Router			/users/{id} [get]
func GetUserById(c *fiber.Ctx) error {
	strid := c.Params("id")
	id, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}
	user, err := repositories.GetUserById(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.UserModelToResponse(user)
	return c.JSON(response.NewResponse(resp))
}

// GetSelfUser godoc
//
//	@Summary		Muestra al mismo usuario que hizo la peticion
//	@Security		ApiKeyAuth
//	@Description	Muestra un usuario propio que esta logeado actualmente
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.BaseResponse[response.UserResponse]	"Respuesta exitosa"
//	@Router			/users/self [get]
func GetSelfUser(c *fiber.Ctx) error {
	userEmail := c.Locals("user_email").(string)
	user, err := repositories.GetUserByEmail(userEmail)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}

	if _, err := repositories.GetUserById(user.ID); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.UserModelToResponse(user)
	return c.JSON(response.NewResponse(resp))
}

// CreateUser godoc
//
//	@Summary		Crea un nuevo usuario
//	@Description	Crea un nuevo usuario en la aplicación.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			userRequest	body		request.UserRequest								true	"Solicitud de creación de usuario"
//	@Success		200			{object}	response.BaseResponse[response.UserResponse]	"Respuesta exitosa"
//	@Router			/users/register [post]
func CreateUser(c *fiber.Ctx) error {
	model := request.UserRequest{}
	if _, err := helpers.ValidateRequest(c.Body(), &model); err != nil {
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

// LoginUser godoc
//
//	@Summary		Inicia sesión de usuario
//	@Description	Inicia sesión de usuario con credenciales proporcionadas.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			loginRequest	body		request.LoginRequest							true	"Solicitud de inicio de sesión"
//	@Success		200				{object}	response.BaseResponse[response.LoginResponse]	"Respuesta exitosa"
//	@Router			/users/login [post]
func LoginUser(c *fiber.Ctx) error {
	model := request.LoginRequest{}
	if _, err := helpers.ValidateRequest(c.Body(), &model); err != nil {
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
	lgResp := mapper.UserResponseToLoginResponse(resp)
	token, err := auth.GenerateToken(lgResp)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	lgResp.Token = token

	c.Cookie(&fiber.Cookie{
		Name:        "token",
		Value:       token,
		Expires:     lgResp.Exp,
		Secure:      true,
		HTTPOnly:    true,
		SessionOnly: false,
	})
	return c.JSON(response.NewResponse(lgResp))
}

// UpdateUserImage godoc
//
//	@Summary		Actualiza la imagen de usuario
//	@Security		ApiKeyAuth
//	@Description	Actualiza la imagen de usuario identificado por su ID.
//	@Tags			users
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			user_img	formData	file											true	"Imagen de usuario"
//	@Success		200			{object}	response.BaseResponse[response.UserResponse]	"Respuesta exitosa"
//	@Router			/users/img [patch]
func UpdateUserImage(c *fiber.Ctx) error {
	userEmail := c.Locals("user_email").(string)
	user, err := repositories.GetUserByEmail(userEmail)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}

	file, err := c.FormFile("user_img")
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}
	file.Filename = "user_" + user.Role.Name + "_" + "image_" + strconv.FormatUint(user.ID, 10)
	url_img, _, err := helpers.UploadFile(file, "user_images/", false)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}

	user.ImageUrl = url_img

	if _, err := repositories.UpdateUser(user); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}

	resp := mapper.OnlyUserModelToResponse(user)
	return c.JSON(response.MessageResponse("user updated successfully", resp))
}

// RecoverPassword godoc
//
//	@Summary		Recupera la contraseña de usuario
//	@Description	Envía un correo electrónico con una nueva contraseña generada.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			passwordResetRequest	body		request.PasswordResetRequest					true	"Solicitud de restablecimiento de contraseña"
//	@Success		200						{object}	response.BaseResponse[response.UserResponse]	"Respuesta exitosa"
//	@Router			/users/recover [post]
func RecoverPassword(c *fiber.Ctx) error {
	req := request.PasswordResetRequest{}
	if _, err := helpers.ValidateRequest(c.Body(), &req); err != nil {
		for _, v := range err {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(err))
	}
	user, err := repositories.GetUserByEmail(req.Email)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}

	newPassword := helpers.GeneratePassword(10, 1, 1, 1)

	if err := common.SendResetPasswordEmail(user, newPassword); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}

	user.Password = auth.Encrypt_password(newPassword)
	if _, err := repositories.UpdateUser(user); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.OnlyUserModelToResponse(user)
	return c.JSON(response.MessageResponse("check your email", resp))
}

// UpdateUser godoc
//
//	@Summary		Actualiza los detalles de usuario desde otro usuario con un rol superior
//	@Security		ApiKeyAuth
//	@Description	Actualiza los detalles de usuario identificado por su ID. Solo los de un rol superior pueden utilizar esta endpoint
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id					path		int												true	"ID del usuario"
//	@Param			updateUserRequest	body		request.UpdateUserRequest						true	"Solicitud de actualización de usuario"
//	@Success		200					{object}	response.BaseResponse[response.UserResponse]	"Respuesta exitosa"
//	@Router			/users/{id} [patch]
func UpdateUser(c *fiber.Ctx) error {
	strid := c.Params("id")
	id, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}
	model := request.UpdateUserRequest{}
	if _, err := helpers.ValidateRequest(c.Body(), &model); err != nil {
		for _, v := range err {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(err))
	}

	userEmail := c.Locals("user_email").(string)
	if _, err := repositories.GetUserByEmail(userEmail); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}

	user, err := repositories.GetUserById(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse(err.Error()))
	}

	if model.Password != "" {
		model.Password = auth.Encrypt_password(model.Password)
	}
	updateUser := mapper.UpdateUserRequestToModel(model, user)

	if _, err := repositories.UpdateUser(updateUser); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.OnlyUserModelToResponse(updateUser)
	return c.JSON(response.MessageResponse("user updated successfully", resp))
}

// DeletePet godoc
//
//	@Summary		Elimina un usuario desde otro usuario con un rol superior
//	@Security		ApiKeyAuth
//	@Description	Elimina un usuario identificada por su ID. Solo los de un rol superior pueden utilizar este endpoint
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User id"
//	@Success		200	{object}	response.BaseResponse[response.UserResponse]
//	@Router			/users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	strid := c.Params("id")
	id, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse(err.Error()))
	}

	userEmail := c.Locals("user_email").(string)
	if _, err := repositories.GetUserByEmail(userEmail); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse(err.Error()))
	}

	if _, err := repositories.GetUserById(id); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}

	deletedUser, err := repositories.DeleteUser(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.UserModelToResponse(deletedUser)
	return c.JSON(response.MessageResponse("user eliminated successfully", resp))
}

// DeletePet godoc
//
//	@Summary		Elimina un usuario
//	@Security		ApiKeyAuth
//	@Description	Elimina un usuario identificada por su ID.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.BaseResponse[response.UserResponse]
//	@Router			/users [delete]
func DeleteSelfUser(c *fiber.Ctx) error {
	userEmail := c.Locals("user_email").(string)
	user, err := repositories.GetUserByEmail(userEmail)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse(err.Error()))
	}

	deletedUser, err := repositories.DeleteUser(user.ID)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.UserModelToResponse(deletedUser)
	return c.JSON(response.MessageResponse("user eliminated successfully", resp))
}

// UpdateUser godoc
//
//	@Summary		Actualiza los detalles de usuario
//	@Security		ApiKeyAuth
//	@Description	Actualiza los detalles de usuario identificado por su ID.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			updateUserRequest	body		request.UpdateUserRequest						true	"Solicitud de actualización de usuario"
//	@Success		200					{object}	response.BaseResponse[response.UserResponse]	"Respuesta exitosa"
//	@Router			/users [patch]
func UpadteSelfUser(c *fiber.Ctx) error {
	model := request.UpdateUserRequest{}
	if _, err := helpers.ValidateRequest(c.Body(), &model); err != nil {
		for _, v := range err {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(err))
	}

	userEmail := c.Locals("user_email").(string)
	user, err := repositories.GetUserByEmail(userEmail)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}

	if model.Password != "" {
		model.Password = auth.Encrypt_password(model.Password)
	}
	updateUser := mapper.UpdateUserRequestToModel(model, user)

	if _, err := repositories.UpdateUser(updateUser); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.OnlyUserModelToResponse(updateUser)
	return c.JSON(response.MessageResponse("user updated successfully", resp))
}
