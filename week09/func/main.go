package main

import (
	"fmt"
)
func main(){
	liters, err := PaintCalc(-5.0, 10.0)
  if err != nil {
    fmt.Println("Error:", err) 
    return
  }
  fmt.Printf("%0.2f liters needed\n", liters)
}
func PaintCalc(w float64, h float64) (float64, error) { 
  if w <= 0 {
    return 0, fmt.Errorf("가로 값이 0이하 입니다. %.2f", w) 
  }
  if h <= 0 {
    return 0, fmt.Errorf("세로 값이 0이하 입니다. %.2f", h) 
  }
  
  liters := w * h / 10
  return liters, nil 
}