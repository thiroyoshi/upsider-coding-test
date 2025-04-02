package invoices

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"upsider-coding-test/cmd/api/model"
	requestView "upsider-coding-test/cmd/api/view/request/invoices"
	"upsider-coding-test/internal/auth"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const HUNDRED = 100

type InvoiceResponse struct {
	ID                string  `json:"id"`
	CompanyID         string  `json:"companyId"`
	BusinessPartnerID string  `json:"businessPartnerId"`
	IssueDate         string  `json:"issueDate"`
	PaymentAmount     int     `json:"paymentAmount"`
	Fee               int     `json:"fee"`
	FeeRate           float64 `json:"feeRate"`
	ConsumptionTax    int     `json:"consumptionTax"`
	TaxRate           float64 `json:"taxRate"`
	TotalAmount       int     `json:"totalAmount"`
	PaymentDueDate    string  `json:"paymentDueDate"`
	Status            string  `json:"status"`
}

type PostController struct {
	companyRepo *model.CompanyRepository
	partnerRepo *model.BusinessPartnerRepository
	invoiceRepo *model.InvoiceRepository
}

func NewPostController(db *gorm.DB) *PostController {
	return &PostController{
		companyRepo: model.NewCompanyRepository(db),
		partnerRepo: model.NewBusinessPartnerRepository(db),
		invoiceRepo: model.NewInvoiceRepository(db),
	}
}

func (c *PostController) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req, err := requestView.NewPostRequest(ctx)
		if err != nil {
			slog.Warn("failed to bind json", "error", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// get company ID from context
		companyUUID, err := getCompanyIDFromContext(ctx)
		if err != nil {
			slog.Warn("failed to get company ID", "error", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// check if partner exists
		if _, err := c.partnerRepo.FindByID(req.PartnerID); err != nil {
			slog.Warn("failed to find partner", "error", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// convert string to uuid
		partnerUUID, err := uuid.Parse(req.PartnerID)
		if err != nil {
			slog.Warn("failed to parse partner ID", "error", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid partner ID format"})
			return
		}

		// calculate Fee
		feeRate := 4
		fee := req.Amount * feeRate / HUNDRED

		// calculate tax
		taxRate := 10
		tax := fee * taxRate / HUNDRED

		// calculate total amount
		totalAmount := req.Amount + fee + tax

		// convert string to time
		paymentDueDate, err := time.Parse("20060102", req.PaymentDue)
		if err != nil {
			slog.Warn("failed to parse payment due date", "error", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// create invoice model
		invoice := model.Invoice{
			CompanyID:         companyUUID,
			BusinessPartnerID: partnerUUID,
			IssueDate:         time.Now(),
			PaymentAmount:     req.Amount,
			Fee:               fee,
			FeeRate:           float64(feeRate) / HUNDRED,
			ConsumptionTax:    tax,
			TaxRate:           float64(taxRate) / HUNDRED,
			TotalAmount:       totalAmount,
			PaymentDueDate:    paymentDueDate,
			Status:            "unpaid",
		}
		if err := c.invoiceRepo.Create(&invoice); err != nil {
			slog.Error("failed to create invoice", "error", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Convert invoice model to response
		ctx.JSON(http.StatusOK, InvoiceResponse{
			ID:                invoice.ID.String(),
			CompanyID:         invoice.CompanyID.String(),
			BusinessPartnerID: invoice.BusinessPartnerID.String(),
			IssueDate:         invoice.IssueDate.Format("2006-01-02"),
			PaymentAmount:     invoice.PaymentAmount,
			Fee:               invoice.Fee,
			FeeRate:           invoice.FeeRate,
			ConsumptionTax:    invoice.ConsumptionTax,
			TaxRate:           invoice.TaxRate,
			TotalAmount:       invoice.TotalAmount,
			PaymentDueDate:    invoice.PaymentDueDate.Format("2006-01-02"),
			Status:            invoice.Status,
		})
	}
}

func getCompanyIDFromContext(ctx *gin.Context) (uuid.UUID, error) {
	cid, exists := ctx.Get(auth.CompanyIDKey)
	if !exists {
		return uuid.Nil, errors.New("company ID not found in context")
	}

	companyID, ok := cid.(uuid.UUID)
	if !ok {
		slog.Error("invalid company ID type",
			"type", fmt.Sprintf("%T", cid),
			"value", fmt.Sprintf("%v", cid))
		return uuid.Nil, fmt.Errorf("invalid company ID type: expected uuid.UUID, got %T", cid)
	}

	return companyID, nil
}
