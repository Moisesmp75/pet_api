package mapper

import (
	"pet_api/src/dto/request"
	"pet_api/src/dto/response"
	"pet_api/src/models"
)

func ReportRequestToModel(req request.ReportRequest) models.Report {
	return models.Report{
		ReporterUserID: req.ReporterUserID,
		ReportedUserID: req.ReportedUserID,
		Description:    req.Description,
	}
}

func ReportModelToResponse(report models.Report) response.ReportResponse {
	return response.ReportResponse{
		ID:           report.ID,
		ReporterUser: *OnlyUserModelToResponse(report.ReporterUser),
		ReportedUser: *OnlyUserModelToResponse(report.ReportedUser),
		ReportDate:   report.CreatedAt,
		Description:  report.Description,
	}
}

func ReportModelsToResponse(reports []models.Report) []response.ReportResponse {
	resp := make([]response.ReportResponse, len(reports))

	for i, v := range reports {
		resp[i] = ReportModelToResponse(v)
	}

	return resp
}
