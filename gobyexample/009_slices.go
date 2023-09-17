package main

import (
	"fmt"
	"slices"
)

func main() {
	var a []int
	fmt.Println("a=>", a)
	a = append(a, 10)
	fmt.Println("a=>", len(a))

	var s []string
	s = make([]string, 3)
	fmt.Println("s=>", len(s), s)
	s[0] = "a"
	s[1] = "b"
	s = append(s, "d", "e")
	fmt.Println("s=>", len(s), s)
	fmt.Println("s[:4]", s[:4])
	fmt.Println("s[:4]", s[2:])

	t := []string{"a", "b", "", "d", "e"}
	if slices.Equal(t, s) {
		fmt.Println("t==s")
	}
	if slices.Contains(s, "d") {
		fmt.Println("True d in s")
	}
}
