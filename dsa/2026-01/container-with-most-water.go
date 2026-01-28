package dsa

func ContainerWithMostWater(nums []int) int {
	start, end, maxArea := 0, len(nums)-1, 0

	for start < end {
		width := end - start
		height := min(nums[start], nums[end])
		currentArea := width * height
		maxArea = max(maxArea, currentArea)

		if nums[start] < nums[end] {
			start++
		} else {
			end--
		}
	}
	return maxArea
}
