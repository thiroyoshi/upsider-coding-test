package invoices

import (
	"log/slog"
	"net/http"

	requestView "upsider-coding-test/cmd/api/view/request/invoices"

	"github.com/gin-gonic/gin"
)

type invoice struct {
}

type GetRequest struct {
	invoices []invoice `json:"invoices"`
}

type GetController struct {
}

func NewGetController() *GetController {
	return &GetController{}
}

func (c *GetController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		slog.Info("get controller")

		req, err := requestView.NewGetRequest(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		slog.Info("get request", "req", req)

		// get data form database

		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}
