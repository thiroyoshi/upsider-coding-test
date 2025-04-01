package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type APIKey struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	APIKey    string    `gorm:"type:varchar(64);uniqueIndex;not null"`
	Name      string    `gorm:"type:varchar(255);not null"`
	ExpiresAt *time.Time
	IsActive  bool `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User `gorm:"foreignKey:UserID"`
}

type APIKeyRepository struct {
	db *gorm.DB
}

func NewAPIKeyRepository(db *gorm.DB) *APIKeyRepository {
	return &APIKeyRepository{db: db}
}

// FindByAPIKey retrieves APIKey and associated User data by the API key string
func (r *APIKeyRepository) FindByAPIKey(apiKey string) (*APIKey, error) {
	var result APIKey
	if err := r.db.Preload("User").
		Where("api_key = ? AND is_active = true", apiKey).
		Where("expires_at IS NULL OR expires_at > ?", time.Now()).
		First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
