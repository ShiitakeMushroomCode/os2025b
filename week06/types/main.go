package main

import (
	"fmt"
	"reflect"
)

func main(){
	// var name string
	// var id int
	// name = "Kim"
	// id = 1000

	// var name string = "Kim"
	// var id int = 1000

	// var name = "Kim"
	// var id = 1000

	name := "Kim"
	id := 1000

	fmt.Println(name,reflect.TypeOf(name))
	fmt.Println(id, reflect.TypeOf(id))
}