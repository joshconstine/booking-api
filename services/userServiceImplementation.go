package services

import (
	requests "booking-api/data/request"
	responses "booking-api/data/response"
	"booking-api/repositories"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type userServiceImplementation struct {
	userRepository repositories.UserRepository
	Validate       *validator.Validate
}

func NewUserServiceImplementation(userRepository repositories.UserRepository, validate *validator.Validate) UserService {
	return &userServiceImplementation{
		userRepository: userRepository,
		Validate:       validate,
	}
}
func (t *userServiceImplementation) FindAll() []responses.UserResponse {
	result := t.userRepository.FindAll()

	var users []responses.UserResponse
	for _, value := range result {
		user := responses.UserResponse{
			UserID: value.ID,
			Email:  value.Email,
		}
		users = append(users, user)
	}
	return users
}

func (t *userServiceImplementation) FindByUserID(userID string) responses.UserResponse {

	result := t.userRepository.FindByUserID(userID)

	return result
}

func (t *userServiceImplementation) FindByPublicUserID(publicUserID string) (responses.UserResponse, error) {
	result, err := t.userRepository.FindByPublicUserID(publicUserID)

	if err != nil {
		return responses.UserResponse{}, err

	}
	return result, nil

}
func (t *userServiceImplementation) IsAdmin(userID string) bool {
	result := t.userRepository.IsAdmin(userID)

	return result
}

func (t *userServiceImplementation) FindById(id uint) responses.UserResponse {
	result := t.userRepository.FindById(id)

	user := responses.UserResponse{
		UserID: result.ID,
		Email:  result.Email,
	}

	return user
}

func (t *userServiceImplementation) FindByEmail(email string) (responses.UserResponse, error) {
	result, err := t.userRepository.FindByEmail(email)
	if err != nil {
		return responses.UserResponse{}, err
	}

	return result.MapUserToResponse(), nil
}

func (t *userServiceImplementation) FindByEmailPublic(email string) (responses.PublicUserResponse, error) {
	result, err := t.userRepository.FindByEmail(email)
	if err != nil {
		return responses.PublicUserResponse{}, err
	}

	return result.MapUserToPublicResponse(), nil
}
func (t *userServiceImplementation) CreateUser(user *requests.CreateUserRequest) error {
	// validate request
	fmt.Println("validating user")
	// err := t.Validate.Struct(user)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return err
	// }

	fmt.Println("after user")
	result := t.userRepository.Create(user)

	return result
}

func (t *userServiceImplementation) UpdateUser(user *requests.UpdateUserRequest) error {
	// validate request
	err := t.Validate.Struct(user)
	if err != nil {
		return err
	}

	result := t.userRepository.Update(user)

	return result
}
