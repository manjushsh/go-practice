package algorithms

func SortWorst(arrayToSort []int) []int {

	arrayLen := len(arrayToSort)
	for arrayLen != 0 {
		for index := 1; index < arrayLen; index++ {
			if arrayToSort[index-1] > arrayToSort[index] {
				arrayToSort[index-1], arrayToSort[index] = arrayToSort[index], arrayToSort[index-1]
			}
		}
		arrayLen--
	}
	return arrayToSort
}
