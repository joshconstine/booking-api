package repositories

import (
	"booking-api/constants"
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/models"
	"gorm.io/gorm"
	"log"
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

func (t *userRepositoryImplementation) FindByPublicUserID(userID string) (response.UserResponse, error) {
	var user models.User
	result := t.Db.Where("public_user_id = ?", userID).First(&user)

	if result.Error != nil {
		return response.UserResponse{}, result.Error
	}
	return user.MapUserToResponse(), nil

}
func (t *userRepositoryImplementation) IsAdmin(userID string) bool {
	var memberships []models.Membership

	result := t.Db.Model(&models.Membership{}).Where("user_id = ?", userID).Find(&memberships)

	if result.Error != nil {
		return false
	}

	for _, membership := range memberships {
		if membership.RoleID == constants.USER_ROLE_ACCOUNT_OWNER_ID || membership.RoleID == constants.USER_ROLE_ACCOUNT_MANAGER_ID {
			return true

		}
	}

	return false
}

func (t *userRepositoryImplementation) FindByEmail(email string) (models.User, error) {
	var user models.User
	result := t.Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (t *userRepositoryImplementation) FindByUserID(userID string) response.UserResponse {
	var user models.User
	result := t.Db.Model(models.User{}).Where("id = ?", userID).Preload("Chats.Messages").First(&user)
	if result.Error != nil {
		return response.UserResponse{}
	}
	return user.MapUserToResponse()

}

func (t *userRepositoryImplementation) Create(user *request.CreateUserRequest) error {
	userToInsert := models.User{
		ID:             user.UserID,
		Username:       user.Username,
		Email:          user.Email,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		PhoneNumber:    user.PhoneNumber,
		Gender:         "male",
		DOB:            user.DOB,
		ProfilePicture: "",
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
	result := t.Db.Model(&models.User{}).Where("public_user_id = ?", user.UserID).First(&userToUpdate)

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
