package auth

import (
	"errors"
	"net/http"

	"upsider-coding-test/cmd/api/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ErrMissingAPIKey = errors.New("missing API key")
	ErrInvalidAPIKey = errors.New("invalid API key")
)

const (
	APIKeyHeader = "X-API-Key" //nolint:gosec
	CompanyIDKey = "companyID"
)

func APIKeyAuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	repo := model.NewAPIKeyRepository(db)

	return func(c *gin.Context) {
		apiKey := c.GetHeader(APIKeyHeader)
		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, ErrMissingAPIKey)
			c.Abort()
			return
		}

		// Find API key and associated user
		apiKeyData, err := repo.FindByAPIKey(apiKey)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusUnauthorized, ErrInvalidAPIKey)
				c.Abort()
				return
			}

			c.JSON(http.StatusInternalServerError, err)
			c.Abort()
			return
		}

		// Store company ID in context
		c.Set(CompanyIDKey, apiKeyData.User.CompanyID)
		c.Next()
	}
}
