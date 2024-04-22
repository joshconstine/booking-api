package repositories

import (
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
	"log"

	"gorm.io/gorm"
)

type userRepositoryImplementation struct {
	Db *gorm.DB
}

func NewUserRepositoryImplementation(Db *gorm.DB) UserRepository {
	return &userRepositoryImplementation{Db: Db}
}

func (t *userRepositoryImplementation) FindAll() []models.User {
	var users []models.User
	result := t.Db.Find(&users)
	if result.Error != nil {
		return []models.User{}
	}

	return users
}

func (t *userRepositoryImplementation) FindById(id uint) models.User {
	var user models.User
	result := t.Db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return models.User{}
	}

	return user
}

func (t *userRepositoryImplementation) FindByEmail(email string) models.User {
	var user models.User
	result := t.Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return models.User{}
	}

	return user
}

func (t *userRepositoryImplementation) FindByUserID(userID string) response.UserResponse {
	var user models.User
	result := t.Db.Where("user_id = ?", userID).First(&user)
	if result.Error != nil {
		return response.UserResponse{}
	}

	result = t.Db.Model(&models.Chat{}).Where("user_id = ?", user.UserID).Find(&user.Chats)
	if result.Error != nil {
		return response.UserResponse{}
	}

	return user.MapUserToResponse()

}

func (t *userRepositoryImplementation) Create(user *request.CreateUserRequest) error {
	userToInsert := models.User{
		UserID:      user.UserID,
		Username:    user.Username,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
	}

	result := t.Db.Model(&models.User{}).Create(&userToInsert)
	if result.Error != nil {
		return result.Error
	}

	//log the result
	log.Println(result)

	return nil
}

func (t *userRepositoryImplementation) Update(user *request.UpdateUserRequest) error {
	var userToUpdate models.User
	result := t.Db.Model(&models.User{}).Where("user_id = ?", user.UserID).First(&userToUpdate)

	if result.Error != nil {
		return result.Error
	}

	userToUpdate.Username = user.Username
	userToUpdate.FirstName = user.FirstName
	userToUpdate.LastName = user.LastName
	userToUpdate.PhoneNumber = user.PhoneNumber

	result = t.Db.Save(&userToUpdate)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t *userRepositoryImplementation) Delete(user models.User) models.User {
	result := t.Db.Delete(&user)
	if result.Error != nil {
		return models.User{}
	}

	return user
}
