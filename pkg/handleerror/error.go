package handleerror

import (
	"fmt"

	"github.com/pkg/errors"
)

const (
	NoType = ErrorType(iota)
	BagRequest
	NotFound
)

type ErrorType uint

type customError struct {
	errorType     ErrorType
	originalError error
	contextInfo   errorContext
}

type errorContext struct {
	Field   string
	Message string
}

func (error customError) Error() string {
	return error.originalError.Error()
}

func (et ErrorType) New(msg string) error {
	return customError{errorType: et, originalError: errors.New(msg)}
}

func (et ErrorType) Newf(msg string, args ...interface{}) error {
	err := fmt.Errorf(msg, args...)
	return customError{errorType: et, originalError: err}
}

func (et ErrorType) Wrap(err error, msg string) error {
	return et.Wrapf(err, msg)
}

func (et ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	newErr := errors.Wrapf(err, msg, args...)
	return customError{errorType: et, originalError: newErr}
}

func New(msg string) error {
	return customError{errorType: NoType, originalError: errors.New(msg)}
}

func Newf(msg string, args ...interface{}) error {
	return customError{errorType: NoType, originalError: errors.New(fmt.Sprintf(msg, args...))}
}

func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

func Cause(err error) error {
	return errors.Cause(err)
}

func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedError := errors.Wrapf(err, msg, args...)
	if customErr, ok := err.(customError); ok {
		return customError{
			errorType:     customErr.errorType,
			originalError: wrappedError,
			contextInfo:   customErr.contextInfo,
		}
	}
	return customError{errorType: NoType, originalError: wrappedError}
}

func AddErrorContext(err error, field, message string) error {
	context := errorContext{Field: field, Message: message}
	if customErr, ok := err.(customError); ok {
		return customError{errorType: customErr.errorType, originalError: customErr.originalError, contextInfo: context}
	}
	return customError{errorType: NoType, originalError: err, contextInfo: context}
}

func GetErrorContext(err error) map[string]string {
	emptyError := errorContext{}
	if customErr, ok := err.(customError); ok || customErr.contextInfo != emptyError {
		return map[string]string{"field": customErr.contextInfo.Field, "message": customErr.contextInfo.Message}
	}
	return nil
}

func GetType(err error) ErrorType {
	if customErr, ok := err.(customError); ok {
		return customErr.errorType
	}
	return NoType
}
