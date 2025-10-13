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
	if err != nil{
		log.Fatal(err) //report the error, exit the program
	}
	fmt.Println(input)
}