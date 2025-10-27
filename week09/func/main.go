package main

import (
	"fmt"
	"math"
)
func main(){
	fmt.Print(floatParts(5.1))
}
func floatParts(number float64) (iP int, fP float64){
	wN := math.Floor(number)
	return int(wN), number - wN
}