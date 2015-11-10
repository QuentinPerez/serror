// Package serror implements functions to manipulate errors.
package serror

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/mgutz/ansi"
)

// smartError is a smart implementation of error.
type smartError struct {
	errorFomatted string
}

func format(str string) string {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return str
	}
	return fmt.Sprintf("%s%s, line %s%d:%s %s()\n\t➜ %s%s%s",
		ansi.White,
		path.Base(file),
		ansi.Yellow,
		line,
		ansi.White,
		strings.Split(path.Base(runtime.FuncForPC(pc).Name()), ".")[1],
		ansi.Red,
		str,
		ansi.Reset,
	)
}

// NewString returns an error using the given string.
func NewString(str string) error {
	return &smartError{format(str)}
}

// NewError returns an error using the given error.
func NewError(err error) error {
	return &smartError{format(err.Error())}
}

func (s *smartError) Error() string {
	return s.errorFomatted
}
