package main

import (
	"fmt"

	mdl "github.com/nathanaelsanilo/my-app/model"
	logger "github.com/nathanaelsanilo/my-app/service/logger"
	"github.com/nathanaelsanilo/my-app/service/user"
)

func main() {
	logger.BuildSugarLogger().Info("Initialized!")

	me := mdl.BuildUser("nathan@mail.com", "nathan", 1)
	fmt.Println("me :", me)

	chocolate := mdl.BuildProduct(1, "chocolate", "snack")
	fmt.Println("chocolate :", chocolate)

	// logger, _ := zap.NewProduction()
	// sugar := logger.Sugar()
	// sugar.Infow("failed to fetch URL", "url", "some-url", "attempt", 3)

	user_service := user.NewUserService()

	user_service.Create("mike")
	user_service.Create("john")
	user_service.Create("alicia")
	user_service.Create("audrey")

	fmt.Printf("users: %p \n", user_service.GetUsers())

	usr, err := user_service.GetById(4)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("User : %+v \n", usr)
}
