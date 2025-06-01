package kerr

import (
	"errors"
	"fmt"
)

const (
	wrapFormat    = "%w; %w"
	wrapMsgFormat = "%s; %w"
)

// WrapMsg wraps errors with the message. The pattern for errs is FIFO.
// Every next error should be more specific than previous one.
func WrapMsg(msg string, errs ...error) error {
	if len(errs) == 0 {
		return errors.New(msg)
	}

	var resErr error

	for _, err := range errs {
		if err != nil {
			resErr = wrap(resErr, err)
		}
	}

	return fmt.Errorf(wrapMsgFormat, msg, resErr)
}

// WrapErr wraps errors. The pattern for errs is FIFO.
// Every next error should be more specific than previous one.
func WrapErr(errs ...error) error {
	if len(errs) == 0 {
		return nil
	}

	var resErr error

	for _, err := range errs {
		if err != nil {
			resErr = wrap(resErr, err)
		}
	}

	return resErr
}

func wrap(err1, err2 error) error {
	if err1 == nil && err2 == nil {
		return nil
	}

	if err1 == nil {
		return err2
	}

	if err2 == nil {
		return err1
	}

	return fmt.Errorf(wrapFormat, err1, err2)
}
