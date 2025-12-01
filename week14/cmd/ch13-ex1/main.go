package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := range 3 {
		fmt.Println(msg, ":", i)
		time.Sleep(2000 * time.Millisecond)
	}
}

func say2(msg string) {
	for i := range 3 {
		fmt.Println(msg, ":", i)
		time.Sleep(1500 * time.Millisecond)
	}
}

func main() {
	start := time.Now()

	// 4.5초 걸림
	// go say("고루틴") // 새 고루틴에서 실행
	// say2("메인")     // 메인 고루틴에서 실행

	// 6초 걸림
	go say2("고루틴") // 새 고루틴에서 실행
	say("메인")     // 메인 고루틴에서 실행

	fmt.Println("전체 실행 시간 : ", time.Since(start))
}