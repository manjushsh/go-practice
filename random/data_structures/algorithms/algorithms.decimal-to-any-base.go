package algorithms

import "fmt"

func DecimalToAnyBaseDefault1(num, base int) string {
	if num == 0 {
		return "0"
	}
	var result string
	for num > 0 {
		remainder := num % base
		if remainder < 10 {
			result = string(rune(remainder)+'0') + result
		} else {
			result = string(rune(remainder-10)+'A') + result
		}
		num /= base
	}
	return result
}

func DecimalToAnyBaseDefault2(num, base int) string {
	var result string
	for num > 0 {
		remainder := num % base
		result = fmt.Sprintf("%X", remainder) + result
		num /= base
	}
	return result
}

func DecimalToAnyBase(num, base int) string {
	const charset = "0123456789ABCDEF"
	var result string
	for num > 0 {
		remainder := num % base
		result = string(charset[remainder]) + result
		num /= base
	}
	return result
}
