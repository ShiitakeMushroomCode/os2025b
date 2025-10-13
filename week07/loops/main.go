package main

import (
	"fmt"
	"strings"
	"time"
)

func main(){
	var now time.Time = time.Now()
	var year int = now.Year()
	// month := now.Month()
	var month time.Month = now.Month()
	var day int = now.Day()
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)

	univ := "Go$ Inha$"
	changer := strings.NewReplacer("$", "!")
	fixed := changer.Replace(univ)
	fmt.Println(fixed)
}