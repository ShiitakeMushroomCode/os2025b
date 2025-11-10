package main

import "fmt"

func main() {
	subj := [4]string{"Go", "Javascript", "Python", "Linux"}
	subjS := subj[:3]
	subj[0] = "Java"
	subjS = append(subjS, "Go")
	for _, sub := range subj{
		fmt.Println(sub)
	}
	fmt.Println("==================")
	for i:=0; i<len(subjS);i++{
		fmt.Println(subjS[i])	
	}
}