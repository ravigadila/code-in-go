package main

import "fmt"

func sum(nums ...int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}

func main() {
	res := sum(1, 2, 3, 4, 5)
	fmt.Println("res ", res)
}
