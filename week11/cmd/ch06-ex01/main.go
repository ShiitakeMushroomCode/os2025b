package main

import "fmt"

func main() {
	subj := []string{"Go", "Javascript", "Python", "Linux"}
	subjS := subj[1:3]
	for _, sub := range subj{
		fmt.Println(sub)
	}
	fmt.Println("==================")
	for i:=0; i<len(subjS);i++{
		fmt.Println(subjS[i])	
	}
}