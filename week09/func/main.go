package main

import (
	"fmt"
)
func main(){
	PrintPaintCalc(5.1,6.5)
}
func PrintPaintCalc(w float64, h float64){
	fmt.Printf("%0.2f liters needed\n", (w * h / 10))
}