package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/calendar"
)

func main() {
	event := calendar.Event{}
	err := event.SetTitle("My Birthday!")
	if err != nil{
		log.Fatal(err)
	}
	err = event.SetYear(2025)
	if err != nil{
		log.Fatal(err)
	}
	err = event.SetMonth(11)
	// err = date.SetMonth(19)
	if err != nil{
		log.Fatal(err)
	}
	err = event.SetDay(24)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%d년 %d월 %d일은 %s",event.Year(),event.Month(),event.Day(),event.Title())
}
