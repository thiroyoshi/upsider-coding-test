package invoices

import (
	"log/slog"
	"net/http"

	requestView "upsider-coding-test/cmd/api/view/request/invoices"

	"github.com/gin-gonic/gin"
)

type PostController struct {
}

func NewPostController() *PostController {
	return &PostController{}
}

func (c *PostController) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		req, err := requestView.NewPostRequest(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		slog.Info("post request", "req", req)

		slog.Info("post controller")
		ctx.JSON(http.StatusCreated, gin.H{"message": "created"})
	}
}
