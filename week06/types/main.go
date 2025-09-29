package main

import (
	"fmt"
	"reflect"
)

func main(){
	
	var float_d float64
	var int_d int
	var bool_d bool
	var str_d string

	fmt.Println(float_d, reflect.TypeOf(float_d))
	fmt.Println(int_d, reflect.TypeOf(int_d))
	fmt.Println(bool_d, reflect.TypeOf(bool_d))
	fmt.Println(str_d, reflect.TypeOf(str_d))
}