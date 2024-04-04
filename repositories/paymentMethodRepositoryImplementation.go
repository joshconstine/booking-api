package repositories

import (
	requests "booking-api/data/request"
	"booking-api/data/response"
	responses "booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type PaymentMethodRepositoryImplementation struct {
	Db *gorm.DB
}

func NewPaymentMethodRepositoryImplementation(db *gorm.DB) PaymentMethodRepository {
	return &PaymentMethodRepositoryImplementation{Db: db}
}

func (t *PaymentMethodRepositoryImplementation) FindAll() []response.PaymentMethodResponse {
	var paymentMethods []models.PaymentMethod
	var response []response.PaymentMethodResponse

	result := t.Db.Find(&paymentMethods)
	if result.Error != nil {
		return []responses.PaymentMethodResponse{}
	}

	var item responses.PaymentMethodResponse
	for _, paymentMethod := range paymentMethods {
		item.ID = paymentMethod.ID
		item.Name = paymentMethod.Name

		response = append(response, item)

	}
	return response
}

func (t *PaymentMethodRepositoryImplementation) FindById(id uint) response.PaymentMethodResponse {
	var paymentMethod models.PaymentMethod
	var response response.PaymentMethodResponse

	result := t.Db.First(&paymentMethod, id)
	if result.Error != nil {

		return response
	}

	response.ID = paymentMethod.ID
	response.Name = paymentMethod.Name

	return response
}

func (t *PaymentMethodRepositoryImplementation) Create(paymentMethod requests.CreatePaymentMethodRequest) response.PaymentMethodResponse {
	paymentMethodModel := models.PaymentMethod{
		Name: paymentMethod.Name,
	}
	result := t.Db.Create(&paymentMethodModel)
	if result.Error != nil {
		return response.PaymentMethodResponse{}
	}

	return response.PaymentMethodResponse{
		ID:   paymentMethodModel.ID,
		Name: paymentMethodModel.Name,
	}
}
