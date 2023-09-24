package main

import "fmt"

func numsqrt(a int) int {
	return a * a
}

func ptrsqrt(a *int) int {
	return *a * *a
}

func main() {
	i := 10
	fmt.Println(numsqrt(i))
	fmt.Println(ptrsqrt(&i))
}
