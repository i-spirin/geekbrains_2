package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	Field1 interface{}
	Field2 interface{}
}

func main() {

	var fillingMap = map[string]interface{}{
		"Field1": "123", "Field2": "321",
	}
	fmt.Println(fillingMap)

	f := Foo{}
	fmt.Println("Before reflect", f)
	mapToFoo(&f, fillingMap)
	fmt.Println("After reflect", f)

}

func mapToFoo(in interface{}, fillingMap map[string]interface{}) error {
	// pointer to struct - addressable
	ps := reflect.ValueOf(in)
	// struct
	if ps.Kind() != reflect.Ptr && ps.Kind() != reflect.Interface {
		return fmt.Errorf("argument in is not pointer or interface")
	}
	s := ps.Elem()
	if s.Kind() != reflect.Struct {
		return fmt.Errorf("in is not struct")
	}
	for key, value := range fillingMap {

		field := s.FieldByName(key)
		if field.IsValid() && field.CanSet() {
			val := reflect.ValueOf(value)
			field.Set(val)
		}
	}
	return nil
}
