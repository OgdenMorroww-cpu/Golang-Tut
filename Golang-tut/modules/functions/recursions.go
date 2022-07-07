package functions

import "fmt"

func MyRecursion() {
	fmt.Println(factorials(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}

func factorials(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorials(n-1)
}
