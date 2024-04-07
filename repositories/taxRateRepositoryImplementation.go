package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type TaxRateRepositoryImplementation struct {
	Db *gorm.DB
}

func NewTaxRateRepositoryImplementation(db *gorm.DB) TaxRateRepository {
	return &TaxRateRepositoryImplementation{Db: db}
}

func (r *TaxRateRepositoryImplementation) FindAll() []response.TaxRateResponse {
	var taxRates []models.TaxRate
	result := r.Db.Find(&taxRates)
	if result.Error != nil {
		return []response.TaxRateResponse{}
	}

	var taxRateResponses []response.TaxRateResponse
	for _, taxRate := range taxRates {
		taxRateResponses = append(taxRateResponses, response.TaxRateResponse{
			ID:         taxRate.ID,
			Percentage: taxRate.Percentage,
			Name:       taxRate.Name,
		})
	}

	return taxRateResponses
}

func (r *TaxRateRepositoryImplementation) FindById(id uint) response.TaxRateResponse {
	var taxRate models.TaxRate
	result := r.Db.Where("id = ?", id).First(&taxRate)
	if result.Error != nil {
		return response.TaxRateResponse{}
	}

	return response.TaxRateResponse{
		ID:         taxRate.ID,
		Percentage: taxRate.Percentage,
		Name:       taxRate.Name,
	}
}

func (r *TaxRateRepositoryImplementation) Create(taxRate request.CreateTaxRateRequest) response.TaxRateResponse {

	taxRateModel := models.TaxRate{
		Percentage: taxRate.Percentage,
		Name:       taxRate.Name,
	}

	result := r.Db.Create(&taxRateModel)
	if result.Error != nil {
		return response.TaxRateResponse{}
	}

	return response.TaxRateResponse{
		ID:         taxRateModel.ID,
		Percentage: taxRateModel.Percentage,
		Name:       taxRateModel.Name,
	}
}
