package array

// FindLowestNum finds the lowest number in a given slice of integers.
// It returns the lowest number and its index in the slice.
// If the slice is empty, it returns -1 and -1.
func FindLowestNum(arr []int) (int, int) {
	if len(arr) == 0 {
		return -1, -1 // Return -1 for both value and index if the array is empty
	}

	lowest := arr[0]
	index := 0

	for i, num := range arr {
		if num < lowest {
			lowest = num
			index = i
		}
	}

	return lowest, index
}