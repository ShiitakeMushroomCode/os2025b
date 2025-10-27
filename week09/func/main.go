package main

import (
	"fmt"
)
func main(){
	amount := 0.0 
	temp := 0.0
	temp = PaintCalc(5.1,6.5)
	amount += temp
	fmt.Printf("%0.2f liters needed\n", temp)
	temp = PaintCalc(6.1,2.5)
	amount += temp
	fmt.Printf("%0.2f liters needed\n", temp)
	fmt.Printf("Total %0.2f liters needed\n", amount)
}
func PaintCalc(w float64, h float64) float64{
	return (w * h / 10)
}