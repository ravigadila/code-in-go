package main

import (
	"fmt"
)

func main() {
	fmt.Println("Introduction to IF/else in go")
	var a int = 10
	var b int = 20

	if b%a == 0 {
		fmt.Println("b is divisible by a")
	}
	if n := 15; n%10 == 0 {
		fmt.Println("n div by 10")
	} else if n%5 == 0 {
		fmt.Println("n div by 5")
	} else {
		fmt.Println("n is strange")
	}
}
