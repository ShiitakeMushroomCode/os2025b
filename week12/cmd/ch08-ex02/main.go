package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)


func main() {
	var s1 magazine.Subscriber
	s1.Name = "Mark"
	s1.Rate = 5.6
	fmt.Println(s1)

	var e1 magazine.Employee
	e1.Name = "John"
	e1.Salary = 15000
	e1.Address.City = "Incheon"
	fmt.Println(e1)
}