package serror

import (
	"fmt"
	"testing"
)

func TestSimpleError(t *testing.T) {
	fmt.Println(Errorf("simple error"))
}

func TestFormatError(t *testing.T) {
	fmt.Println(Errorf("format %s", "error"))
}

func returnAnError() error {
	err := Errorf("return an error")
	return err
}

func TestConcatSmartError(t *testing.T) {
	fmt.Println(Errorf("Concat smarterror: %v", returnAnError()))
}
