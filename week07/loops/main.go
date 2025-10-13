package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil{
		log.Fatal(err) //report the error, exit the program
	}

	input = strings.TrimSpace(input)
	score, err := strconv.ParseFloat(input, 64)
	if err != nil{
		log.Fatal(err) //report the error, exit the program
	}
	fmt.Print(score)
	if score > 60{
		fmt.Println(" Good")
	} else{
		fmt.Println(" Fail")
	}
}