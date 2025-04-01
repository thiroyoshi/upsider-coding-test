package invoices

import (
	val "upsider-coding-test/internal/validation"

	"github.com/gin-gonic/gin"
)

type GetRequest struct {
	StartAt string `json:"startAt"`
	EndAt   string `json:"endAt"`
}

func NewGetRequest(ctx *gin.Context) (GetRequest, error) {
	startAtStr := ctx.Query("startAt")
	if err := val.ValidateDateFormat(startAtStr); err != nil {
		return GetRequest{}, err
	}

	endAtStr := ctx.Query("endAt")
	if err := val.ValidateDateFormat(endAtStr); err != nil {
		return GetRequest{}, err
	}

	if err := val.ValidateDateRange(startAtStr, endAtStr); err != nil {
		return GetRequest{}, err
	}

	return GetRequest{
		StartAt: startAtStr,
		EndAt:   endAtStr,
	}, nil
}
