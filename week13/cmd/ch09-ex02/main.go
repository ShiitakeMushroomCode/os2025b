package main

import (
	"fmt"
	"log"
	"week13/pkg/calendar"
)

func main() {
	date := calendar.Date{}
	err := date.SetYear(2025)
	if err != nil{
		log.Fatal(err)
	}
	err = date.SetMonth(11)
	// err = date.SetMonth(19)
	if err != nil{
		log.Fatal(err)
	}
	err = date.SetDay(24)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(date.PrintDate())
}
