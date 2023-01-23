package errors

import (
	"errors"
	"fmt"
)

func LpError(customErr, cause error) error {
	return &LPError{
		cause:     cause,
		customErr: customErr,
	}
}

type LPError struct {
	cause     error
	customErr error
}

var MaxTranscodeAttempts = errors.New("hit max transcode attempts")

func (e *LPError) Error() string {
	return fmt.Sprintf("%s: %s", e.customErr.Error(), e.cause.Error())
}

func (e *LPError) Is(err error) bool {
	return errors.Is(err, e.cause) || errors.Is(err, e.customErr)
}
