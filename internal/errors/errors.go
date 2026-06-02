package appErrors

import "errors"

var (
	ErrInvalidFieldLength      error = errors.New("wrong lengt of the field")
	ErrInvalidDepartmentNumber error = errors.New("department has same incorrect number")
	ErrInvalidTime             error = errors.New("the specified time is future")
)
