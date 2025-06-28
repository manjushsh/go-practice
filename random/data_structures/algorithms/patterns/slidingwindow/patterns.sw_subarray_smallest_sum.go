package slidingwindow

func SmallestSubarrayWithGivenSum(numbers []int, targetSum int) (int, int) {
	leftPointer := 0
	smallestLength := len(numbers) + 1 // Start with impossible length
	windowSum := 0

	// Try each position as the right edge of our window
	for rightPointer := 0; rightPointer < len(numbers); rightPointer++ {
		// Add the new number to our window
		windowSum += numbers[rightPointer]

		// Try to shrink the window from the left while it's still valid
		for windowSum >= targetSum {
			// Update our best result if this window is smaller
			currentLength := rightPointer - leftPointer + 1
			if currentLength < smallestLength {
				smallestLength = currentLength
			}

			// Remove the leftmost number and shrink the window
			windowSum -= numbers[leftPointer]
			leftPointer++
		}
	}

	// Check if we found any valid window
	if smallestLength == len(numbers)+1 {
		return -1, -1 // No solution found
	}
	return leftPointer - 1, smallestLength
}

func MaxSubarrayWithGivenSum(numbers []int, exactSum int) (int, int) {
	leftPointer := 0
	longestLength := 0
	windowSum := 0
	bestStartIndex := -1

	// Try each position as the right edge of our window
	for rightPointer := 0; rightPointer < len(numbers); rightPointer++ {
		// Add the new number to our window
		windowSum += numbers[rightPointer]

		// If sum is too big, shrink window from left
		for windowSum > exactSum && leftPointer <= rightPointer {
			windowSum -= numbers[leftPointer]
			leftPointer++
		}

		// Check if we found the exact sum we want
		if windowSum == exactSum {
			currentLength := rightPointer - leftPointer + 1
			if currentLength > longestLength {
				longestLength = currentLength
				bestStartIndex = leftPointer
			}
		}
	}

	// Check if we found any valid window
	if longestLength == 0 {
		return -1, -1 // No solution found
	}
	return bestStartIndex, longestLength
}
