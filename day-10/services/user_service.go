package services

import (
	"fmt"

	"github.com/nathanaelsanilo/my-app/models"
	"github.com/nathanaelsanilo/my-app/stores"
)

type UserService interface {
	GetAll() []models.User
	GetDetail(userId int) (models.User, error)
	DeleteById(userId int) error
	Update(userId int, dto models.UserReqDto) (models.User, error)
	Create(dto models.UserReqDto) models.User
}

type userService struct {
	store stores.UserStore
}

func NewUserService(store stores.UserStore) UserService {
	return &userService{
		store: store,
	}
}

func (s *userService) GetAll() []models.User {
	return s.store.FindAll()
}

func (s *userService) GetDetail(userId int) (models.User, error) {
	u, err := s.store.FindById(userId)

	if err != nil {
		return models.User{}, fmt.Errorf("Error:UserService.GetDetail:%w", err)
	}

	return u, nil
}

func (s *userService) DeleteById(userId int) error {
	err := s.store.DeleteById(userId)

	if err != nil {
		return fmt.Errorf("Error:UserService.DeleteById:%w", err)
	}

	return nil
}

func (s *userService) Update(userId int, dto models.UserReqDto) (models.User, error) {
	u := models.NewUser(
		userId,
		dto.Name,
		dto.IsActive,
		dto.Weight,
	)

	saved, err := s.store.Update(userId, u)

	if err != nil {
		return models.User{}, fmt.Errorf("Error:UserService.Update:%w", err)
	}

	return saved, nil
}

func (s *userService) Create(dto models.UserReqDto) models.User {
	u := models.User{
		Name:     dto.Name,
		IsActive: dto.IsActive,
		Weight:   dto.Weight,
	}

	saved := s.store.Create(u)

	return saved
}
