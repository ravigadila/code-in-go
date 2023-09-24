package main

import "fmt"

func incVal() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	incV1 := incVal()
	fmt.Println(incV1())
	fmt.Println(incV1())
	fmt.Println(incV1())
	fmt.Println(incV1())

	incV2 := incVal()
	fmt.Println(incV2())
	fmt.Println(incV2())
	fmt.Println(incV2())

}
