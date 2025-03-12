package algorithms

// IsNumInList checks if a number is in a list
// and returns true if it is, false otherwise.

import "slices"

func IsNumInList(list []int, num int) bool {
	// https://pkg.go.dev/slices#Contain
	if list == nil {
		return false
	}
	if len(list) == 0 {
		return false
	}
	return slices.Contains(list, num) || false
}
