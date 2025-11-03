package main

import (
	"fmt"
	"log"
	"week10/pkg/keyboard"
)

func main() {
	fmt.Print("1. 섭씨 -> 화씨\n2. 화씨 -> 섭씨\n메뉴 번호를 입력하세요 : ")
	choice, err := keyboard.GetInt()
	if err!=nil{
		log.Fatal(err)
	}
	if choice == 1{
		fmt.Print("섭씨 온도를 입력하세요 : ")
		celsius, err := keyboard.GetFloat()
		if err!=nil{
			log.Fatal(err)
		}
		fmt.Printf("섭씨 %.2f도는 화씨 %.2f도입니다.\n", celsius, (celsius * 9/5) + 32)
	}else if choice == 2{
		fmt.Print("화씨 온도를 입력하세요 : ")
		fahrenheit, err := keyboard.GetFloat()
		if err!=nil{
			log.Fatal(err)
		}
		fmt.Printf("화씨 %.2f도는 섭씨 %.2f도입니다.", fahrenheit, (fahrenheit - 32) * 5/9)
	}else{
		log.Fatal("존재하지 않는 메뉴 번호입니다.")
	}
}
