package algorithms

func SumAllNumList(list []int) int {
	sum := 0
	if list == nil {
		return sum
	}
	for _, num := range list {
		sum += num
	}
	return sum
}

func SumAllNumListRecursive(list []int) int {
	if list == nil {
		return 0
	}
	if len(list) == 0 {
		return 0
	}
	return list[0] + SumAllNumListRecursive(list[1:])
}
