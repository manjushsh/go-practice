package algorithms

// IsNumInList checks if a number is in a list
// and returns true if it is, false otherwise.

import "slices"

func IsNumInList(list []int, num int) bool {
	// https://pkg.go.dev/slices#Contains
	return slices.Contains(list, num)
}
