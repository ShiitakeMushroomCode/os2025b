package main

import "fmt"


func main() {
	arrayString := [7]string{"도","레","미","파","솔","라","시"}
	arrayInt := [3]int{-9, 11, 7}
	for i := 0; i<len(arrayInt); i++{
		fmt.Printf("%d번째 arrayInt의 값: %d\n", i+1, arrayInt[i])
	}
	for i := 0; i<len(arrayString); i++{
		fmt.Printf(arrayString[i])
	}
	fmt.Printf("#v\n", arrayInt)
	fmt.Printf("#v\n", arrayString)
}