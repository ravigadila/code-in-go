package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func engVal(a int) (int, string) {
	return a, "integer"
}

func main() {
	var x, y int = 100, 200
	fmt.Println("x+y = ", plus(x, y))
	fmt.Println("10+20+30 = ", plusplus(10, 20, 30))
	a, s := engVal(10)
	fmt.Println("Eng = ", a, s)
}
