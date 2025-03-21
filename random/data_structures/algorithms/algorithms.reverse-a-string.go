package algorithms

import "strings"

func ReverseAString(s string) string { //	Swap
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReverseAString2(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}
	return string(runes)
}

func ReverseAStringRecursive(s string) string {
	if len(s) == 0 {
		return s
	}
	return ReverseAStringRecursive(s[1:]) + s[:1]
}

func ReverseAString3(s string) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

func ReverseAString4(s string) string {
	var reversed string
	for _, rune := range s {
		reversed = string(rune) + reversed
	}
	return reversed
}
