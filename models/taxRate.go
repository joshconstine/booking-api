package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

type TaxRate struct {
	gorm.Model
	Percentage float64
	Name       string
}

func (t *TaxRate) TableName() string {
	return "tax_rates"
}

func (t *TaxRate) MapTaxRateToResponse() response.TaxRateResponse {
	return response.TaxRateResponse{
		ID:         t.ID,
		Percentage: t.Percentage,
		Name:       t.Name,
	}
}
