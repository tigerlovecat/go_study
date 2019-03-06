package errno

import "fmt"

type Errno struct {
	Status int
	Info   string
}

func (err Errno) Error() string {
	return err.Info
}

// Err represents an error
type Err struct {
	Status int
	Info   string
	Err    error
}

func (err *Err)New(newErr error) *Err {
	return &Err{Status: err.Status, Info: err.Info, Err: newErr}
}

func (err *Err) Add(message string) error {
	err.Info += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Info += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Status, err.Info, err.Err)
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Status
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Status, OK.Info
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Status, typed.Info
	case *Errno:
		return typed.Status, typed.Info
	default:
	}

	return InternalServerError.Status, err.Error()
}
