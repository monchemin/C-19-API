package errors

import "errors"

func InvalidRequestData() error {
	return errors.New("invalid request data")
}

func EmptyResultData() error {
	return errors.New("empty data")
}

func Unauthorized() error {
	return errors.New("access denied for this resource")
}

func Unknown() error {
	return errors.New("unknown error")
}
