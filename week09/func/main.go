package main

import (
	"fmt"
)
func main(){
	f:=123
	s:=456
  Swap(f,s)
	fmt.Println("Main Value\nfirst :",f,"second :",s,"\n ")
  Pswap(&f,&s)
	fmt.Println("Main Value\nfirst :",f,"second :",s)
}

func Swap(f int, s int){
	temp:=f
	f=s
	s=temp
	fmt.Println("Value Swap\nfirst :",f,"second :",s)
}

func Pswap(f *int, s *int){
	temp:=*f
	*f=*s
	*s=temp
	fmt.Println("Pointer Swap\nfirst :",*f,"second :",*s)
}

