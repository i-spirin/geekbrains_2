package timederror_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/i-spirin/geekbrains_2/lesson_01/panic_handle/timederror"
)

func TestNew(t *testing.T) {
	err := timederror.New("checkingError")
	if !strings.Contains(err.Error(), "checkingError") {
		t.Fatalf("Wrong error message want substring `checkingError' in %s", err.Error())
	}
}

func ExampleNew() {
	err := timederror.New("checkingError")
	if err != nil {
		fmt.Print(err.Error())
	}
}
