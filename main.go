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

// 리턴값에 변수명을 지정해줄 수 있음
func lengAndUpperV2(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done") // defer는 함수가 끝나고 실행됨
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// range 는 for문에서 사용할 수 있음
func superAdd(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total = total + num
	}
	return total
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canIDrinkSwitch(age int) bool {
	switch {
	case age < 18:
		return false
	case 18 == 18:
		return true
	}
	return false
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

	totalLength, upperName = lengAndUpperV2("ricepotato2")
	fmt.Println(totalLength)
	fmt.Println(upperName)

	repeatMe("ricepotato", "sukjun", "sagong")

	total := superAdd(1, 3, 3, 4, 5, 4, 7, 8, 2, 0)
	fmt.Println(total)

	fmt.Println(canIDrink(16))
	fmt.Println(canIDrinkSwitch(30))
}
