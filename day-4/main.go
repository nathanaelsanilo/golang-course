package main

import (
	"errors"
	"fmt"
)

type Describer interface {
	Describe() string
}

type User struct {
	Name string
}

type Product struct {
	Name  string
	Price float32
}

// declare a method with the same name as interface
// then use value receiver
func (u User) Describe() string {
	return "User : " + u.Name
}

func (p Product) Describe() string {
	return "Product : " + p.Name
}

// polymorphism
func PrintInfo(payload Describer) {
	fmt.Println("Describe payload :", payload.Describe())
}

func ThrowError() error {
	return errors.New("Example of throw an error!")
}

func main() {
	// PrintInfo(User{Name: "nathan"})
	// PrintInfo(Product{Name: "nathan", Price: 35_000_000.25})

	// fmt.Println("Error :", ThrowError())

	val, err := ValidateAge(0)
	if err != nil {
		fmt.Println(err)
		var valErr *ValidationError
		errors.As(err, &valErr)

		// newer version of unpacking error
		valErr2, ok := errors.AsType[*ValidationError](err)
		if ok {
			fmt.Println("valErr2: ", valErr2)
		}
	} else {
		fmt.Println(val)
	}

	circle1, errCircle := BuildCircle(0)
	if errCircle != nil {
		fmt.Println("err :", errCircle)
	} else {
		fmt.Println("circle :", circle1)
	}

}
