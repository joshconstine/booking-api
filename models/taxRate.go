package models

import (
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
