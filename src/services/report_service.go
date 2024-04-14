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

// GetAllReports godoc
//
// @Summary Lista todos los informes
// @Description Obtiene una lista paginada de todos los informes.
// @Tags reports
// @Accept json
// @Produce json
// @Param offset query int false "Offset para paginación"
// @Param limit query int false "Límite de resultados por página"
// @Success 200 {object} response.BaseResponsePag[response.ReportResponse] "Respuesta exitosa"
// @Router /reports [get]
func GetAllReports(c *fiber.Ctx) error {
	offset, limit, errors := helpers.ValidatePaginationParams(c.Query("offset", "0"), c.Query("limit", "10"))
	if len(errors) > 0 {
		for _, v := range errors {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(errors))
	}
	totalItems := repositories.CountReports()
	reports, err := repositories.GetAllReports(offset, limit)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.ReportModelsToResponse(reports)
	paguination := common.GeneratePagination(totalItems, limit, int64(offset))
	return c.JSON(response.NewResponsePagination(resp, paguination))
}

// GetReportById godoc
//
// @Summary Obtiene un informe por ID
// @Description Obtiene los detalles de un informe según su ID.
// @Tags reports
// @Accept json
// @Produce json
// @Param id path int true "ID del informe"
// @Success 200 {object} response.BaseResponse[response.ReportResponse] "Respuesta exitosa"
// @Router /reports/{id} [get]
func GetReportById(c *fiber.Ctx) error {
	strid := c.Params("id")
	id, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	report, err := repositories.GetReportById(id)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.ReportModelToResponse(report)
	return c.JSON(response.NewResponse(resp))
}

// CreateReport godoc
//
// @Summary Crea un nuevo informe
// @Description Crea un nuevo informe en la aplicación.
// @Tags reports
// @Accept json
// @Produce json
// @Param reportRequest body request.ReportRequest true "Solicitud de informe"
// @Success 200 {object} response.BaseResponse[response.ReportResponse] "Respuesta exitosa"
// @Router /reports [post]
func CreateReport(c *fiber.Ctx) error {
	model := request.ReportRequest{}
	if _, err := helpers.ValidateRequest(c.Body(), &model); err != nil {
		for _, v := range err {
			log.Println(v)
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorsResponse(err))
	}
	reporterUser, err := repositories.GetUserById(model.ReporterUserID)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	reportedUser, err := repositories.GetUserById(model.ReportedUserID)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse(err.Error()))
	}
	if reporterUser.ID == reportedUser.ID {
		log.Println("you can't report yourself")
	}
	newReport := mapper.ReportRequestToModel(model)
	newReport.ReportedUser = reportedUser
	newReport.ReporterUser = reporterUser
	if _, err := repositories.CreateReport(newReport); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse(err.Error()))
	}
	resp := mapper.ReportModelToResponse(newReport)
	return c.JSON(response.NewResponse(resp))
}
