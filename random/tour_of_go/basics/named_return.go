package basics

import (
	"fmt"
	"math"
	"math/cmplx"
)

func Reverse(x,y int) (int, int) {
	return y, x
}

func RandomSplit(num int) (x, y int) {
	x = num - 1
	y = num - x
	return
}

func DataTypes() {
	var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	num8   uint8	  = 255		// 256 will throw error
	pi	   float64	  = math.Pi
)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", num8, num8)
	fmt.Printf("Type: %T Value: %v\n", pi, pi)
}

func TypeCoversions() {
	var x, y = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	const langauge = "Go"
	fmt.Println(x, y, z, langauge)

	// Bit shifting

	const (
		// Create a huge number by shifting a 1 bit left 100 places.
		// In other words, the binary number that is 1 followed by 100 zeroes.
		Big = 1 << 100
		// Shift it right again 99 places, so we end up with 1<<1, or 2.
		Small = Big >> 99
	)

	fmt.Println(Big * 1.0, Small)

}


