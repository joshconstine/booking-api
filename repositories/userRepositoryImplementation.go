package repositories

import (
	"booking-api/data/request"
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

func (t *userRepositoryImplementation) Create(user request.CreateUserRequest) int {
	result := t.Db.Create(&user)
	if result.Error != nil {
		return 0
	}

	//log the result
	log.Println(result)

	return 1
}

func (t *userRepositoryImplementation) Update(user models.User) models.User {
	result := t.Db.Save(&user)
	if result.Error != nil {
		return models.User{}
	}

	return user
}

func (t *userRepositoryImplementation) Delete(user models.User) models.User {
	result := t.Db.Delete(&user)
	if result.Error != nil {
		return models.User{}
	}

	return user
}
