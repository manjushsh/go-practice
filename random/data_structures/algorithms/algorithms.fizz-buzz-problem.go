package algorithms

import (
	"fmt"
	"strings"
)

func FizzBuzz(n int) string {
	var result []string
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0: //	or i%15 == 0
			result = append(result, "FizzBuzz")
		case i%3 == 0:
			result = append(result, "Fizz")
		case i%5 == 0:
			result = append(result, "Buzz")
		default:
			result = append(result, fmt.Sprintf("%d", i))
		}
	}
	return strings.Join(result, ", ")
}

func FizzBuzzOld1(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, string(i))
		}
	}
	return result
}
