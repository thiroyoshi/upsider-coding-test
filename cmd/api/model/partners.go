package model

import (
	"time"

	"gorm.io/gorm"
)

type BusinessPartner struct {
	ID                 string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	CompanyID          string    `gorm:"type:uuid;not null" json:"companyId"`
	CompanyName        string    `gorm:"not null" json:"companyName"`
	RepresentativeName string    `gorm:"not null" json:"representativeName"`
	PhoneNumber        string    `gorm:"not null;size:20" json:"phoneNumber"`
	PostalCode         string    `gorm:"not null;size:8" json:"postalCode"`
	Address            string    `gorm:"not null;type:text" json:"address"`
	CreatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	Company            Company   `gorm:"foreignKey:CompanyID" json:"company"`
}

type BusinessPartnerRepository struct {
	db *gorm.DB
}

func NewBusinessPartnerRepository(db *gorm.DB) *BusinessPartnerRepository {
	return &BusinessPartnerRepository{db: db}
}

func (r *BusinessPartnerRepository) FindByID(id string) (*BusinessPartner, error) {
	var partner BusinessPartner
	if err := r.db.First(&partner, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &partner, nil
}
