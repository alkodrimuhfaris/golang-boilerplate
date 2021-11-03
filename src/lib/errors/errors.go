package errors

import "errors"

var (
	// ErrInternalServerError will throw if internal server is error
	ErrInternalServerError = errors.New("internal server error")
	// ErrNotFound will throw if data is not found
	ErrNotFound = errors.New("error not found")
	// ErrConflict will throw if data is conflict
	ErrConflict = errors.New("error coflict")
	// ErrInvalidTimestamp will throw if timestamp is invalid
	ErrInvalidTimestamp = errors.New("invalid timestamp")
	// ErrAlreadyExist will throw if query is already exist
	ErrAlreadyExist = errors.New("data is already exist")
	// ErrBadRequest will throw if user request is corrupt
	ErrBadRequest = errors.New("invalid request")
)
