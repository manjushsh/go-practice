package dsa

func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if index, found := numMap[diff]; found {
			return []int{i, index}
		}
		numMap[nums[i]] = i
	}
	return []int{-1, -1}
}
