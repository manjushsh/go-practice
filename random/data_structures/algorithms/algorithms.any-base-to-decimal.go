package algorithms

import "math"

func AnyBaseToDecimal(num string, base int) int {
	var decimalValue int
	length := len(num)

	for i, digit := range num {
		if digit >= '0' && digit <= '9' {
			decimalValue += int(digit-'0') * int(math.Pow(float64(base), float64(length-i-1)))
		} else if digit >= 'A' && digit <= 'F' {
			decimalValue += (int(digit-'A') + 10) * int(math.Pow(float64(base), float64(length-i-1)))
		} else if digit >= 'a' && digit <= 'f' {
			decimalValue += (int(digit-'a') + 10) * int(math.Pow(float64(base), float64(length-i-1)))
		}
	}

	return decimalValue
}