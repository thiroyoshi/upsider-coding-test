package invoices

import (
	val "upsider-coding-test/internal/validation"

	"github.com/gin-gonic/gin"
)

type GetRequest struct {
	StartAt int `json:"startAt"`
	EndAt   int `json:"endAt"`
}

func NewGetRequest(ctx *gin.Context) (GetRequest, error) {
	startAtStr := ctx.Query("startAt")
	startAt, err := val.ValidateDateFormat(startAtStr)
	if err != nil {
		return GetRequest{}, err
	}

	endAtStr := ctx.Query("endAt")
	endAt, err := val.ValidateDateFormat(endAtStr)
	if err != nil {
		return GetRequest{}, err
	}

	if err := val.ValidateDateRange(startAtStr, endAtStr); err != nil {
		return GetRequest{}, err
	}

	return GetRequest{
		StartAt: startAt,
		EndAt:   endAt,
	}, nil
}
