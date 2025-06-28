package slidingwindow

func SumOfSubArrays(originalElements []int, windowSize int) []int {
	if windowSize < 0 || windowSize > len(originalElements) {
		return make([]int, 0)
	}

	windowStart, sumOfWindowElements := 0, 0
	results := make([]int, (len(originalElements)-windowSize)+1)

	for windowEnd := 0; windowEnd < len(originalElements); windowEnd++ {
		// or could be "for windowEnd := range originalElements {"
		sumOfWindowElements += originalElements[windowEnd]

		if windowEnd >= windowSize-1 {
			results[windowStart] = sumOfWindowElements
			sumOfWindowElements -= originalElements[windowStart]
			windowStart++
		}
		// Increment windowEnd will be done by the for loop
	}
	return results
}

func MaxSumOfSubArrays(originalElements []int, windowSize int) int {
	if windowSize < 0 || len(originalElements) == 0 || windowSize > len(originalElements) {
		return -1
	}

	windowStart, sumOfWindowElements := 0, 0
	results := make([]int, (len(originalElements)-windowSize)+1)

	for windowEnd := range originalElements {
		// or could be "for windowEnd := range originalElements {"
		sumOfWindowElements += originalElements[windowEnd]

		if windowEnd >= windowSize-1 {
			results[windowStart] = sumOfWindowElements
			sumOfWindowElements -= originalElements[windowStart]
			windowStart++
		}
		windowEnd++
	}
	return maxInSlice(results)
}



// maxInSlice returns the maximum value in a slice of ints.
// Assumes the slice is non-empty.
func maxInSlice(nums []int) int {
	maxVal := nums[0]
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
