package models

import (
	"booking-api/data/response"

	"gorm.io/gorm"
)

//tax rate has a unique name+percentage constraint

type TaxRate struct {
	gorm.Model
	Percentage float64 `gorm:"index:tax_percent,unique;"`
	Name       string  `gorm:"index:tax_percent,unique;"`
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
