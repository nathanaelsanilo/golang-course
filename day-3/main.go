package main

import "fmt"

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string
	Age    int
	PointX float32 `json:"point_x"`
	PointY float32 `json:"point_y"` // struct tags. similar with @JsonProperty in jakarta package. GO use reflect (in java reflection, the ability to see itself)
}

type Address struct {
	City    string
	Country string
}

// GO doesnt have inheritance
type Employee struct {
	User
	Address    // embedding
	Department string
}

func main() {

	u := User{
		Id:     1, // public fields always use Uppercase. private fields use lowercase
		Name:   "nathan",
		Email:  "nathan@mail.com",
		Age:    31,
		PointX: 0,
		PointY: 0,
	}
	fmt.Println("User :", u)
	fmt.Println(u.Greet())
	u.Walk()
	fmt.Println("User :", u)

	u2 := BuildUser()
	u2.Name = "He who remain"
	u2.Walk()
	fmt.Println(u2)

	empl := BuildEmployee()
	fmt.Println("empl :", empl)

	laptop := BuildProduct()
	laptop.ID = 1
	laptop.Name = "ASUS Laptop"
	laptop.Price = 3000
	laptop.Stock = 10
	laptop.ApplyDiscount(0.1)
	fmt.Println("laptop :", laptop)

	mouse := BuildProduct()
	mouse.ID = 2
	mouse.Name = "Razer Mouse"
	mouse.Price = 120
	mouse.Stock = 5
	fmt.Println("mouse :", mouse)

}

// factory function replace constructor
func BuildUser() User {
	return User{
		Id:     0,
		Name:   "",
		Email:  "",
		Age:    0,
		PointX: 0,
		PointY: 0,
	}
}

func BuildAddress() Address {
	return Address{
		City:    "",
		Country: "",
	}
}

func BuildEmployee() Employee {
	u := BuildUser()
	u.Age = 31
	u.Id = 1
	u.Name = "nathan"
	u.Email = "nathan@mail.com"

	address := BuildAddress()
	address.City = "Jakarta"
	address.Country = "ID"

	return Employee{
		User:       u,
		Address:    address,
		Department: "IT",
	}
}

// value receiver : copy the struct. cannot modify
func (u User) Greet() string {
	return "Hello " + u.Name
}

// pointer receiver : modify class member. u is the actual struct
func (u *User) Walk() {
	u.PointX = 10.5
	u.PointY = 22.5
}
