package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/keyboard"
)

func main() {
	fmt.Printf("실수 입력: ")
	number, err := keyboard.GetFloat()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%.2f", number)
}