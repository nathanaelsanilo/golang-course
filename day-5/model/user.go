package model

import service "github.com/nathanaelsanilo/my-app/service/logger"

type User struct {
	Name  string
	Id    int
	email string
}

// exposed
func BuildUser(email, name string, id int) User {
	return User{
		Name:  name,
		Id:    id,
		email: email,
	}
}

// private : accessible within this package
func validateEmail(email string) bool {
	return len(email) > 0
}

func init() {
	service.BuildSugarLogger().Info("Initialized!")
}
