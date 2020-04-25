package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

//이름을 받으면 첫번째 리턴으로 글자의 숫자, 1. 이름 으로 출력해주는 함수
func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("I'm Done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}
func main() {

	totalLength, upperName := lenAndUpper("sangwoo cho")
	fmt.Println(totalLength, upperName)
	repeatMe("one", "two", "three")
}
