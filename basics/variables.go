package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 3
	fmt.Println(a)

	b := "1"
	fmt.Println(b)

	fmt.Println(reflect.TypeOf(b))

	var c int
	fmt.Println(c)
}