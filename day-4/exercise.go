package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float32
	Height float32
}

type ShapeError struct {
	Field   string
	Message string
}

func (s *ShapeError) Error() string {
	return fmt.Sprintf("Validation failed : %s - %s", s.Field, s.Message)
}

func (c Circle) PrintShapeInfo(s Shape) string {
	area := s.Area()
	perimeter := s.Perimeter()
	return fmt.Sprintf("Info area %f, perimeter %f", area, perimeter)
}

func (r Rectangle) PrintShapeInfo(s Shape) string {
	return ""
}

func (c Circle) Area() float64 {
	return (c.Radius * 2) * math.Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func BuildCircle(radius float64) (Circle, error) {
	if radius <= 0 {
		return Circle{}, &ShapeError{Field: "radius", Message: "radius is not valid"}
	}

	return Circle{Radius: radius}, nil
}
