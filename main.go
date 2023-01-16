package main

import (
	"fmt"
	"strings"
)

// 함수정의. func 키워드로 시작
// 타입을 지정해줘야 함, 뒤에만 타입을 지정해줘도 됨
func multiply(a, b int) int {
	return a * b
}

// 함수가 여러개의 값을 리턴할 수 있음
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// ...을 사용하면 여러개의 인자를 받을 수 있음
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	const name string = "ricepotato"  // initialize constants
	var nameVar string = "ricepotato" // initialize variables
	nameExp := "ricepotato"           // initialize var. only in function

	fmt.Println("Hello World")
	fmt.Println(name)
	fmt.Println(nameVar)
	fmt.Println(nameExp)

	fmt.Println(multiply(2, 3))
	fmt.Println(lenAndUpper("ricepotato")) // 길이와 대문자로 바꾼 문자열을 리턴
	totalLength, upperName := lenAndUpper("ricepotato")
	fmt.Println(totalLength)
	fmt.Println(upperName)

	repeatMe("ricepotato", "sukjun", "sagong")
}
