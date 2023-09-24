package main

import (
	"fmt"
	"maps"
)

func main() {
	var d1 = make(map[string]int)
	fmt.Println("d1=>", d1)
	d1["k1"] = 10
	d1["k2"] = 20
	fmt.Println("len d1=>", len(d1))
	fmt.Println("val k2=>", d1["k2"])
	fmt.Println("val k2=>", d1["NA"]) // returns null value
	d2 := map[string]int{"k1": 10, "k2": 20}
	fmt.Println("d2=>", d2)

	if maps.Equal(d1, d2) {
		fmt.Println("Ditto")
	}
}
