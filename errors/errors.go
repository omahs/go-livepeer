package errors

import "fmt"

func NewMaxTranscodeAttemptsError(cause error) error {
	return &MaxTranscodeAttemptsErr{
		cause: cause,
	}
}

type MaxTranscodeAttemptsErr struct {
	cause error
}

func (e *MaxTranscodeAttemptsErr) Error() string {
	return fmt.Sprintf("hit max transcode attempts: %s", e.cause.Error())
}

func (e *MaxTranscodeAttemptsErr) Unwrap() error {
	return e.cause
}
