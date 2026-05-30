package models

type User struct {
	ID       int
	Name     string
	IsActive bool
	Weight   int32
}

func NewUser(id int, name string, isActive bool, weight int32) User {
	return User{
		ID:       id,
		Name:     name,
		IsActive: isActive,
		Weight:   weight,
	}
}

type UserReqDto struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	Weight   int32  `json:"weight,omitempty"`
}
