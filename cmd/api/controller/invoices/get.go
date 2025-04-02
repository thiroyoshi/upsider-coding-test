package invoices

import (
	"log/slog"
	"net/http"
	"time"

	"upsider-coding-test/cmd/api/model"
	requestView "upsider-coding-test/cmd/api/view/request/invoices"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GetController struct {
	invoiceRepo *model.InvoiceRepository
}

func NewGetController(db *gorm.DB) *GetController {
	return &GetController{
		invoiceRepo: model.NewInvoiceRepository(db),
	}
}

func (c *GetController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		slog.Info("get controller")

		// Get request
		req, err := requestView.NewGetRequest(ctx)
		if err != nil {
			slog.Warn("failed to bind json", "error", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// convert string to time
		startDate, err := time.Parse("20060102", req.StartAt)
		if err != nil {
			slog.Error("failed to parse start date", "error", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		endDate, err := time.Parse("20060102", req.EndAt)
		if err != nil {
			slog.Error("failed to parse end date", "error", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Get invoices
		invoices, err := c.invoiceRepo.ListByDateRange(startDate, endDate)
		if err != nil {
			slog.Error("failed to get invoices", "error", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Response
		ctx.JSON(http.StatusOK, gin.H{
			"invoices": invoices,
		})
	}
}
