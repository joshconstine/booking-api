package repositories

import (
	"booking-api/constants"
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
	"fmt"
	"log/slog"
	"strconv"

	"gorm.io/gorm"
)

type bookingRepositoryImplementation struct {
	Db *gorm.DB
}

func NewBookingRepositoryImplementation(Db *gorm.DB) BookingRepository {
	return &bookingRepositoryImplementation{Db: Db}
}

func (t *bookingRepositoryImplementation) FindAll() []response.BookingResponse {
	var bookings []models.Booking
	result := t.Db.Find(&bookings)
	if result.Error != nil {
		return []response.BookingResponse{}
	}

	var response []response.BookingResponse
	for _, booking := range bookings {
		response = append(response, booking.MapBookingToResponse())
	}

	return response

}

func (t *bookingRepositoryImplementation) FindById(id string) response.BookingInformationResponse {
	var booking models.Booking
	result := t.Db.Model(&models.Booking{}).Where(
		"id = ?", id).
		Preload("Payments.PaymentMethod").
		Preload("Details").
		Preload("CostItems").
		Preload("User").
		Preload("CostItems.BookingCostType").
		Preload("CostItems.TaxRate").
		Preload("Documents.Document").
		Preload("BookingStatus").
		Preload("Entities").
		First(&booking)

	if result.Error != nil {
		return response.BookingInformationResponse{}
	}

	return booking.MapBookingToInformationResponse()

}

func (t *bookingRepositoryImplementation) Create(booking *request.CreateBookingRequest) error {

	bookingToCreate := booking.MapCreateBookingRequestToBooking()
	result := t.Db.Model(&models.Booking{}).Create(&bookingToCreate)
	if result.Error != nil {
		slog.Error(result.Error.Error())
		return result.Error
	}

	return nil
}

func (t *bookingRepositoryImplementation) Update(booking models.Booking) models.Booking {
	result := t.Db.Save(&booking)
	if result.Error != nil {
		return models.Booking{}
	}

	return booking
}
func (t *bookingRepositoryImplementation) DoesEntityAllowInstantBooking(entityID uint, entityType string) bool {
	var entityRules []models.EntityBookingRule
	var allowInstantBooking bool
	result := t.Db.Model(&models.EntityBookingRule{}).Where(
		"entity_id = ? AND entity_type = ?",
		entityID,
		entityType,
	).First(&entityRules)

	if result.Error != nil {
		return false
	}

	allowInstantBooking = entityRules[0].AllowInstantBooking

	fmt.Println(allowInstantBooking)
	if result.Error != nil {
		return false
	}

	return allowInstantBooking
}

func (t *bookingRepositoryImplementation) CheckIfEntitiesCanBeBooked(request *request.CreateBookingRequest) (bool, error) {
	for _, entity := range request.EntityRequests {

		//****************************INSTANT BOOKING CHECK********************************
		allowInstantBooking := t.DoesEntityAllowInstantBooking(entity.EntityID, entity.EntityType)

		if !allowInstantBooking {
			//check if there is an inquiry for the entity
			//if there is an inquiry, check if the entitybookingrequest was approved
			//if it was approved, the request.entitybookingrequest start and end times must be within the approved time
			var inquery models.Inquiry

			if request.InquiryID == 0 {
				errorMsg := fmt.Sprintf("entity (ID: %s , Type: %s) requires an inquiry to be made", strconv.Itoa(int(entity.EntityID)), entity.EntityType)
				return false, fmt.Errorf(errorMsg)
			}

			result := t.Db.Model(&models.Inquiry{}).Where("id = ?", request.InquiryID).First(&inquery)
			if result.Error != nil {
				return false, result.Error
			}

			for _, entityInquiry := range inquery.BookingRequests {
				if entityInquiry.EntityID == entity.EntityID && entityInquiry.EntityType == entity.EntityType {
					//check if the entitybookingrequest was approved
					if entityInquiry.InquiryStatusID == constants.INQUIRY_STATUS_APPROVED_ID {
						continue
					} else if entityInquiry.InquiryStatusID == constants.INQUIRY_STATUS_APPROVAL_EXPIRED_ID {
						errorMsg := fmt.Sprintf("entity (ID: %s , Type: %s) inquiry approval has expired", strconv.Itoa(int(entity.EntityID)), entity.EntityType)
						return false, fmt.Errorf(errorMsg)
					} else if entityInquiry.InquiryStatusID == constants.INQUIRY_STATUS_CANCELLED_ID {
						errorMsg := fmt.Sprintf("entity (ID: %s , Type: %s) inquiry has been cancelled", strconv.Itoa(int(entity.EntityID)), entity.EntityType)
						return false, fmt.Errorf(errorMsg)
					} else if entityInquiry.InquiryStatusID == constants.INQUIRY_STATUS_DECLINED_ID {
						errorMsg := fmt.Sprintf("entity (ID: %s , Type: %s) inquiry has been declined", strconv.Itoa(int(entity.EntityID)), entity.EntityType)
						return false, fmt.Errorf(errorMsg)
					} else if entityInquiry.InquiryStatusID == constants.INQUIRY_STATUS_NEW_ID {
						errorMsg := fmt.Sprintf("entity (ID: %s , Type: %s) inquiry has not been approved", strconv.Itoa(int(entity.EntityID)), entity.EntityType)
						return false, fmt.Errorf(errorMsg)
					} else {
						errorMsg := fmt.Sprintf("entity (ID: %s , Type: %s) inquiry has an unknown status", strconv.Itoa(int(entity.EntityID)), entity.EntityType)
						return false, fmt.Errorf(errorMsg)
					}

				}
			}
		}

		//****************************TIMEBLOCK OVERLAP CHECK********************************
		var entityTimeblocks []models.EntityTimeblock
		result := t.Db.Model(&models.EntityTimeblock{}).Where(
			"entity_id = ? AND entity_type = ?",
			entity.EntityID,
			entity.EntityType,
		).Find(&entityTimeblocks)

		if result.Error != nil {
			return false, result.Error
		}

		for _, timeblock := range entityTimeblocks {
			if (entity.StartTime.Equal(timeblock.StartTime) || entity.StartTime.After(timeblock.StartTime)) && (entity.StartTime.Before(timeblock.EndTime)) ||
				((entity.EndTime.Equal(timeblock.EndTime) || entity.EndTime.Before(timeblock.EndTime)) && (entity.EndTime.After(timeblock.StartTime))) ||
				(entity.StartTime.Before(timeblock.StartTime) && entity.EndTime.After(timeblock.EndTime)) {
				errorMsg := fmt.Sprintf("entity (ID: %s , Type: %s) is already booked for the time range %s to %s", strconv.Itoa(int(entity.EntityID)), entity.EntityType, entity.StartTime, entity.EndTime)
				return false, fmt.Errorf(errorMsg)
			}
		}
	}

	return true, nil
}
