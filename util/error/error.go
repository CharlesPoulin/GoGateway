package errors

import (
"google.golang.org/grpc/status"
)// ErrNotFound is for something that not found
func ErrNotFound(model string) error {
	return status.Errorf(404, "the %s resource not found", model)
}// ErrBadRequest is for something that bad request
func ErrBadRequest(msg string) error {
	return status.Errorf(400, msg)
}func DuplicateTransaction() error {
	return status.Errorf(401, "Duplicate Transaction")
}