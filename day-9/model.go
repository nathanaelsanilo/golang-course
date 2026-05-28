package main

type User struct {
	ID       int
	Name     string
	IsActive bool
	Weight   float32
}

func NewUser(name string, isActive bool, weight float32) User {
	return User{Name: name, IsActive: isActive, Weight: weight}
}

type UserCreateReqDto struct {
	Name     string  `json:"name"`
	IsActive bool    `json:"is_active"`
	Weight   float32 `json:"weight,omitempty"`
}
