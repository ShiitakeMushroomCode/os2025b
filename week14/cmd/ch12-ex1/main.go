package main

import (
	"fmt"
)

func main() {
	//LIFO
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("main")
}