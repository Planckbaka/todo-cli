package errors

import (
	"errors"
)

var (
	ErrTaskNotFound       = errors.New("task not found")
	ErrInvalidID          = errors.New("invalid task ID")
	ErrDatabaseConnection = errors.New("database connection failed")
)
