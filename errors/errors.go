package errors

import "errors"

func InvalidRequestData() error {
	return errors.New("invalid request data")
}

func EmptyResultData() error {
	return errors.New("empty data")
}
