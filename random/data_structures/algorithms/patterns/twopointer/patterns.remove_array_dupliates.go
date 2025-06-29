package twopointer

func RemoveArrayDuplicates(originalArray []int) []int {
	nextNonDup := 1

	for currentElement := 1; currentElement < len(originalArray); currentElement++ {
		if originalArray[nextNonDup-1] != originalArray[currentElement] {
			originalArray[nextNonDup] = originalArray[currentElement]
			nextNonDup++
		}
	}
	return originalArray[:nextNonDup]
}
