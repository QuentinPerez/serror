package serror

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewString(t *testing.T) {
	fmt.Println(NewString("foo"))
}

func TestNewError(t *testing.T) {
	fmt.Println(NewError(errors.New("bar")))
}
