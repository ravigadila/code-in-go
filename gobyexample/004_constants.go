package main

import (
	"fmt"
)

const s string = "Ravi"

func main() {
	var v1 = "ABC"
	fmt.Println("S out: ", s)

	v1 = "DEF"
	fmt.Println(v1)

	const n = 500000000
	const d = 3e20 / n

	fmt.Printf("d %T %v \n", d, d)

	// var a1 = 40 / 10
	// you cannot assign a non-constant value (like a variable) to a constant,"
	// const x = 40 / a1  // ERROR

	// fmt.Printf("x %T %v \n", x, x)
	const a2 = "GO is FUN"
	fmt.Printf("a2 %T %v \n", a2, a2)
}
