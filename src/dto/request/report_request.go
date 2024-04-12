package request

type ReportRequest struct {
	ReporterUserID uint64 `json:"reporter_user_id" validate:"required"`
	ReportedUserID uint64 `json:"reported_user_id" validate:"required"`
	Description    string `json:"description" validate:"required,gt=0"`
}
