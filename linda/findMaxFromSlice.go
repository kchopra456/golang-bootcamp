package main

import (
	"fmt"
)

func maxFromSlice() {
	sl := []int{4, 3, 5, 6, 7, 3, 5, 8}

	max, index := sl[0], 0
	for i, v := range sl {
		if max < v {
			max = v
			index = i
		}
	}
	fmt.Printf("Slice: %v, maximum value: %d, at index: %d", sl, max, index)
}
