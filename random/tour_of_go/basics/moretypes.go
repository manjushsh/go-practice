package basics

import (
	"fmt"
	"strings"
	"math"

	"golang.org/x/tour/pic"
)

func BasicPointer() {
	num, str := 32, "hello"
	pointerInt, pointerStr := &num, &str

	fmt.Println(num, "-->", pointerInt, str, "-->", pointerStr, *pointerInt, *pointerStr)
}

func BasicArray() {
	var arr [10]int
	arr2 := [2]int{19, 55}
	fmt.Println(arr, arr2)
}

/**	
	A slice does not store any data, it just describes a section of an underlying array.
	Changing the elements of a slice modifies the corresponding elements of its underlying array.
	Other slices that share the same underlying array will see those changes.

	When slicing, you may omit the high or low bounds to use their defaults instead.
	The default is zero for the low bound and the length of the slice for the high bound. 
**/
func BasicSlice1() {
	primes := [6]int{2,3,5,7,11}	// Last element will be automatically 0 as array len is 6 but we have 5 values
	sliceOfPrimes := primes[1:4]	// should return [3,5,7] i.e, includes index 1 but omits index 4
	fullSlice := primes[:]
	startToLen := primes[2:]
	zeroIndexToLen := primes[:4]
	fmt.Println(sliceOfPrimes, fullSlice, startToLen, zeroIndexToLen)
}

/** 
	A slice has both a length and a capacity.
	The length of a slice is the number of elements it contains.
	The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
	You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
	The zero value of a slice is nil
**/
func SliceAndCapacity() {
	array := []int{2, 3, 5, 7, 11, 13}
	modifiedArrSlice := array[:]
	modifiedArrSlice = modifiedArrSlice[:4]
	modifiedArrSlice = modifiedArrSlice[2:]

	fmt.Printf("len = %d and capacity = %d. Array %v \n", len(modifiedArrSlice), cap(modifiedArrSlice), modifiedArrSlice)

	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("Nil")
	}

	a := make([]int, 5)
	fmt.Println("a", a)

	b := make([]int, 0, 5)
	fmt.Println("b", b, len(b), cap(b))

	c := b[:2]
	fmt.Println("c", c)

	d := c[2:5]
	fmt.Println("d", d)
}

func TwoDSlices() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func BasicRanges() {
	array := []int{2, 3, 5, 7, 11, 13}
	for index, value := range array {
		fmt.Println(index, value * 2)
	}
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = uint8(x ^ y)
			// p[y][x] = uint8(x*y)
			// p[y][x] = uint8((x+y)/2)
		}
	}
	return pic
}

func BasicShowPic() {
	pic.Show(Pic)
}

func BasicMap() {

	type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var m2 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

m2["Linux"] = Vertex{0.00, 0.00}
delete(m, "Google")
elem, ok := m2["Google"]
fmt.Println(m2, m, elem, ok)

}

func WordCount(s string) map[string]int {
	var countMap = make(map[string]int)
	for _, word := range strings.Fields(s) {
		countMap[word]++
	}
	return countMap
}

func BasicWordCount() {
	fmt.Println(WordCount("Hello there. Who is there. there."))
}

func Pythogorus(a, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}

func FuncAsParam(fnParam func(float64, float64) float64) float64 {
	const num1, num2 = 3, 4

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Printf("Square root of %d and %d is %v \n", num1, num2, fnParam(3, 4))
	fmt.Println(hypot(5, 12))
	return fnParam(num1, num2)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func Closure() {
	pos, neg := adder(), adder()
	fmt.Println("Sums")
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first, second := 0, 1
    return func() int {
        ret := first
        first, second = second, first+second
        return ret
    }
}

func Fibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
