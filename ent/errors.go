package ent

import "fmt"

func NewValidationError(name string, err error) *ValidationError {
	return &ValidationError{
		Name: name,
		err:  err,
	}
}

func NewNotFoundError(id int) *NotFoundError {
	return &NotFoundError{label: fmt.Sprintf("id=%d", id)}
}
