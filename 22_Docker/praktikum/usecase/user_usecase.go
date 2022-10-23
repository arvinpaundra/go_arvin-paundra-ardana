package usecase

import (
	"github.com/arvinpaundra/clean-architecture/dto"
	"github.com/arvinpaundra/clean-architecture/model"
	"github.com/arvinpaundra/clean-architecture/repository"
)

type UserUsecase interface {
	GetAll() ([]model.User, error)
	Create(user dto.UserDTO) error
	Login(userDTO dto.UserDTO) (model.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{
		userRepository: userRepo,
	}
}

func (u *userUsecase) GetAll() ([]model.User, error) {
	users, err := u.userRepository.GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userUsecase) Create(user dto.UserDTO) error {
	newUser := model.User{
		Email:    user.Email,
		Password: user.Password,
	}

	err := u.userRepository.Create(newUser)

	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) Login(userDTO dto.UserDTO) (model.User, error) {
	user, err := u.userRepository.Login(userDTO.Email, userDTO.Password)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
