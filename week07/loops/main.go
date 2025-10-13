package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main(){
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	fmt.Println(input)
	log.Fatal(err) //report the error, exit the program
}