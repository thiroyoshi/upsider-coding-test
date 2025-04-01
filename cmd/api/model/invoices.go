package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invoice struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	CompanyID         uuid.UUID `gorm:"type:uuid;not null"`
	BusinessPartnerID uuid.UUID `gorm:"type:uuid;not null"`
	IssueDate         time.Time `gorm:"not null"`
	PaymentAmount     int       `gorm:"not null"`
	Fee               int       `gorm:"not null"`
	FeeRate           float64   `gorm:"type:decimal(5,2);not null"`
	ConsumptionTax    int       `gorm:"not null"`
	TaxRate           float64   `gorm:"type:decimal(5,2);not null"`
	TotalAmount       int       `gorm:"not null"`
	PaymentDueDate    time.Time `gorm:"not null"`
	Status            string    `gorm:"type:varchar(20);not null"`
	CreatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type InvoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

func (r *InvoiceRepository) ListByDateRange(start, end time.Time) ([]Invoice, error) {
	var invoices []Invoice
	if err := r.db.Where("issue_date >= ? AND payment_due_date <= ?", start, end).
		Where("status = 'unpaid'").
		Find(&invoices).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}

func (r *InvoiceRepository) Create(invoice *Invoice) error {
	return r.db.Create(invoice).Error
}
