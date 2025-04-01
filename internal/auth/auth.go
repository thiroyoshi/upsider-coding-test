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
	APIKeyHeader = "X-API-Key"
	CompanyIDKey = "companyID"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	repo := model.NewAPIKeyRepository(db)

	return func(c *gin.Context) {
		apiKey := c.GetHeader(APIKeyHeader)
		if apiKey == "" {
			c.AbortWithError(http.StatusUnauthorized, ErrMissingAPIKey)
			return
		}

		// Find API key and associated user
		apiKeyData, err := repo.FindByAPIKey(apiKey)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithError(http.StatusUnauthorized, ErrInvalidAPIKey)
				return
			}
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// Store company ID in context
		c.Set(CompanyIDKey, apiKeyData.User.CompanyID)
		c.Next()
	}
}
