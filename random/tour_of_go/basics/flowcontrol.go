package basics

import (
	"fmt"
	"runtime"
)

func BasicFor() {
	for i := 0; i < 3; i++ {
		fmt.Println(i + 1)
	}
}

func NoInitAndPost() {
	sum := 0
	for ; sum < 2000; {
		sum += 1
	}
	fmt.Println(sum)
}

func WhileWithFor() {
	sum := 1
	for sum < 200 {
		sum += sum
	}
	fmt.Println(sum)

	// If you omit the loop condition it loops forever, so an infinite loop is compactly expressed. 
}

func BasicIf(num int) {
	if num %2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}
}

func IfWithStatement(x, y int) {
	if sum := x + y; sum < 10 {
		fmt.Println("Single digit sum. Sum:", sum)
	} else {
		fmt.Println("Number is greater than 10. Sum:", sum)
	}
}

func SwitchDetectOS() {
	switch os := runtime.GOOS; os {
	case "darwin": {
		fmt.Println("MacOS detected")
	}
	case "windows": {
	fmt.Println("Windows detected")
	}
	case "linux": {
		fmt.Println("Linux detected")
	}
	default: {
		fmt.Println("What OS are you using?")
	}
	}
}

func BasicDefer() {
	/**	
	 A defer statement defers the execution of a function until the surrounding function returns.

	The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns. 

	Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order. 
	LIFO
	
	**/
	defer fmt.Println("How are you?")	// <-- Pushed to stack at 0	- This runs second
	defer fmt.Println("There")			// <-- Pushed to stack at 1 - This runs first
	fmt.Println("Hello")
}



