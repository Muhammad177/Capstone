package dto

import (
	"Capstone/models"
	"time"
)

type GetReportResponse struct {
	ID             int                    `json:"id"`
	Reason         string                 `json:"reason"`
	Reporter       ReporterResponse       `json:"reporter"`
	ReportedThread ReportedTheardResponse `json:"reportedThread"`
	CreateAt       time.Time              `json:"createdAt"`
	UpdateAt       time.Time              `json:"updateAt"`
}

func NewGetReportResponse(data models.Report) *GetReportResponse {
	return &GetReportResponse{
		ID:             int(data.ID),
		Reason:         data.Reason,
		Reporter:       *NewReporterResponse(data.ReportUser),
		ReportedThread: *newReportedTheardResponse(data.ReportThread),
		CreateAt:       data.CreatedAt,
		UpdateAt:       data.UpdatedAt,
	}
}

type ReportedTheardResponse struct {
	ID    int    `json: "thread_ID"`
	Title string `json: "title"`
}

func newReportedTheardResponse(reportedThread models.Thread) *ReportedTheardResponse {
	return &ReportedTheardResponse{
		ID:    int(reportedThread.ID),
		Title: reportedThread.Title,
	}
}

type ReporterResponse struct {
	ID       int    `json:"user_ID"`
	Username string `json:"username"`
}

func NewReporterResponse(reporter models.User) *ReporterResponse {
	return &ReporterResponse{
		ID:       int(reporter.ID),
		Username: reporter.Username,
	}
}

type GetReportsResponse []GetReportResponse

func NewGetReportsResponse(data []models.Report) *GetReportsResponse {
	result := GetReportsResponse{}

	for _, each := range data {
		result = append(result, *NewGetReportResponse(each))
	}

	return &result
}
