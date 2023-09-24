package main

import (
	"fmt"
)

func main() {
	n := []int{5, 10, 15}
	for i, v := range n {
		fmt.Println("=>", i, v)
	}

	fruits := map[string]string{"a": "apple", "b": "ball"}
	for k, v := range fruits {
		fmt.Println(k, "for", v)
	}
}
