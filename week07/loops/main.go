package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main(){
	for i := 0; i < 10; i++{
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil{
			log.Fatal(err) 
		}
		result, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil{
			log.Fatal(err) 
		}
		dice := rand.Intn(100) + 1
		fmt.Print("정답을 입력하세요(1~100 사이의 정수) : ")
		if result == dice{
			fmt.Println("정답입니다.")
			break
		} else if result > dice{
			fmt.Println("정답보다 큰 수 입니다.")
		} else{
			fmt.Println("정답보다 작은 수 입니다.")
		}
	}
}