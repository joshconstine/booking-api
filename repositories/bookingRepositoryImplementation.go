package repositories

import (
	"booking-api/constants"
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
	"booking-api/pkg/database"
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
		Preload("Entities.BookingStatus").
		Preload("Entities.Timeblock").
		Preload("Entities.Documents").
		Preload("Entities.BookingCostItems").
		First(&booking)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return response.BookingInformationResponse{}
	}

	resp := booking.MapBookingToInformationResponse()

	for j, entity := range resp.Entities {
		entityName := GetNameForEntity(entity.EntityID, entity.EntityType, t.Db)
		resp.Entities[j].Name = entityName
		thumbnail := GetThumbnailForEntity(entity.EntityID, entity.EntityType, t.Db)

		resp.Entities[j].Thumbnail = thumbnail
	}
	return resp

}
func calculatePaymentDueDate(bookingStartDate time.Time) time.Time {
	//the due date will be 90 days before the booking start date if the startdate is < 90 days from now the due date is now
	dueDate := bookingStartDate.AddDate(0, 0, -90)
	if dueDate.Before(time.Now()) {
		now := time.Now()
		return now

	}
	return dueDate
}

// TODO: Something can be improvewd here,
func (t *bookingRepositoryImplementation) GetSnapshot(request request.GetBookingSnapshotRequest) []response.BookingSnapshotResponse {
	//limit := 10
	var bookings []models.Booking
	//result := t.Db.Model(&models.Booking{}).Preload("BookingStatus").Preload("Details").Preload("Entities").Limit(limit).Find(&bookings)

	result := t.Db.Scopes(database.Paginate(bookings, &request.PaginationRequest, t.Db)).Preload("BookingStatus").Preload("Details").Preload("Entities").Find(&bookings)
	if result.Error != nil {
		return []response.BookingSnapshotResponse{}
	}

	var response []response.BookingSnapshotResponse
	for _, booking := range bookings {
		response = append(response, booking.MapBookingToSnapshotResponse())
	}

	for i, snapshot := range response {
		var user models.User
		result := t.Db.Model(&models.User{}).Where("id = ?", snapshot.UserID).Select("FirstName", "LastName").First(&user)
		if result.Error != nil {
			fmt.Println(result.Error.Error())
		} else {
			snapshot.Name = user.FirstName + " " + user.LastName
		}

		for j, entity := range snapshot.BookedEntities {
			entityName := GetNameForEntity(entity.EntityID, entity.EntityType, t.Db)
			snapshot.BookedEntities[j].Name = entityName
		}

		// Update the original response slice
		response[i] = snapshot
	}

	return response
}

func GetNameForEntity(entityID uint, entityType string, db *gorm.DB) string {
	var entityResponse response.EntityInfoResponse
	if entityType == constants.BOAT_ENTITY {
		result := db.Model(&models.Boat{}).Where("id = ?", entityID).Select("Name").First(&entityResponse)
		if result.Error != nil {
			fmt.Println(result.Error.Error())
			// return []response.BookingSnapshotResponse{}
		}

		return entityResponse.Name
	} else if entityType == constants.RENTAL_ENTITY {
		result := db.Model(&models.Rental{}).Where("id = ?", entityID).Select("Name").First(&entityResponse)
		if result.Error != nil {
			fmt.Println(result.Error.Error())
			// return []response.BookingSnapshotResponse{}
		}

		return entityResponse.Name

	}
	return ""
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

	if booking.EntityRequests == nil || len(booking.EntityRequests) == 0 {
		booking.EntityRequests = []request.BookEntityRequest{}
	}

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

	//log the booking details
	fmt.Println("Booking Details")
	fmt.Println(bookingToCreate.Details)
	result := t.Db.Model(&models.Booking{}).Create(&bookingToCreate)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		slog.Error(result.Error.Error())
		return "", result.Error
	}

	var earliestStartDate time.Time

	if len(bookingToCreate.Entities) > 0 {
		earliestStartDate = bookingToCreate.Entities[0].Timeblock.StartTime
	} else {
		earliestStartDate = time.Now()

		earliestStartDate = earliestStartDate.AddDate(10, 0, 0)
	}

	for _, entityRequest := range bookingToCreate.Entities {
		if entityRequest.Timeblock.StartTime.Before(earliestStartDate) {
			earliestStartDate = entityRequest.Timeblock.StartTime
		}
	}

	bookingToCreate.Details = models.BookingDetails{
		GuestCount:       booking.Guests,
		BookingStartDate: earliestStartDate,
		PaymentDueDate:   calculatePaymentDueDate(earliestStartDate),
		PaymentComplete:  false,
		DepositPaid:      false,
		DocumentsSigned:  false,
		LocationID:       1,
		BookingID:        bookingToCreate.ID,
	}

	result = t.Db.Model(&models.BookingDetails{}).Where(
		"booking_id = ?", bookingToCreate.ID).Save(&bookingToCreate.Details)
	if result.Error != nil {

		fmt.Println(result.Error.Error())
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
	if request.EntityRequests == nil || len(request.EntityRequests) == 0 {
		return true, nil
	}
	for _, entity := range request.EntityRequests {

		//****************************INSTANT BOOKING CHECK********************************
		allowInstantBooking := t.DoesEntityAllowInstantBooking(entity.EntityID, entity.EntityType)

		// allowInstantBooking = true
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

func (repo *bookingRepositoryImplementation) UpdateBookingStatusForBooking(statusRequest request.UpdateBookingStatusRequest) error {
	result := repo.Db.Model(&models.Booking{}).Where("id = ?", statusRequest.BookingID).Update("booking_status_id", statusRequest.BookingStatusID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
