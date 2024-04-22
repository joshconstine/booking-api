package repositories

import (
	"booking-api/constants"
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
	"fmt"
	"log/slog"
	"strconv"
	"time"

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
		// Preload("User").
		Preload("CostItems.BookingCostType").
		Preload("CostItems.TaxRate").
		Preload("Documents.Document").
		Preload("BookingStatus").
		Preload("Entities").
		First(&booking)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return response.BookingInformationResponse{}
	}

	return booking.MapBookingToInformationResponse()

}

func CalculateNightlyCost(amount float64, startTime time.Time, endTime time.Time) float64 {
	// Ensure endTime is after startTime
	if endTime.Before(startTime) {
		return 0
	}

	// Calculate the number of whole nights between startTime and endTime
	duration := endTime.Sub(startTime)
	nights := int(duration.Hours() / 24)

	return float64(nights) * amount
}

func CalculateCostAmmount(bookingCostTypeID uint, amount float64, startTime time.Time, endTime time.Time) float64 {

	switch bookingCostTypeID {
	case constants.BOOKING_COST_TYPE_BOAT_RENTAL_COST_ID, constants.BOOKING_COST_TYPE_RENTAL_COST_ID:
		return CalculateNightlyCost(amount, startTime, endTime)
	default:
		return amount
	}

}

func CalculateCostsForEntityBooking(entity models.EntityBooking, db *gorm.DB) []models.BookingCostItem {
	var costsForEntity []models.BookingCostItem
	var entityCosts []models.EntityBookingCost

	result := db.Model(&models.EntityBookingCost{}).Where(
		"entity_id = ? AND entity_type = ?",
		entity.EntityID,
		entity.EntityType,
	).Find(&entityCosts)

	if result.Error != nil {
		return []models.BookingCostItem{}
	}

	for _, entityCost := range entityCosts {
		costsForEntity = append(costsForEntity, models.BookingCostItem{
			BookingCostTypeID: entityCost.BookingCostTypeID,
			Amount:            CalculateCostAmmount(entityCost.BookingCostTypeID, entityCost.Amount, entity.Timeblock.StartTime, entity.Timeblock.EndTime),
			TaxRateID:         entityCost.TaxRateID,
			EntityBookingID:   entity.ID,
		})
	}

	return costsForEntity
}

func GetDocumentsForEntity(entity models.EntityBooking, db *gorm.DB) []models.BookingDocument {
	var documentsForEntity []models.BookingDocument
	var entityDocuments []models.EntityBookingDocument

	result := db.Model(&models.EntityBookingDocument{}).Where(
		"entity_id = ? AND entity_type = ?",
		entity.EntityID,
		entity.EntityType,
	).Find(&entityDocuments)

	if result.Error != nil {
		return []models.BookingDocument{}
	}

	for _, entityDocument := range entityDocuments {
		documentsForEntity = append(documentsForEntity, entityDocument.MapEntityBookingDocumentToBookingDocument())
	}

	return documentsForEntity
}

func (t *bookingRepositoryImplementation) Create(booking *request.CreateBookingRequest) (string, error) {

	bookingToCreate := booking.MapCreateBookingRequestToBooking()

	//*************COSTS FOR ENTITIES****************
	var costsForBooking []models.BookingCostItem
	for _, entityBooking := range bookingToCreate.Entities {

		costsForBooking = append(costsForBooking, CalculateCostsForEntityBooking(entityBooking, t.Db)...)
	}

	bookingToCreate.CostItems = costsForBooking

	//************DOCUMENTS FOR ENTITIES ************

	var documentsForBooking []models.BookingDocument
	for _, entityBooking := range bookingToCreate.Entities {
		documentsForBooking = append(documentsForBooking, GetDocumentsForEntity(entityBooking, t.Db)...)
	}

	bookingToCreate.Documents = documentsForBooking

	result := t.Db.Model(&models.Booking{}).Create(&bookingToCreate)
	if result.Error != nil {
		slog.Error(result.Error.Error())
		return "", result.Error
	}

	return bookingToCreate.ID, nil
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

		fmt.Print(allowInstantBooking)
		if !allowInstantBooking {
			// check if there is an inquiry for the entity
			// if there is an inquiry, check if the entitybookingrequest was approved
			// if it was approved, the request.entitybookingrequest start and end times must be within the approved time

			var bookingPermissionRequests []models.EntityBookingPermission

			var user models.User

			result := t.Db.Model(&models.User{}).Where("email = ?", request.Email).First(&user)
			if result.Error != nil {
				return false, result.Error
			}

			result = t.Db.Model(&models.EntityBookingPermission{}).Where("user_id = ? AND entity_id = ? AND entity_type = ?", user.ID, entity.EntityID, entity.EntityType).Find(&bookingPermissionRequests)
			if result.Error != nil {
				return false, result.Error
			}

			if request.InquiryID == 0 {
				errorMsg := fmt.Sprintf("entity (ID: %s , Type: %s) requires an inquiry to be made", strconv.Itoa(int(entity.EntityID)), entity.EntityType)
				return false, fmt.Errorf(errorMsg)
			}

			for _, entityInquiry := range bookingPermissionRequests {
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
