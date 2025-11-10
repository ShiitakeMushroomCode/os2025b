package main

import "fmt"

func main() {
	// var subj []string
	// subj = make([]string, 3)
	// subj := make([]string, 3)
	// subj[0] = "Go"
	// subj[2] = "Python"
	subj := []string{"Go","","Python"}
	for _, sub := range subj{
		fmt.Println(sub)
	}
}