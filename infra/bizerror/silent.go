package bizerror

import "github.com/pkg/errors"

var (
	ErrSilent = errors.New("silent")
)

type silentError struct {
	wrapped error
}

func (e *silentError) Error() string {
	return e.wrapped.Error()
}

func (e *silentError) Unwrap() error {
	return e.wrapped
}

func (e *silentError) Is(target error) bool {
	return target == ErrSilent
}

func (e *silentError) Silent() {}

// Mark error as silent, don't report in prometheus metrics
func Silence(err error) error {
	return &silentError{wrapped: err}
}

func IsSilence(err error) bool {
	return errors.Is(err, ErrSilent)
}
