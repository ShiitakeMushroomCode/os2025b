package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)


func main() {
	weights, err := datafile.GetFloats("meatWeight.txt")
	if err != nil{
		log.Fatal(err)
	}
	sum := 0.0
	for i :=0; i<len(weights);i++{
		sum += weights[i]
	}

	fmt.Printf("평균 고기 소비량 : %.2fkg", sum/float64(len(weights)))
}