package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("a => ", a)
	a[0] = 20
	a[4] = 30
	fmt.Println("0, 4 index", a)
	var b = [2]string{"x", "y"}
	fmt.Println("b=>", b)
}
