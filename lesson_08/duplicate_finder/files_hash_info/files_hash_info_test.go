package files_info_test

import (
	"testing"

	files_info "github.com/i-spirin/geekbrains_2/lesson_08/duplicate_finder/files_hash_info"
)

func TestAdd(t *testing.T) {
	fi := files_info.New()

	err := fi.Add("/tmp/123", "12345")
	if err != nil {
		t.Errorf("Got error for fi.Add")
	}
	if fi.Files["12345"] != "/tmp/123" {
		t.Errorf("Got unexpected result from fi.Files")
	}
	err = fi.Add("/tmp/123", "12345")
	if err == nil {
		t.Errorf("fi.Add does not returned error when expected")
	}
}
