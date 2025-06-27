package slidingwindow

import "fmt"

// SubarrayAverage calculates the average of all contiguous subarrays of size k
// using the sliding window technique for O(n) time complexity
func SubarrayAverage(nums []int, k int) []float64 {
	// Validate input parameters to ensure they are within acceptable bounds
	if len(nums) == 0 || k <= 0 || k > len(nums) {
		// Print error message for invalid parameters
		fmt.Println("Incorrect params provided")
		// Return empty slice for invalid input
		return []float64{}
	}

	// Pre-allocate results slice with exact size needed: total possible subarrays of size k
	results := make([]float64, len(nums)-k+1)
	// Initialize both window pointers to start of array
	windowEnd, windowStart := 0, 0
	// Initialize running sum to track current window total
	sum := 0

	// Iterate through entire array using windowEnd pointer
	for windowEnd < len(nums) {
		// Add current element to running sum (expand window)
		sum += nums[windowEnd]

		// Check if window has reached desired size k
		if windowEnd >= k-1 {
			// Calculate average by dividing sum by window size k
			results[windowStart] = float64(sum) / float64(k)
			// Remove leftmost element from sum (shrink window from left)
			sum -= nums[windowStart]
			// Move window start pointer forward to maintain window size
			windowStart++
		}
		// Move window end pointer forward to next element
		windowEnd++
	}
	// Return slice containing all subarray averages
	return results
}

func SubarrayAveragePersonal(originalElements []int, windowSize int) []float32 {
	if len(originalElements) == 0 || windowSize <= 0 || windowSize > len(originalElements) {
		panic("YOU BROKE THIS.")
		// return []float32{}
	}

	results := make([]float32, (len(originalElements)-windowSize)+1)
	var windowStart, windowEnd int = 0, 0
	sumOfWindowNumbers := 0

	for windowEnd < len(originalElements) {
		sumOfWindowNumbers += originalElements[windowEnd]

		if windowEnd >= windowSize-1 {
			results[windowStart] = float32(sumOfWindowNumbers) / float32(windowSize)
			sumOfWindowNumbers -= originalElements[windowStart]
			windowStart++
		}
		windowEnd++
	}

	return results
}
