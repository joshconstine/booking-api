package services

import (
	requests "booking-api/data/request"
	responses "booking-api/data/response"
	"booking-api/repositories"

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
			ID:    value.ID,
			Email: value.Email,
		}
		users = append(users, user)
	}
	return users
}

func (t *userServiceImplementation) FindById(id uint) responses.UserResponse {
	result := t.userRepository.FindById(id)

	user := responses.UserResponse{
		ID:    result.ID,
		Email: result.Email,
	}

	return user
}

func (t *userServiceImplementation) FindByEmail(email string) responses.UserResponse {
	result := t.userRepository.FindByEmail(email)

	user := responses.UserResponse{
		ID:    result.ID,
		Email: result.Email,
	}

	return user
}

func (t *userServiceImplementation) CreateUser(user requests.CreateUserRequest) responses.UserResponse {
	// validate request
	err := t.Validate.Struct(user)
	if err != nil {
		panic(err)
	}

	result := t.userRepository.Create(user)

	createdUser := t.FindById(uint(result))

	return createdUser
}
