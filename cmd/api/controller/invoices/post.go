package invoices

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct {
}

func NewPostController() *PostController {
	return &PostController{}
}

func (c *PostController) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		slog.Info("post controller")
		ctx.JSON(http.StatusCreated, gin.H{"message": "created"})
	}
}
