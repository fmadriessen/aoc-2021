package main

import (
	"aoc/helper"
	"fmt"
	"strconv"
)

func singleSlidingWindow(depths []int) (result int) {
	prevDepth := 0
	for i, depth := range depths {
		if i != 0 && depth > prevDepth {
			result++
		}
		prevDepth = depth
	}
	return result
}

func threeSlidingWindow(depths []int) (result int) {
	result = 0
	for i := 1; i+2 < len(depths); i++ {
		// Since they have two same measurements, sufficies to compare depth[i - 1] and depth[i + 2]
		if depths[i-1] < depths[i+2] {
			result++
		}
	}
	return result
}

func main() {
	inputs := helper.ReadInput("input")
	depths := make([]int, len(inputs))

	for i, input := range inputs {
		depths[i], _ = strconv.Atoi(input)
	}

	fmt.Printf("Using a single sample sliding window there are %v depth increases\n", singleSlidingWindow(depths))
	fmt.Printf("Using a three sample sliding window there are %v depth increases\n", threeSlidingWindow(depths))
}
