package response

import "time"

type ReportResponse struct {
	ID           uint64       `json:"id"`
	ReporterUser UserResponse `json:"reporter"`
	ReportedUser UserResponse `json:"reported_user"`
	Description  string       `json:"description"`
	ReportDate   time.Time    `json:"report_date"`
}
