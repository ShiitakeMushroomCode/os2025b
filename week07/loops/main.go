package main

import (
	"fmt"
	"reflect"
)

func main(){
	var length float64 = 3.2;
	var width int = 2;
	fmt.Println("Area is", int(length) *(width));
	fmt.Println("length > width", int(length) > (width));
	fmt.Println("형변환", reflect.TypeOf(int(length)));
	fmt.Println("원본", reflect.TypeOf(length));
}