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
	fmt.Print("섭씨온도를 입력하세요 : ")
	celsius, err := GetFloat()
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("섭씨%.2f도는 화씨%.2f도입니다.", celsius, (celsius * 9/5) + 32)
}

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}