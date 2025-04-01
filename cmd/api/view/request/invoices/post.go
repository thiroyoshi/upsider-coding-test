package invoices

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	val "upsider-coding-test/internal/validation"
)

type PostRequest struct {
	PartnerID  string `json:"partnerId"`
	Amount     int    `json:"amount"`
	PaymentDue string `json:"paymentDue"`
}

func NewPostRequest(ctx *gin.Context) (PostRequest, error) {
	var req PostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.Warn("failed to bind json", "error", err)
		return PostRequest{}, err
	}

	// Check if request body is empty
	if req.PartnerID == "" || req.Amount == 0 || req.PaymentDue == "" {
		return PostRequest{}, fmt.Errorf("request body is empty")
	}

	// Validate partner ID is UUID
	if err := val.ValidateUUID(req.PartnerID); err != nil {
		return PostRequest{}, err
	}

	// Validate amount
	if err := val.ValidateAmount(req.Amount); err != nil {
		return PostRequest{}, err
	}

	// Validate payment due date format
	if err := val.ValidateDateFormat(req.PaymentDue); err != nil {
		return PostRequest{}, err
	}

	// Validate payment due date is in the future
	if err := val.ValidatePastDate(req.PaymentDue); err != nil {
		return PostRequest{}, err
	}

	return req, nil
}
