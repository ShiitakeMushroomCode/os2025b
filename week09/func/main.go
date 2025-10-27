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
	fmt.Print("점수값을 입력하세요(0~100점 사이의 실수) : ")
	score, err := GetFloat()
	if err!=nil{
		log.Fatal(err)
	} 
	if score < 0 || score > 100{
		log.Fatal("잘못된 값이 입력되었습니다. 입력값 :", score)
	}
	status := ""
	if score >= 90{
		status="합격"
	}else{
		status="불합격"
	}
	fmt.Printf("%.2f점 %s입니다.", score, status)
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