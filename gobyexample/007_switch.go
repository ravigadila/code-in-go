package main

import (
	"fmt"
	"time"
)

func main() {
	const i int = 2
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend")
	default:
		fmt.Println("workday")
	}

	whichType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i is boolean")
		case int:
			fmt.Println("i in int")
		default:
			fmt.Println("Some Other", t)
		}

	}
	whichType(true)
	whichType(10.2)
}
