package main

import "fmt"

type UserService interface {
	Create(dto UserCreateReqDto) (User, error)
	GetById(userId int) (User, error)
}

type userService struct {
	store *UserStore
}

func NewUserService(s *UserStore) UserService {
	return &userService{
		store: s,
	}
}

func (s *userService) Create(dto UserCreateReqDto) (User, error) {
	fmt.Printf("service.create() : %v \n", dto)
	newUser := NewUser(dto.Name, dto.IsActive, dto.Weight)
	saved := s.store.Create(newUser)
	fmt.Println("user created!")
	return saved, nil
}

func (s *userService) GetById(userId int) (User, error) {
	fmt.Printf("service.getById() : %d \n", userId)

	saved, err := s.store.FindById(userId)
	if err != nil {
		return User{}, err
	}
	fmt.Println("user returned!")
	return saved, nil
}
