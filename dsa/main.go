package main

import (
	"fmt"

	dsa "github.com/manjushsh/go-practice/dsa/2026-01"
)

func main() {
	fmt.Println("Hello there! Lets Go!")
	// fmt.Println(dsa.ContainerWithMostWater([]int{1,8,6,2,5,4,8,3,7}))
	// fmt.Println(dsa.TwoSumSortedArray([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 10))
	fmt.Println(dsa.TwoSum([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 5))
}
