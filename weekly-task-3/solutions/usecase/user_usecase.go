package usecase

import (
	"time"

	"github.com/arvinpaundra/golang-blog/dto"
	"github.com/arvinpaundra/golang-blog/model"
	"github.com/arvinpaundra/golang-blog/repository"
	"github.com/google/uuid"
)

type UserUsecase interface {
	Login(email string, password string) (model.User, error)
	Register(user dto.UserDTO) (model.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) Login(email string, password string) (model.User, error) {
	user, err := u.userRepository.Login(email, password)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userUsecase) Register(userDTO dto.UserDTO) (model.User, error) {
	id := uuid.NewString()
	createdAt := time.Now()
	updatedAt := time.Now()

	user := model.User{
		ID:        id,
		Fullname:  userDTO.Fullname,
		Email:     userDTO.Email,
		Password:  userDTO.Password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	err := u.userRepository.Create(user)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
