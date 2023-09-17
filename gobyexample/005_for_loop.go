package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	var i = 10
	for i < 20 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}
	for true {
		fmt.Println("Some while with breack")
		break
	}
}
