package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
)

type TaxRateRepository interface {
	FindAll() []response.TaxRateResponse
	FindById(id uint) response.TaxRateResponse
	Create(taxRate request.CreateTaxRateRequest) response.TaxRateResponse
}
