package main

import (
	"testing"
)

func TestFill(t *testing.T) {
	fillingMap := map[string]interface{}{
		"Field1": "123", "Field2": "321",
	}

	f := Foo{}
	err := mapToFoo(&f, fillingMap)
	if err != nil {
		t.Errorf(err.Error())
	}
	if f.Field1 != "123" || f.Field2 != "321" {
		t.Errorf("Got unexpected values")
	}
}

func TestError1(t *testing.T) {
	s := "str"
	err := mapToFoo(&s, make(map[string]interface{}))
	if err == nil {
		t.Errorf("Got nil error")
	}
}

func TestError2(t *testing.T) {
	err := mapToFoo(1, make(map[string]interface{}))
	if err == nil {
		t.Errorf("Got nil error")
	}
}

func TestError3(t *testing.T) {
	err := mapToFoo(nil, make(map[string]interface{}))
	if err == nil {
		t.Errorf("Got nil error")
	}
}
