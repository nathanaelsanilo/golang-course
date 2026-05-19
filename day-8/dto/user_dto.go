package dto

import "time"

type UserCreateDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// struct tags
// tell encoding/json how to map fields
type UserDto struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email,omitempty"` // omitted if zero value
	Password    string    `json:"-"`               // exclude
	DateOfBirth time.Time `json:"date_of_birth"`

	// struct zero value = all fields zero
	// encoding/json never considers a struct empty even if all fields are zero
	// omitempty on struct field does nothing
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
