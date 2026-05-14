package user

import (
	"fmt"

	// alias
	mdl "github.com/nathanaelsanilo/my-app/model"
)

type UserService interface {
	Create(name string) (mdl.User, error)
	GetById(id int) (mdl.User, error)
	GetUsers() []mdl.User
}

type inMemoryUserService struct {
	users []mdl.User
}

func NewUserService() UserService {
	return &inMemoryUserService{}
}

func (s *inMemoryUserService) GetUsers() []mdl.User {
	return s.users
}

func (s *inMemoryUserService) Create(name string) (mdl.User, error) {
	if name == "" {
		return mdl.User{}, fmt.Errorf("Name is required!")
	}

	n := len(s.users)
	email := fmt.Sprintf("%s@mail.com", name)

	newUsr := mdl.BuildUser(email, name, n+1)

	s.users = append(s.users, newUsr)

	return newUsr, nil
}

func (s *inMemoryUserService) GetById(id int) (mdl.User, error) {
	for _, usr := range s.users {
		if usr.Id == id {
			return usr, nil
		}
	}

	return mdl.User{}, fmt.Errorf("User not found!")
}
