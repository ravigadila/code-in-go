package main

import "fmt"

func main() {
	// assigning var without specifying type
	var a = 100
	fmt.Printf("type of a %T value %v\n", a, a)

	// specify variable type
	var b string = "hello"
	fmt.Printf("type of b %T value %v\n", b, b)

	var x, y int = 100, 200
	fmt.Println("x y", x, y)
	var nov = 50
	fmt.Println(nov)
	// declare and add values
	ni := "some value"
	fmt.Println(ni)

	var ei int
	fmt.Printf("tyoe %T val %v \n", ei, ei)
}
