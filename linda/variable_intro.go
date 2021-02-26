package main

import (
	"fmt"
)

func varIntro() {
	// runtime default is 0
	var x int
	var y int

	fmt.Printf("Default value of %T is %v\n", x, x)
	fmt.Printf("Default value of %T is %v\n", y, y)

	x = 3
	y = 4

	var mean int = (x + y) / 2
	fmt.Printf("mean: %v", mean)

}
