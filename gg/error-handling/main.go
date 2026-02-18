package main

import (
	"errors"
	"fmt"
	"log"
)

type Person struct{}

func (p *Person) Error(name string) string {
	return fmt.Sprintf(name)
}

//type Error interface {
//	Error() string
//}

type InvalidAgeError struct {
	Age          int
	ErrorMessage string
}

func NewInvalidAgeError(age int, errorMessage string) *InvalidAgeError {
	return &InvalidAgeError{
		Age:          age,
		ErrorMessage: errorMessage,
	}
}

func (e *InvalidAgeError) Error() string {
	// do check a here...
	return fmt.Sprintf("invalid age %d", e.ErrorMessage)
}

func main() {
	invalidError := NewInvalidAgeError(19, "student age is invalid")
	fmt.Println(invalidError.Error())
	res, err := Divide(10, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	//err := fmt.Errorf("some nice error here")

}

func Divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("cannot divide by zero")
	}
	res := a / b
	return res, nil

}
