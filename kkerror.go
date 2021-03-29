package kkerror

import (
	"fmt"
)

var (
	DefaultErrorCode = "000000"
)

type KKError interface {
	error
	Level() Level
	Category() Category
	Code() string
	Message() string
	Unwrap() error
	WrappedError() KKError
}

type DefaultKKError struct {
	ErrorLevel    Level    `json:"error_level,omitempty"`
	ErrorCategory Category `json:"error_category,omitempty"`
	ErrorCode     string   `json:"error_code,omitempty"` // format as 6 digits
	ErrorMessage  string   `json:"message,omitempty"`    // message
	error         KKError  // wrapped error
}

func (k *DefaultKKError) Level() Level {
	return k.ErrorLevel
}

func (k *DefaultKKError) Category() Category {
	return k.ErrorCategory
}

func (k *DefaultKKError) Code() string {
	return k.ErrorCode
}

func (k *DefaultKKError) Message() string {
	return k.ErrorMessage
}

func (k *DefaultKKError) WrappedError() KKError {
	return k.error
}

func (k *DefaultKKError) Unwrap() error {
	return k.error
}

func Error(message string) *DefaultKKError {
	return &DefaultKKError{
		ErrorLevel:    Normal,
		ErrorCategory: Undefined,
		ErrorCode:     DefaultErrorCode,
		ErrorMessage:  message,
	}
}

func WrappedError(error KKError) *DefaultKKError {
	return &DefaultKKError{
		ErrorLevel:    Normal,
		ErrorCategory: Undefined,
		ErrorCode:     DefaultErrorCode,
		error:         error,
	}
}

func (k *DefaultKKError) Error() string {
	if k.ErrorLevel == "" {
		k.ErrorLevel = Normal
	}

	if k.ErrorCategory == "" {
		k.ErrorCategory = Undefined
	}

	if k.ErrorCode == "" {
		k.ErrorCode = DefaultErrorCode
	}

	return fmt.Sprintf("[%s:%s:%s] %s", k.ErrorLevel, k.ErrorCategory, k.ErrorCode, k.ErrorMessage)
}

func (k *DefaultKKError) String() string {
	return k.Error()
}

func (k *DefaultKKError) PrintStack() {
	fmt.Println(k.Error())
	for i, ik := 1, k.WrappedError(); ik != nil; i, ik = i+1, ik.WrappedError() {
		prefix := ""
		for c := 0; c < i; c++ {
			prefix += "  "
		}

		fmt.Println(fmt.Sprintf("%s- %s", prefix, ik.Error()))
	}
}

func (k *DefaultKKError) StringStack() string {
	var rtn = k.Error()
	for i, ik := 1, k.WrappedError(); ik != nil; i, ik = i+1, ik.WrappedError() {
		prefix := ""
		for c := 0; c < i; c++ {
			prefix += "  "
		}

		rtn += fmt.Sprintf("\n%s- %s", prefix, ik.Error())
	}

	return rtn
}

type Level string

const (
	Critical Level = "critical"
	Urgent   Level = "urgent"
	Normal   Level = "normal"
)

type Category string

const (
	Server    Category = "server"
	Client    Category = "client"
	Database  Category = "database"
	Cache     Category = "cache"
	Process   Category = "process"
	Internal  Category = "internal"
	Undefined Category = "undefined"
)
