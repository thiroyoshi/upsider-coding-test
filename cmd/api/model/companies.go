//nolint:tagalign
package model

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID                 string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	CompanyName        string    `gorm:"not null" json:"companyName"`
	RepresentativeName string    `gorm:"not null" json:"representativeName"`
	PhoneNumber        string    `gorm:"not null;size:20" json:"phoneNumber"`
	PostalCode         string    `gorm:"not null;size:8" json:"postalCode"`
	Address            string    `gorm:"not null;type:text" json:"address"`
	CreatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (r *CompanyRepository) FindByID(id string) (*Company, error) {
	var company Company
	if err := r.db.First(&company, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &company, nil
}
