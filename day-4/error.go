package main

import "fmt"

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed %s: %s", e.Field, e.Message)
}

func ValidateAge(age int) (bool, error) {
	if age == 0 {
		return false, &ValidationError{Field: "age", Message: "error occurred!"}
	}

	return true, nil
}
