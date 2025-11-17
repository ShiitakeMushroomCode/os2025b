package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)


func main() {
	counts := make(map[string]int)
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		counts[line]++
	}
	for name, count := range counts {
		fmt.Println(name, ":", count)
	}
}