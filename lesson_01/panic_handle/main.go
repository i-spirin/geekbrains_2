package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"strconv"

	"github.com/i-spirin/geekbrains_2/lesson_01/panic_handle/timederror"
)

func main() {
	err := execution()
	if err != nil {
		fmt.Printf("An error has occured: %v\n", err)
	}
	fmt.Println("Correct end of application")
}

func execution() (err error) {
	defer func() {
		if v := recover(); v != nil {
			err = timederror.New(v)
		}
	}()

	f, err := os.Open("./index")
	if err != nil {
		return timederror.New(err)
	}
	defer func() {
		f.Close()
	}()

	content := make([]byte, 1)

	err = binary.Read(f, binary.BigEndian, &content)

	index, _ := strconv.Atoi(string(content))

	fmt.Println(implicitPanic(index))
	return nil
}

func implicitPanic(index int) int {
	someSlice := []int{1, 2, 3, 4, 5}

	return someSlice[index]

}
