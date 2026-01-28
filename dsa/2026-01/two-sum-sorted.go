package dsa

func TwoSumSortedArray(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	start, end := 0, len(nums)-1
	for start < end {
		sum := nums[start] + nums[end]
		if sum == target {
			// Numbers add up to target, so return index
			return []int{start, end}
		} else if sum < target {
			// If sum is smaller than target, we need larger number at left as right is largest number anyway. so increment start pointer
			start++
		} else {
			// Sum doesnt match and it is higher so reduce right pointer as right element is too large
			end--
		}
	}
	return []int{-1, -1}
}
