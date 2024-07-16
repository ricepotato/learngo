package customerr

import (
	"errors"
)

var (
	ErrType1 = errors.New("some type 1")
	ErrType2 = errors.New("some type 2")
)

func ErrorType1() error {
	return ErrType1
}
func ErrorType2() error {
	return ErrType2
}

type CustomError1 struct {
	Message string
	Err     error
}

func (e *CustomError1) Error() string {
	return e.Message
}

type CustomError2 struct {
	Message string
	Err     error
}

func (e *CustomError2) Error() string {
	return e.Message
}

func customError1() *CustomError1 {
	return &CustomError1{
		Message: "Custom Error 1",
		Err:     ErrType1,
	}
}
func customError2() *CustomError2 {
	return &CustomError2{
		Message: "Custom Error 2",
		Err:     ErrType2,
	}
}

func CustomErr1() error {
	return customError1()
}

func CustomErr2() error {
	return customError2()
}
