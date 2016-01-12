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
	rawString string
	pc        uintptr
	file      string
	line      int
}

func Errorf(format string, a ...interface{}) error {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return fmt.Errorf(format, a...)
	}
	for i := range a {
		switch smart := a[i].(type) {
		case (*smartError):
			a[i] = smart.rawString
		}
	}
	return &smartError{fmt.Sprintf(format, a...), pc, file, line}
}

func (s *smartError) Error() string {
	return fmt.Sprintf("%s%s:%s%d%s %s() %s%s%s",
		ansi.White,
		path.Base(s.file),
		ansi.Yellow,
		s.line,
		ansi.White,
		strings.Split(path.Base(runtime.FuncForPC(s.pc).Name()), ".")[1],
		ansi.Red,
		s.rawString,
		ansi.Reset,
	)
}
