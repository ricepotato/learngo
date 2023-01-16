package main

import "fmt"

func main() {
	const name string = "ricepotato"  // constants
	var nameVar string = "ricepotato" // variables
	nameExp := "ricepotato"           // var. only in function

	fmt.Println("Hello World")
	fmt.Println(name)
	fmt.Println(nameVar)
	fmt.Println(nameExp)
}
