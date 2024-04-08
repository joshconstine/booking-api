package repositories

import (
	"booking-api/data/response"
	"booking-api/models"

	"gorm.io/gorm"
)

type BoatRepositoryImplementation struct {
	Db *gorm.DB
}

func NewBoatRepositoryImplementation(Db *gorm.DB) BoatRepository {
	return &BoatRepositoryImplementation{Db: Db}
}

func (t *BoatRepositoryImplementation) FindAll() []response.BoatResponse {
	var boats []models.Boat
	result := t.Db.Model(&models.Boat{}).Preload("Timeblocks").Preload("Photos").Preload("Bookings").Preload("BookingCostItems").Preload("BookingDurationRule").Find(&boats)
	if result.Error != nil {
		return []response.BoatResponse{}
	}

	var boatResponses []response.BoatResponse
	for _, boat := range boats {
		boatResponses = append(boatResponses, boat.MapBoatToResponse())
	}

	return boatResponses
}

func (t *BoatRepositoryImplementation) FindById(id int) response.BoatResponse {
	var boat models.Boat
	result := t.Db.Model(&models.Boat{}).Where("id = ?", id).Preload("Timeblocks").Preload("Photos").Preload("Bookings").Preload("BookingCostItems").Preload("BookingDurationRule").First(&boat)
	if result.Error != nil {
		return response.BoatResponse{}
	}

	return boat.MapBoatToResponse()

}
func (t *BoatRepositoryImplementation) Create(boat models.Boat) models.Boat {

	result := t.Db.Create(&boat)
	if result.Error != nil {
		return models.Boat{}
	}

	return boat
}

// // Delete implements TagsRepository
// func (t *BoatRepositoryImplementation) Delete(tagsId int) {
// 	var tags modelsBoat
// 	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
// 	helper.ErrorPanic(result.Error)
// }

// // FindAll implements TagsRepository
// func (t *BoatRepositoryImplementation) FindAll() []models.Boat {
// 	var tags []models.Boat
// 	result := t.Db.Find(&tags)
// 	helper.ErrorPanic(result.Error)
// 	return tags
// }

// // FindById implements TagsRepository
// func (t *BoatRepositoryImplementation) FindById(tagsId int) (tags models.Boat, err error) {
// 	var tag models.Boat
// 	result := t.Db.Find(&tag, tagsId)
// 	if result != nil {
// 		return tag, nil
// 	} else {
// 		return tag, errors.New("tag is not found")
// 	}
// }

// // Save implements TagsRepository
// func (t *BoatRepositoryImplementation) Save(tags models.Boat) {
// 	result := t.Db.Create(&tags)
// 	helper.ErrorPanic(result.Error)
// }

// // Update implements TagsRepository
// func (t *BoatRepositoryImplementation) Update(tags models.Boat) {
// 	var updateTag = request.UpdateTagsRequest{
// 		Id:   tags.Id,
// 		Name: tags.Name,
// 	}
// 	result := t.Db.Model(&tags).Updates(updateTag)
// 	helper.ErrorPanic(result.Error)
// }
