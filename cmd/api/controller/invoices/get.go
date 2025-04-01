package invoices

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetController struct {
}

func NewGetController() *GetController {
	return &GetController{}
}

func (c *GetController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		slog.Info("get controller")
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}
