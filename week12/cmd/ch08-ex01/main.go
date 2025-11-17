package main

import "fmt"

type subcriber struct{
	name string
	price int
}

func applyPrice(s* subcriber){
	s.price = 10000
	s.name = "Park Inha"
}

func main() {
	var s1 subcriber
	s1.name = "Kim inha"
	applyPrice(&s1)
	fmt.Println(s1.name)
	fmt.Println(s1.price)
}