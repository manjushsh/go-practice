package twopointer

func PairWithTargetSum(arr []int, targetSum int) []int {
	startPointer := 0
	endPointer := len(arr) - 1

	for startPointer < endPointer {
		currentSum := arr[startPointer] + arr[endPointer]
		if currentSum == targetSum {
			return []int{startPointer, endPointer}
		}
		if currentSum < targetSum {
			startPointer++
		} else {
			endPointer--
		}
	}
	return []int{-1, -1}
}
