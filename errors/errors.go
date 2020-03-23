package errors

import "errors"

func InvalidRequestData() error {
	return errors.New("invalid request data")
}
